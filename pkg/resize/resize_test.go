package resize

import (
	"image"
	"image/color"
	"testing"
)

func TestResize(t *testing.T) {
	tests := []struct {
		name           string
		srcImage       *image.RGBA
		targetWidth    int
		targetHeight   int
		expectedWidth  int
		expectedHeight int
		expectedColor  color.Color
	}{
		{
			name:           "2x2 to 1x1, top-left pixel",
			srcImage:       newTestImage(),
			targetWidth:    1,
			targetHeight:   1,
			expectedWidth:  1,
			expectedHeight: 1,
			expectedColor:  color.RGBA{255, 0, 0, 255},
		},
		{
			name:           "2x2 to 2x2, same size",
			srcImage:       newTestImage(),
			targetWidth:    2,
			targetHeight:   2,
			expectedWidth:  2,
			expectedHeight: 2,
			expectedColor:  color.RGBA{255, 0, 0, 255},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			newImg := Resize(tt.srcImage, tt.targetWidth, tt.targetHeight)

			// リサイズ後の画像サイズをチェック
			if newImg.Bounds().Dx() != tt.expectedWidth || newImg.Bounds().Dy() != tt.expectedHeight {
				t.Errorf("Expected dimensions %dx%d but got %dx%d", tt.expectedWidth, tt.expectedHeight, newImg.Bounds().Dx(), newImg.Bounds().Dy())
			}

			// リサイズ後の画像の色をチェック
			if newImg.At(0, 0) != tt.expectedColor {
				t.Errorf("Expected color %v but got %v", tt.expectedColor, newImg.At(0, 0))
			}
		})
	}
}

func newTestImage() *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, 2, 2))
	img.Set(0, 0, color.RGBA{255, 0, 0, 255})
	img.Set(1, 0, color.RGBA{0, 255, 0, 255})
	img.Set(0, 1, color.RGBA{0, 0, 255, 255})
	img.Set(1, 1, color.RGBA{255, 255, 0, 255})
	return img
}
