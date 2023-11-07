package boids

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	eVector "github.com/hajimehoshi/ebiten/v2/vector"
)

type quadtree struct {
	boundary       *rectangle
	capacity       int
	divided        bool
	boids          []*boid
	ne, nw, se, sw *quadtree
}

func NewQuadtree(boundary *rectangle, capacity int) *quadtree {
	if capacity == 0 {
		capacity = 8
	}
	return &quadtree{
		boundary: boundary,
		capacity: capacity,
	}
}

func (q *quadtree) Subdivide() {
	q.ne = NewQuadtree(q.boundary.subdivide("ne"), q.capacity)
	q.nw = NewQuadtree(q.boundary.subdivide("nw"), q.capacity)
	q.se = NewQuadtree(q.boundary.subdivide("se"), q.capacity)
	q.sw = NewQuadtree(q.boundary.subdivide("sw"), q.capacity)
	q.divided = true

	// Move points to children.
	// This improves performance by placing points in the smallest available rectangle.
	for _, b := range q.boids {
		inserted := q.ne.insert(b) ||
			q.nw.insert(b) ||
			q.se.insert(b) ||
			q.sw.insert(b)

		if !inserted {
			fmt.Println("error: capacity must be greater than 0")
		}
	}

	q.boids = nil
}

func (q *quadtree) insert(b *boid) bool {

	if !q.boundary.contains(b) {
		return false
	}

	if !q.divided {
		if len(q.boids) < q.capacity {
			q.boids = append(q.boids, b)
			return true
		}

		q.Subdivide()
	}

	return (q.ne.insert(b) ||
		q.nw.insert(b) ||
		q.se.insert(b) ||
		q.sw.insert(b))
}

func (q *quadtree) draw(screen *ebiten.Image) {

	for _, b := range q.boids {
		b.draw(screen)
	}

	if drawQuadtree {

		eVector.StrokeRect(
			screen,
			float32(q.boundary.X),
			float32(q.boundary.Y),
			float32(q.boundary.W),
			float32(q.boundary.H),
			1, quadtreeColor, true,
		)
	}

	if q.divided {
		q.ne.draw(screen)
		q.nw.draw(screen)
		q.se.draw(screen)
		q.sw.draw(screen)
	}
}

func (q *quadtree) update() {

	for _, b := range q.boids {
		b.update(q.boids)
	}

	if q.divided {
		q.ne.update()
		q.nw.update()
		q.se.update()
		q.sw.update()
	}
}
