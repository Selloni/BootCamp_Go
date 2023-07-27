package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
)

func main() {
	file, err := os.Create("amazing_logo.png")
	if err != nil {
		log.Fatal("don`t creat file .png")
	}
	img := image.NewRGBA(image.Rect(0, 0, 300, 300))
	fonColor := color.RGBA{0, 50, 100, 200}
	draw.Draw(img, img.Bounds(), &image.Uniform{fonColor}, image.Pt(0, 0), draw.Src)
	lineColor := color.RGBA{200, 30, 30, 255}
	for x := 50; x < 200; x++ {
		y1 := x/3 + 25
		y2 := x/3 + 50
		y3 := x/3 + 75
		img.Set(x, y1, lineColor)
		img.Set(x, y2, lineColor)
		img.Set(x, y3, lineColor)
		y4 := x/3 + 125
		y5 := x/3 + 150
		y6 := x/3 + 175
		img.Set(x+75, y4, lineColor)
		img.Set(x+75, y5, lineColor)
		img.Set(x+75, y6, lineColor)
		for y := 250; y < 300; y++ {
			img.Set(x+50, y, color.RGBA{0, 80, 8, 88})
		}
	}
	png.Encode(file, img)
}
