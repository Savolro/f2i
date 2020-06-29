package rgba

import (
	"fmt"
	"image"
	"log"
)

// EndByte symbolizes the end of data in an image
const EndByte = 0x69

// Encode encodes specified data into a RGBA image of given size
func Encode(data []byte, width, height int) (*image.RGBA, error) {
	// 4 bytes can be stored in every pixel (RGBA)
	maxLength := width * height * 4
	if len(data) > maxLength {
		return nil, fmt.Errorf("data does not fit the image size")
	}
	img := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{width, height}})
	img.Pix = data
	if len(img.Pix) < maxLength {
		img.Pix = append(img.Pix, EndByte)
	}
	for i := len(img.Pix); i < maxLength; i++ {
		img.Pix = append(img.Pix, 0)
	}
	return img, nil
}

// Decode decodes a RGBA image to an original data
func Decode(img *image.RGBA) ([]byte, error) {
	i := len(img.Pix) - 1
	for ; img.Pix[i] == 0 && i >= 0; i-- {
	}
	if img.Pix[i] != EndByte {
		log.Println(img.Pix[i])
		return nil, fmt.Errorf("last byte of an image is not an end byte")
	}
	return img.Pix[:i], nil
}
