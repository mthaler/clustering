package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"
)

func main() {
	faithful := readCsvFile("faithful.csv")
	var points []point
	for _, v := range faithful {
		// only add a point if this is not a line with labels
		if v[0] != "eruptions" {
			x, err := strconv.ParseFloat(v[0], 64)
			if err != nil {
				log.Fatal("Could not parse number of eruptions")
			}
			y, err := strconv.ParseFloat(v[1], 64)
			if err != nil {
				log.Fatal("Could not parse waiting time")
			}
			p := point{x: x, y: y}
			points = append(points, p)
		}
	}
	maxErruptions := 0.0
	for _, p := range points {
		if p.x > maxErruptions {
			maxErruptions = p.x
		}
	}
	maxWaiting := 0.0
	for _, p := range points {
		if p.y > maxWaiting {
			maxWaiting = p.y
		}
	}
	centroids := generateCentroids(2, maxErruptions, maxWaiting)
	fmt.Printf("%+v\n", centroids)
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
func generateCentroids(k int, maxX float64, maxY float64) []centroid {
	var c []centroid
	c = make([]centroid, k)
	for i := 0; i < k; i++ {
		c[i] = centroid{x: rand.Float64() * maxX, y: rand.Float64() * maxY}
	}
	return c
}
