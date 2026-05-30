package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
)

func main() {
	fmt.Println("Hello, World!")
}

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
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
