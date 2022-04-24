package services

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/jpeg"
	"strings"
)

func DecodeBaseImageToBytes(base64Image string) ([]byte, error) {
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(base64Image))
	m, _, err := image.Decode(reader)
	if err != nil {
		return make([]byte, 0), err
	}

	buf := new(bytes.Buffer)
	err = jpeg.Encode(buf, m, nil)
	if err != nil {
		return make([]byte, 0), err
	}

	return buf.Bytes(), nil
}
