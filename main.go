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
	continueImg := gocv.IMRead("continue_button.png", gocv.IMReadColor)
	if continueImg.Empty() {
		panic("找不到 continue_button.png")
	}

	for {
		for i := 0; i < numDisplays; i++ {
			bounds := screenshot.GetDisplayBounds(i)
			img, _ := screenshot.CaptureRect(bounds)
			mat, _ := gocv.ImageToMatRGB(img)

			result := gocv.NewMat()
			gocv.MatchTemplate(mat, allowImg, &result, gocv.TmCcoeffNormed, gocv.NewMat())

			_, maxVal, _, maxLoc := gocv.MinMaxLoc(result)
			if maxVal > 0.9 {
				fmt.Printf("在螢幕 %d 找到 Allow 按鈕! 相似度: %.2f\n", i, maxVal)
				x := bounds.Min.X + maxLoc.X + allowImg.Cols()/2
				y := bounds.Min.Y + maxLoc.Y + allowImg.Rows()/2
				robotgo.Move(x, y)
				robotgo.Click("left")
				time.Sleep(1 * time.Second)
			}

			gocv.MatchTemplate(mat, continueImg, &result, gocv.TmCcoeffNormed, gocv.NewMat())

			_, maxVal, _, maxLoc = gocv.MinMaxLoc(result)
			if maxVal > 0.98 {
				fmt.Printf("在螢幕 %d 找到 Continue 按鈕! 相似度: %.2f\n", i, maxVal)
				x := bounds.Min.X + maxLoc.X + continueImg.Cols()/2
				y := bounds.Min.Y + maxLoc.Y + continueImg.Rows()/2
				robotgo.Move(x, y)
				robotgo.Click("left")
				time.Sleep(1 * time.Second)
			}

			mat.Close()
			result.Close()
		}
		time.Sleep(500 * time.Millisecond)
	}
}
