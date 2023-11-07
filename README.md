# go-boids

This is an implementation of boids using the Ebiten 2d game engine in Go, and uses a quadtree for optimization.
Using the quadtree makes the time complexity of calculating boid trajectories `O(log n)` instead of `O(n^2)`.

## Running the application

Modify any parameters in `internal/boids/params.go` to change behavior and boid counts.
Run the application with `go run cmd/boids/main.go`.
Move the up/down/left/right arrow keys to modify alignment, cohesion, and separation values.