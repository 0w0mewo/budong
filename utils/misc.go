package utils

import (
	"bytes"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

func ImageBytesFmt(b []byte) string {
	_, imgType, err := image.Decode(bytes.NewBuffer(b))
	if err != nil {
		return "unknown"
	}

	return imgType
}
