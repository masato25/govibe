package main

import (
	"fmt"
	"time"

	"github.com/go-vgo/robotgo"
	"github.com/kbinani/screenshot"
	"gocv.io/x/gocv"
)

func main() {
	numDisplays := screenshot.NumActiveDisplays()
	fmt.Printf("偵測到 %d 個螢幕\n", numDisplays)
	allowImg := gocv.IMRead("allow_button.png", gocv.IMReadColor)
	if allowImg.Empty() {
		panic("找不到 allow_button.png")
	}

	for {
		bounds := screenshot.GetDisplayBounds(0)
		img, _ := screenshot.CaptureRect(bounds)
		mat, _ := gocv.ImageToMatRGB(img)

		result := gocv.NewMat()
		gocv.MatchTemplate(mat, allowImg, &result, gocv.TmCcoeffNormed, gocv.NewMat())

		_, maxVal, _, maxLoc := gocv.MinMaxLoc(result)
		if maxVal > 0.8 {
			fmt.Printf("找到 Allow 按鈕! 相似度: %.2f\n", maxVal)
			x := bounds.Min.X + maxLoc.X + allowImg.Cols()/2
			y := bounds.Min.Y + maxLoc.Y + allowImg.Rows()/2
			robotgo.MoveMouse(x, y)
			robotgo.MouseClick("left")
			time.Sleep(1 * time.Second)
		}

		mat.Close()
		result.Close()
		time.Sleep(500 * time.Millisecond)
	}
}
