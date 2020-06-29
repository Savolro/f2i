package rgba

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
			name:    "last byte empty",
			data:    []byte{1, 2, 3},
			expData: []byte{1, 2, 3, EndByte},
			width:   1,
			height:  1,
			err:     false,
		}, {
			name:    "only first byte",
			data:    []byte{1},
			expData: []byte{1, EndByte, 0, 0},
			width:   1,
			height:  1,
			err:     false,
		}, {
			name:    "bigger image",
			data:    []byte{1, 2, 3, 4, 5},
			expData: []byte{1, 2, 3, 4, 5, EndByte, 0, 0},
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
				assert.Equal(t, tt.expData, img.Pix)
				assert.Equal(t, tt.width, img.Rect.Max.X)
				assert.Equal(t, tt.height, img.Rect.Max.Y)
			}
		})
	}
}

func TestDecode(t *testing.T) {
	tests := []struct {
		name string
		img  *image.RGBA
		data []byte
		err  bool
	}{
		{
			name: "valid image",
			img: &image.RGBA{
				Pix: []byte{1, 2, 3, EndByte},
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
		}, {
			name: "end byte in data",
			img: &image.RGBA{
				Pix: []byte{1, EndByte, 3, EndByte},
				Rect: image.Rectangle{
					Max: image.Point{
						X: 1,
						Y: 1,
					},
				},
			},
			data: []byte{1, EndByte, 3},
			err:  false,
		}, {
			name: "invalid image",
			img: &image.RGBA{
				Pix: []byte{1, EndByte, 3, 4},
				Rect: image.Rectangle{
					Max: image.Point{
						X: 1,
						Y: 1,
					},
				},
			},
			data: nil,
			err:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := Decode(tt.img)
			assert.Equal(t, tt.err, err != nil, err)
			assert.Equal(t, tt.data, data)
		})
	}
}
