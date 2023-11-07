package boids

import (
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

type boid struct {
	// color
	color int

	// movement
	position     vector
	velocity     vector
	acceleration vector

	// image
	image        *ebiten.Image
	imageOptions *ebiten.DrawImageOptions
	imageWidth   int
	imageHeight  int
}

func NewBoid(color int) *boid {
	image := getImage(color)
	return &boid{
		// info
		color: color,
		// movement
		position: vector{
			X: rand.Float64() * float64(screenWidth),
			Y: rand.Float64() * float64(screenHeight),
		},
		velocity: vector{
			X: randomVelocity(maxSpeed),
			Y: randomVelocity(maxSpeed),
		},
		// image
		image:        image,
		imageOptions: &ebiten.DrawImageOptions{},
		imageWidth:   image.Bounds().Dx(),
		imageHeight:  image.Bounds().Dy(),
	}
}

func (b *boid) draw(screen *ebiten.Image) {
	b.imageOptions.GeoM.Reset()
	b.imageOptions.GeoM.Rotate(-1*math.Atan2(b.velocity.Y*-1, b.velocity.X) + math.Pi/2)
	b.imageOptions.GeoM.Translate(b.position.X, b.position.Y)
	screen.DrawImage(b.image, b.imageOptions)
}

func (b *boid) update(boids []*boid) {
	b.rules(boids)
	b.movement()
	b.checkSpeed(minSpeed, maxSpeed)
	b.checkPosition(screenWidth, screenHeight)
}

func (b *boid) rules(others []*boid) {

	alignment := vector{}
	alignmentTotal := 0
	cohesion := vector{}
	cohesionTotal := 0
	separation := vector{}
	separationTotal := 0

	for _, other := range others {

		// ignore other colors if not colorblind
		if !colorblind && b.color != other.color {
			continue
		}

		d := b.position.distance(other.position)
		if b != other {
			// alignment
			if d < perceptionMap[0] {
				alignment.add(other.velocity)
				alignmentTotal++
			}
			// cohesion
			if d < perceptionMap[1] {
				cohesion.add(other.position)
				cohesionTotal++
			}
			// separation
			if d < perceptionMap[2] {
				diff := b.position
				diff.subtract(other.position)
				diff.divide(d)
				separation.add(diff)
				separationTotal++
			}
		}
	}

	// alignment
	if alignmentTotal > 0 {
		alignment.divide(float64(alignmentTotal))
		alignment.setMagnitude(maxSpeed)
		alignment.subtract(b.velocity)
		alignment.limit(maxForce)
	}

	// cohesion
	if cohesionTotal > 0 {
		cohesion.divide(float64(cohesionTotal))
		cohesion.subtract(b.position)
		cohesion.setMagnitude(maxSpeed)
		cohesion.subtract(b.velocity)
		cohesion.setMagnitude(maxForce)
	}

	// separation
	if separationTotal > 0 {
		separation.divide(float64(separationTotal))
		separation.setMagnitude(maxSpeed)
		separation.subtract(b.velocity)
		separation.setMagnitude(maxForce)
	}

	b.acceleration.add(alignment)
	b.acceleration.add(cohesion)
	b.acceleration.add(separation)
	b.acceleration.divide(3)
}

func (b *boid) movement() {
	b.position.add(b.velocity)
	b.velocity.add(b.acceleration)
	b.velocity.limit(maxSpeed)
	b.acceleration.multiply(0.0)
}

func (b *boid) checkSpeed(min, max float64) {
	d := math.Sqrt(math.Pow(b.velocity.X, 2) + math.Pow(b.velocity.Y, 2))
	if d < min {
		r := d / min
		b.velocity.X /= r
		b.velocity.Y /= r
	} else if d > max {
		r := d / max
		b.velocity.X /= r
		b.velocity.Y /= r
	}
}

// wraparound display
func (b *boid) checkPosition(width, height int) {
	b.position.X = modulo(b.position.X, width)
	b.position.Y = modulo(b.position.Y, height)
}
