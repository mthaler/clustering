package main

type centroid struct {
	x float64
	y float64
}

func (c centroid) average(points []point) {
	if len(points) == 0 {
		return
	}

	x := 0.0
	y := 0.0
	for _, p := range points {
		x += p.x
		y += p.y
	}
	x /= float64(len(points))
	y /= float64(len(points))
	c.x = x
	c.y = y
}
