package boids

import "math"

type vector struct {
	X, Y float64
}

func (v *vector) add(v2 vector) {
	v.X += v2.X
	v.Y += v2.Y
}

func (v *vector) subtract(v2 vector) {
	v.X -= v2.X
	v.Y -= v2.Y
}

func (v *vector) divide(z float64) {
	v.X /= z
	v.Y /= z
}

func (v *vector) multiply(z float64) {
	v.X *= z
	v.Y *= z
}

func (v *vector) distance(v2 vector) float64 {
	return math.Sqrt(math.Pow(v2.X-v.X, 2) + math.Pow(v2.Y-v.Y, 2))
}

func (v *vector) limit(max float64) {
	magSq := v.magnitudeSquared()
	if magSq > max*max {
		v.divide(math.Sqrt(magSq))
		v.multiply(max)
	}
}

func (v *vector) normalize() {
	mag := math.Sqrt(v.X*v.X + v.Y*v.Y)
	v.X /= mag
	v.Y /= mag
}

func (v *vector) setMagnitude(z float64) {
	v.normalize()
	v.X *= z
	v.Y *= z
}

func (v *vector) magnitudeSquared() float64 {
	return v.X*v.X + v.Y*v.Y
}
