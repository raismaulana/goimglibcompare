package main

import (
	"github.com/davidbyttow/govips/v2/vips"
)

func vipsz(file []byte) ([]byte, error) {
	image1, err := vips.NewImageFromBuffer(file)
	if err != nil {
		panic(err)
	}

	image1.Thumbnail(200, 200, vips.InterestingCentre)
	z, _, err := image1.ExportJpeg(vips.NewJpegExportParams())
	if err != nil {
		panic(err)
	}

	return z, nil
}
