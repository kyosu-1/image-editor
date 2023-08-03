package filters

import (
	"image"
	"image/color"
)

// ApplyGrayscale 強度をパラメータとして受け取り、画像にグレースケールフィルターを適用します。
// 強度は0から1の範囲で、0が元の画像、1が完全なグレースケールを意味します。
func ApplyGrayscale(img image.Image, strength float64) image.Image {
	if strength < 0 {
		strength = 0
	} else if strength > 1 {
		strength = 1
	}

	bounds := img.Bounds()
	grayImg := image.NewGray(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			oldColor := color.RGBAModel.Convert(img.At(x, y)).(color.RGBA)
			grayValue := 0.299*float64(oldColor.R) + 0.587*float64(oldColor.G) + 0.114*float64(oldColor.B)
			gray := uint8(strength*grayValue + (1-strength)*float64(oldColor.R))
			grayImg.Set(x, y, color.Gray{Y: gray})
		}
	}

	return grayImg
}
