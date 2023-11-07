package boids

import "image/color"

const (
	screenWidth  = 1080
	screenHeight = 640
	numClasses   = 5   // maximum of 6 classes
	numEach      = 500 // careful when scaling
	minSpeed     = 0.3
	maxSpeed     = 2.8
	maxForce     = 0.8
	colorblind   = true
	drawQuadtree = true
)

var perceptionMap = map[int]float64{
	0: 50.0, // alignment
	1: 50.0, // cohesion
	2: 50.0, // separation
}

var quadtreeColor = color.NRGBA{
	R: 255, G: 255, B: 255, A: 30,
}
