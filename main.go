package main

import (
	"image/jpeg"
	"log"
	"os"

	"github.com/nfnt/resize"
)

func main() {
	file, err := os.Open("test.jpg")
	if err != nil {
		log.Fatal(err)
	}

	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	method0 := resize.NearestNeighbor
	method1 := resize.Bilinear
	method2 := resize.Bicubic
	method3 := resize.MitchellNetravali
	method4 := resize.Lanczos2
	method5 := resize.Lanczos3

	m0 := resize.Resize(1000, 0, img, method0)
	out, err := os.Create("resize0.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	jpeg.Encode(out, m0, nil)

	m1 := resize.Resize(1000, 0, img, method1)
	out, err = os.Create("resize1.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	jpeg.Encode(out, m1, nil)

	m2 := resize.Resize(1000, 0, img, method2)
	out, err = os.Create("resize2.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	jpeg.Encode(out, m2, nil)

	m3 := resize.Resize(1000, 0, img, method3)
	out, err = os.Create("resize3.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	jpeg.Encode(out, m3, nil)

	m4 := resize.Resize(1000, 0, img, method4)
	out, err = os.Create("resize4.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	jpeg.Encode(out, m4, nil)

	m5 := resize.Resize(1000, 0, img, method5)
	out, err = os.Create("resize5.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	jpeg.Encode(out, m5, nil)
}
