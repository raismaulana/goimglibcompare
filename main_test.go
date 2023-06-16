package main

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/davidbyttow/govips/v2/vips"
)

func Benchmark_Stdlib(b *testing.B) {

	buf := getFile()
	b.Run("Stdlib", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, err := stdlib(context.Background(), buf.Bytes())
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

func Benchmark_Vips(b *testing.B) {
	vips.LoggingSettings(func(messageDomain string, messageLevel vips.LogLevel, message string) {
		log.Printf("[%v.%v] %v", messageDomain, messageLevel, message)
	}, vips.LogLevelError)
	vips.Startup(&vips.Config{
		ConcurrencyLevel: 1,
		MaxCacheFiles:    1,
		MaxCacheMem:      1,
		CollectStats:     true,
	})
	defer vips.Shutdown()
	stt := vips.MemoryStats{}
	buf := getFile()
	runCount := int64(0)
	b.Run("Vips", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, err := vipsz(buf.Bytes())
			if err != nil {
				b.Fatal(err)
			}
		}
		runCount = int64(b.N)
	})
	vips.ReadVipsMemStats(&stt)
	fmt.Printf("Benchmark_Vips/Vips Average Memory Usage: %d B/op\n", stt.Mem/runCount)
	fmt.Printf("%+v\n", stt)
}
