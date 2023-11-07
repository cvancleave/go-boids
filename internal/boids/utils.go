package boids

import (
	"fmt"
	"image"
	"log"
	"math/rand"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

func getImage(item int) *ebiten.Image {

	// load image
	file, err := os.Open(fmt.Sprintf("./assets/%d.png", item))
	if err != nil {
		log.Fatal(err)
	}

	// decode image from file
	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	// create ebiten version of image
	return ebiten.NewImageFromImage(img)
}

func randomVelocity(max float64) float64 {
	return rand.Float64()*max - max/2
}

func modulo(v float64, m int) float64 {
	if v > float64(m) {
		return v - float64(m)
	} else if v < 0 {
		return v + float64(m)
	}
	return v
}
