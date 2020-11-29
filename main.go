package main

import (
	"image"
	"image/color"
	"image/gif"
	"math"
	"os"
)

type Circle struct {
	X, Y, R float64
}

func (c *Circle) Brightness(x, y float64) uint8 {
	var dx, dy float64 = c.X - x, c.Y - y
	d := math.Sqrt(dx*dx+dy*dy) / c.R
	if d > 1 {
		return 0
	} else {
		return 255
	}
}

func main() {
	var w, h int = 1000, 1000
	unit := 0
	white := color.RGBA{255, 255, 255, 0xff}
	black := color.RGBA{0, 0, 0, 0xff}

	var palette = []color.Color{
		color.RGBA{0x00, 0x00, 0x00, 0xff},
		color.RGBA{0xff, 0xff, 0xff, 0xff},
	}

	var images []*image.Paletted
	var delays []int
	steps := 20
	for step := 0; step < steps; step++ {
		img := image.NewPaletted(image.Rect(0, 0, w, h), palette)
		images = append(images, img)
		delays = append(delays, 0)

		unit += step

		for i := 25; i < w; i += 100 {
			for j := 25; j < h; j += 100 {
				printSquare(i, j, unit/2, img, white)
				printSquare(i+unit, j+unit, unit/2, img, white)

				printSquare(i+unit, j, unit/2, img, black)
				printSquare(i, j+unit, unit/2, img, black)
			}
		}
	}

	f, _ := os.OpenFile("rgb.gif", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	gif.EncodeAll(f, &gif.GIF{
		Image: images,
		Delay: delays,
	})
}

func printSquare(x, y, radius int, img *image.Paletted, color color.RGBA) {
	for i := x - radius; i < x+radius; i++ {
		for j := y - radius; j < y+radius; j++ {
			img.Set(i, j, color)
		}
	}
}
