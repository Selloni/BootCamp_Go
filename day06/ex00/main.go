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
	for x := 20; x < 300; x++ {
		y := x/3 + 15
		img.Set(x, y, lineColor)
	}

	png.Encode(file, img)
}
