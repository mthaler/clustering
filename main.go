package main

import (
	"encoding/csv"
	"fmt"
	"image/color"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
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

	var points2 plotter.XYs
	points2 = make(plotter.XYs, len(points))
	for _, p := range points {
		points2 = append(points2, plotter.XY{X: p.x, Y: p.y})
	}

	// Create a new plot
	pl := plot.New()
	scatter, err := plotter.NewScatter(points2)
	if err != nil {
		panic(err)
	}
	scatter.GlyphStyle.Color = color.RGBA{R: 255, A: 255}
	scatter.GlyphStyle.Radius = vg.Points(4)
	// Add the scatter plot to the plot and set the axes labels
	pl.Add(scatter)
	pl.Title.Text = "Old Faithful"
	pl.X.Label.Text = "Eruptions"
	pl.Y.Label.Text = "Waiting"
	// Save the plot to a PNG file
	if err := pl.Save(4*vg.Inch, 4*vg.Inch, "faithful2.png"); err != nil {
		panic(err)
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
