package main

import (
	"fmt"
	"image"
	"image/png"
	"os"
)

func combineImages(images []image.Image) image.Image {
	width := images[0].Bounds().Dx()
	height := images[0].Bounds().Dy() * len(images)
	canvas := image.NewRGBA(image.Rect(0, 0, width, height))
	var y int
	for _, img := range images {
		for i := 0; i < img.Bounds().Dy(); i++ {
			for j := 0; j < img.Bounds().Dx(); j++ {
				canvas.Set(j, y+i, img.At(j, i))
			}
		}
		y += img.Bounds().Dy()
	}
	return canvas
}

func main() {
	images := make([]image.Image, 3)
	for i := 0; i < 3; i++ {
		f, _ := os.Open(fmt.Sprintf("image%d.png", i+1))
		img, _ := png.Decode(f)
		images[i] = img
	}
	combined := combineImages(images)
	f, _ := os.Create("combined.png")
	png.Encode(f, combined)
}
