package boids

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Canvas struct {
	quadtree *quadtree
	allBoids []*boid

	// ui
	selectedRule int
}

func StartCanvas() {

	boundary := NewRectangle(0, 0, screenWidth, screenHeight)
	quadtree := NewQuadtree(boundary, 8)

	allBoids := []*boid{}
	for i := 0; i < numClasses; i++ {
		for j := 0; j < numEach; j++ {
			b := NewBoid(i)
			allBoids = append(allBoids, b)
			quadtree.insert(b)
		}
	}

	canvas := &Canvas{
		quadtree: quadtree,
		allBoids: allBoids,
	}

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowResizingMode(2) // 2 = resizable

	if err := ebiten.RunGame(canvas); err != nil {
		panic(err)
	}
}

// mandatory for interface implementation
func (c *Canvas) Update() error {
	boundary := NewRectangle(0, 0, screenWidth, screenHeight)
	quadtree := NewQuadtree(boundary, 8)

	for _, b := range c.allBoids {
		quadtree.insert(b)
	}

	c.quadtree = quadtree
	c.quadtree.update()

	c.handleKeypress()
	return nil
}

// mandatory for interface implementation
func (c *Canvas) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)

	// draw quadtree and boids
	c.quadtree.draw(screen)

	ebitenutil.DebugPrint(screen,
		fmt.Sprintf(
			"FPS: %0.2f\nAlignment %0.2f\nCohesion %0.2f\nSeparation %0.2f\nTotal: %d",
			ebiten.ActualTPS(), perceptionMap[0], perceptionMap[1], perceptionMap[2], len(c.allBoids),
		),
	)
}

// mandatory for interface implementation
func (c *Canvas) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

// up down left right to modify align/cohesion/separation values
func (c *Canvas) handleKeypress() {
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		c.selectedRule++
		c.selectedRule %= 3
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		c.selectedRule += 2
		c.selectedRule %= 3
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) {
		perceptionMap[c.selectedRule] += 5
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) {
		perceptionMap[c.selectedRule] -= 5
	}
}
