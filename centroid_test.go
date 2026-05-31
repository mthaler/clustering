package main

import "testing"

func TestAverage(t *testing.T) {
	c := centroid{x: 1, y: 1}
	c.average([]point{{x: 1, y: 1}, {x: 2, y: 2}, {x: 3, y: 3}})
	if c.x != 2 {
		t.Errorf("x=%g, expected 2\n", c.x)
	}
	if c.y != 2 {
		t.Errorf("y=%g, expected 2\n", c.y)
	}
}
