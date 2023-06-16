package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	_ "runtime/pprof"

	"github.com/davidbyttow/govips/v2/vips"
)

func main() {
	imglib := flag.String("lib", "", "current environment")
	flag.Parse()
	println(*imglib)
	switch *imglib {
	case "stdlib":
		runStdLib()
	case "vips":
		runVips()
	default:
		println("i dont understand, try run with go run . -imglib stdlib")
	}
}

func getFile() *bytes.Buffer {
	f, err := os.Open("./assets/test-limit.jpg")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, f)
	if err != nil {
		panic(err)
	}

	return buf
}

const (
	MAX_WIDTH_THUMBNAIL  = 200
	MAX_HEIGHT_THUMBNAIL = 200
)

func getThumbnailSize(originWidth, originHeight int) (int, int) {
	if originWidth <= MAX_WIDTH_THUMBNAIL && originHeight <= MAX_HEIGHT_THUMBNAIL {
		// no need to be resize
		return originWidth, originHeight
	}
	var width, height = uint(originWidth), uint(originHeight)

	// Preserve aspect ratio
	if width > MAX_WIDTH_THUMBNAIL {
		height = uint(height * MAX_WIDTH_THUMBNAIL / width)
		if height < 1 {
			height = 1
		}
		width = MAX_WIDTH_THUMBNAIL
	}

	if height > MAX_HEIGHT_THUMBNAIL {
		width = uint(width * MAX_HEIGHT_THUMBNAIL / height)
		if width < 1 {
			width = 1
		}
		height = MAX_HEIGHT_THUMBNAIL
	}
	return int(width), int(height)
}

func runStdLib() {
	buf := getFile()

	r, err := stdlib(context.Background(), buf.Bytes())
	if err != nil {
		panic(err)
	}

	ioutil.WriteFile("stdlib-result.jpeg", r, 0644)
}

func runVips() {
	vips.Startup(&vips.Config{
		ConcurrencyLevel: 1,
		MaxCacheFiles:    1,
		CollectStats:     true,
	})
	defer vips.Shutdown()
	stt := vips.MemoryStats{}
	buf := getFile()

	r, err := vipsz(buf.Bytes())
	if err != nil {
		panic(err)
	}
	ioutil.WriteFile("vips-result.jpeg", r, 0644)

	vips.ReadVipsMemStats(&stt)
	fmt.Printf("%+v\n", stt)
}
