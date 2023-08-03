package main

import (
	"syscall/js"
	"bytes"
	"image/jpeg"
	"image"

	"github.com/kyosu-1/image-editor/pkg/filters"
	"github.com/kyosu-1/image-editor/pkg/resize"
	"github.com/kyosu-1/image-editor/pkg/rotate"
)

func loadImage(p js.Value) (image.Image, error) {
	// JavaScriptから渡されたバイナリデータを画像として読み込みます。
	jsArray := js.Global().Get("Uint8Array").New(p)
	byteData := make([]byte, jsArray.Get("length").Int())
	js.CopyBytesToGo(byteData, jsArray)
	reader := bytes.NewReader(byteData)
	return jpeg.Decode(reader)
}

func imageToBytes(img image.Image) []byte {
	var buffer bytes.Buffer
	err := jpeg.Encode(&buffer, img, nil)
	if err != nil {
		return nil
	}
	return buffer.Bytes()
}

// byte配列をUint8Arrayに変換する関数
func bytesToUint8Array(bytes []byte) js.Value {
	uint8Array := js.Global().Get("Uint8Array").New(len(bytes))
	js.CopyBytesToJS(uint8Array, bytes)
	return uint8Array
}

func applyGrayscaleWrapper() js.Func {
	return js.FuncOf(func(this js.Value, p []js.Value) interface{} {
		img, err := loadImage(p[0])
		if err != nil {
			return nil
		}
		strength := p[1].Float()
		newImg := filters.ApplyGrayscale(img, strength) // 強度を適用する処理を追加
		resultBytes := imageToBytes(newImg)

		return bytesToUint8Array(resultBytes)
	})
}

func resizeWrapper() js.Func {
	return js.FuncOf(func(this js.Value, p []js.Value) interface{} {
		img, err := loadImage(p[0])
		if err != nil {
			return nil
		}
		newWidth := p[1].Int()
		newHeight := p[2].Int()
		newImg := resize.Resize(img, newWidth, newHeight) // 新しいサイズを適用
		resultBytes := imageToBytes(newImg)
		return bytesToUint8Array(resultBytes)
	})
}

func rotateWrapper() js.Func {
	return js.FuncOf(func(this js.Value, p []js.Value) interface{} {
		img, err := loadImage(p[0])
		if err != nil {
			return nil
		}
		angle := p[1].Float()
		newImg := rotate.Rotate(img, angle) // 角度を適用
		resultBytes := imageToBytes(newImg)
		return bytesToUint8Array(resultBytes)
	})
}

func flipHorizontalWrapper() js.Func {
	return js.FuncOf(func(this js.Value, p []js.Value) interface{} {
		img, err := loadImage(p[0])
		if err != nil {
			return nil
		}
		newImg := rotate.FlipHorizontal(img)
		resultBytes := imageToBytes(newImg)
		return bytesToUint8Array(resultBytes)
	})
}

func main() {
	c := make(chan struct{}, 0)

	js.Global().Set("applyGrayscale", applyGrayscaleWrapper())
	js.Global().Set("resize", resizeWrapper())
	js.Global().Set("rotate", rotateWrapper())
	js.Global().Set("flipHorizontal", flipHorizontalWrapper())

	<-c
}
