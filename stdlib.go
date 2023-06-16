package main

import (
	"bytes"
	"context"
	"image"
	"image/jpeg"
	"image/png"

	"golang.org/x/image/bmp"
	"golang.org/x/image/draw"
)

func stdlib(ctx context.Context, file []byte) ([]byte, error) {
	fs := new(bytes.Buffer)
	_, err := fs.Write(file)
	if err != nil {
		return nil, err
	}
	src, s, err := image.Decode(bytes.NewBuffer(file))
	if err != nil {
		return nil, err
	}

	width, height := src.Bounds().Dx(), src.Bounds().Dy()
	if width <= MAX_WIDTH_THUMBNAIL && height <= MAX_HEIGHT_THUMBNAIL {
		return fs.Bytes(), nil
	}

	x, y := getThumbnailSize(width, height)
	dst := image.NewRGBA(image.Rect(0, 0, x, y))
	draw.NearestNeighbor.Scale(dst, dst.Rect, src, src.Bounds(), draw.Over, nil)

	output := new(bytes.Buffer)
	switch s {
	case "jpeg", "jpg", "webp":
		err = jpeg.Encode(output, dst, nil)
		if err != nil {
			return nil, err
		}
	case "png":
		err = png.Encode(output, dst)
		if err != nil {
			return nil, err
		}
	case "bmp":
		err = bmp.Encode(output, dst)
		if err != nil {
			return nil, err
		}
	default:
		return fs.Bytes(), nil
	}

	return output.Bytes(), nil
}
