package main

import (
	"fmt"
	"image"
	"image/color"
	_ "image/png"
	"os"
	"sync"
)

// turns the colors in the image into int values that the then will be turned into a playable maze
// 0 - empty space
// 1 - wall
// 2 - starting point
// 3 - danger zone
// 5 - exit point
// Maybe more types of cells will be added later
func colorToCode(pixelColor color.Color) uint8 {
	r, g, b, a := pixelColor.RGBA()
	if r == 65535 {
		return 3
	}
	if g == 65535 {
		return 2
	}
	if b == 65535 {
		return 5
	}
	if r == 0 && g == 0 && b == 0 && a == 65535 {
		return 1
	} else {
		return 0
	}
}

func processRow(img image.Image, y int, imgArray [][]uint8, wg *sync.WaitGroup, mu *sync.Mutex, imgWidth int) {
	defer wg.Done()

	var yArr []uint8

	for x := 0; x < imgWidth; x++ {
		yArr = append(yArr, colorToCode(img.At(x, y)))
	}

	mu.Lock()
	imgArray[y] = yArr
	mu.Unlock()
}

// is slightly worse on small images than the non parallel version,
// taking in consideration the size of the images i pretend to use it might be
// beter to remove go functions but the difference is only about 36 microseconds and this
// way looks waaaay more cooler
func TransformImage(levelName string) *MapTensor {
	imgFile, err := os.Open(fmt.Sprintf("./levels/%s", levelName))
	if err != nil {
		panic(err.Error())
	}

	img, _, err := image.Decode(imgFile)
	if err != nil {
		panic(err.Error())
	}
	imgWidth := img.Bounds().Size().X
	imgHeight := img.Bounds().Size().Y
	imgArray := make([][]uint8, imgHeight)

	var wg sync.WaitGroup

	var mu sync.Mutex

	for y := 0; y < imgHeight; y++ {
		wg.Add(1)
		go processRow(img, y, imgArray, &wg, &mu, imgWidth)
	}
	wg.Wait()

	return &MapTensor{
		Width:  uint8(imgWidth),
		Height: uint8(imgHeight),
		Tensor: imgArray,
	}
}
