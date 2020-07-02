package nrgba

import (
	"image"
	"testing"

	"github.com/bmizerany/assert"
)

func TestEncode(t *testing.T) {
	tests := []struct {
		name    string
		data    []byte
		expData []byte
		width   int
		height  int
		err     bool
	}{
		{
			name:    "full write",
			data:    []byte{1, 2, 3, 4},
			expData: []byte{1, 2, 3, 4},
			width:   1,
			height:  1,
			err:     false,
		}, {
			name:    "partial write",
			data:    []byte{1, 2, 3},
			expData: []byte{1, 2, 3, 0},
			width:   1,
			height:  1,
			err:     false,
		}, {
			name:    "bigger image",
			data:    []byte{1, 2, 3, 4, 5},
			expData: []byte{1, 2, 3, 4, 5, 0, 0, 0},
			width:   2,
			height:  1,
			err:     false,
		}, {
			name:    "too big data",
			data:    []byte{1, 2, 3, 4, 5},
			expData: nil,
			width:   1,
			height:  1,
			err:     true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			img, err := Encode(tt.data, tt.width, tt.height)
			assert.Equal(t, tt.err, err != nil, err)
			if !tt.err {
				assert.Equal(t, tt.expData, img.Pix, img.Pix)
				assert.Equal(t, tt.width, img.Rect.Max.X)
				assert.Equal(t, tt.height, img.Rect.Max.Y)
			}
		})
	}
}

func TestDecode(t *testing.T) {
	tests := []struct {
		name      string
		img       *image.NRGBA
		data      []byte
		trimZeros bool
		err       bool
	}{
		{
			name: "valid image",
			img: &image.NRGBA{
				Pix: []byte{1, 2, 3},
				Rect: image.Rectangle{
					Min: image.Point{
						X: 0,
						Y: 0,
					},
					Max: image.Point{
						X: 1,
						Y: 1,
					},
				},
			},
			data: []byte{1, 2, 3},
			err:  false,
		},
		{
			name: "full image",
			img: &image.NRGBA{
				Pix: []byte{1, 2, 3, 4},
				Rect: image.Rectangle{
					Min: image.Point{
						X: 0,
						Y: 0,
					},
					Max: image.Point{
						X: 1,
						Y: 1,
					},
				},
			},
			data: []byte{1, 2, 3, 4},
			err:  false,
		},
		{
			name: "zeroed image, trim",
			img: &image.NRGBA{
				Pix: []byte{1, 2, 3, 0},
				Rect: image.Rectangle{
					Min: image.Point{
						X: 0,
						Y: 0,
					},
					Max: image.Point{
						X: 1,
						Y: 1,
					},
				},
			},
			data:      []byte{1, 2, 3},
			trimZeros: true,
			err:       false,
		},
		{
			name: "zeroed image, no trim",
			img: &image.NRGBA{
				Pix: []byte{1, 2, 3, 0},
				Rect: image.Rectangle{
					Min: image.Point{
						X: 0,
						Y: 0,
					},
					Max: image.Point{
						X: 1,
						Y: 1,
					},
				},
			},
			data:      []byte{1, 2, 3, 0},
			trimZeros: false,
			err:       false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := Decode(tt.img, tt.trimZeros)
			assert.Equal(t, tt.err, err != nil, err)
			assert.Equal(t, tt.data, data)
		})
	}
}
