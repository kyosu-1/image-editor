package resize

import (
	"image"
)

// Resize は画像を指定された幅と高さにリサイズします。
func Resize(img image.Image, width, height int) image.Image {
	newImg := image.NewRGBA(image.Rect(0, 0, width, height))
	bounds := img.Bounds()
	xRatio := float64(bounds.Dx()) / float64(width)
	yRatio := float64(bounds.Dy()) / float64(height)

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			srcX := int(float64(x) * xRatio)
			srcY := int(float64(y) * yRatio)
			newImg.Set(x, y, img.At(srcX+bounds.Min.X, srcY+bounds.Min.Y))
		}
	}

	return newImg
}
