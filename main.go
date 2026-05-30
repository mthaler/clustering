package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Hello, World!")
}

// calculates the Euclidan distance between two points
func distance(p1 point, p2 point) float64 {
	return math.Sqrt((p1.x-p2.x)*(p1.x-p2.x) + (p1.y-p2.y)*(p1.y-p2.y))
}
