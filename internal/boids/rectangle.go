package boids

type rectangle struct {
	X, Y, W, H float64
	L, R, T, B float64
}

func NewRectangle(x, y, w, h float64) *rectangle {
	return &rectangle{
		X: x, Y: y, W: w, H: h,
		L: x,
		R: x + w,
		T: y,
		B: y + h,
	}
}

func (r *rectangle) contains(b *boid) bool {
	return (r.L <= b.position.X && b.position.X <= r.R && r.T <= b.position.Y && b.position.Y <= r.B)
}

func (r *rectangle) subdivide(quadrant string) *rectangle {
	switch quadrant {
	case "ne":
		return NewRectangle(r.X+r.W/2, r.Y, r.W/2, r.H/2)
	case "nw":
		return NewRectangle(r.X, r.Y, r.W/2, r.H/2)
	case "se":
		return NewRectangle(r.X+r.W/2, r.Y+r.H/2, r.W/2, r.H/2)
	case "sw":
		return NewRectangle(r.X, r.Y+r.H/2, r.W/2, r.H/2)
	}
	return nil
}
