package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	fmt.Println("Hello, World!")
}

// calculates the Euclidan distance between two points
func distance(p1 point, p2 point) float64 {
	return math.Sqrt((p1.x-p2.x)*(p1.x-p2.x) + (p1.y-p2.y)*(p1.y-p2.y))
}

// generates k centroids
func generateCentroids(k int) []centroid {
	var c []centroid
	c = make([]centroid, k)
	for i := 0; i < k; i++ {
		c[i] = centroid{x: rand.Float64(), y: rand.Float64()}
	}
	return c
}
