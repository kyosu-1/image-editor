package rotate

import (
	"image"
	"math"
)

// Rotate は画像を指定された角度で回転します。
func Rotate(img image.Image, angle float64) image.Image {
	bounds := img.Bounds()
	width, height := bounds.Dx(), bounds.Dy()
	newImg := image.NewRGBA(image.Rect(0, 0, width, height))
	cx, cy := width/2, height/2

	radians := angle * math.Pi / 180

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			// 回転変換
			srcX := int(float64(cx) + (float64(x-cx) * math.Cos(radians)) + (float64(y-cy) * math.Sin(radians)))
			srcY := int(float64(cy) - (float64(x-cx) * math.Sin(radians)) + (float64(y-cy) * math.Cos(radians)))
			if srcX >= 0 && srcX < width && srcY >= 0 && srcY < height {
				newImg.Set(x, y, img.At(srcX, srcY))
			}
		}
	}

	return newImg
}

// FlipHorizontal は画像を水平方向に反転します。
func FlipHorizontal(img image.Image) image.Image {
	bounds := img.Bounds()
	width := bounds.Dx()
	newImg := image.NewRGBA(bounds)

	for x := 0; x < width; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			newImg.Set(width-1-x, y, img.At(x, y))
		}
	}

	return newImg
}
