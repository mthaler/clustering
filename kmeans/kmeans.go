package kmeans	

import (
    "math/rand/v2"
)

 func centroide() Point {
	x := rand.Float64()
	y := rand.Float64()
	return Point{X: x, Y: y}
 }