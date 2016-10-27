package main

import (
	"image"
	"log"
	"os"
	"image/png"
	"image/draw"
)

var (
	zero = image.Point{0, 0}
)

func read(s string) image.Image {
	file, err := os.Open(s); if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatalf("%s: %v", file, err)
	}
	return img
}

func write(path string, img image.Image) {
	var (
		o *os.File
		err error
	)

	if path != "-" {
		o, err = os.Create(path); if err != nil {
			log.Fatal(err)
		}
	}

	err = png.Encode(o, img)
	if err != nil {
		log.Fatal(err)
	}
}

func stitch(images []image.Image) image.Image {
	var x, y int
	for _, img := range images {
		x = img.Bounds().Max.X
		y += img.Bounds().Max.Y
	}
	dst := image.NewRGBA(image.Rect(0, 0, x, y))

	x, y = 0, 0
	for _, img := range images {
		draw.Draw(dst, img.Bounds().Add(image.Point{x, y}), img, zero, draw.Src)
		y += img.Bounds().Max.Y
	}
	return dst
}

