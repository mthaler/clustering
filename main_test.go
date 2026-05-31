package main

import "testing"

func TestDistance(t *testing.T) {
	p1 := point{x: 1, y: 0}
	p2 := point{x: 2, y: 0}
	d := distance(p1, p2)
	if d != 1.0 {
		t.Errorf("Distance is %g, should be 1", d)
	}
}
