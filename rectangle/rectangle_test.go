package main

import "testing"

func TestPerimeter(t *testing.T) {
	tc := []struct {
		length, breadth, peri int
	}{
		{2, 4, 12},
		{3, 4, 14},
		{4, 5, 18},
		{2, 0, 4},
		{2, -1, 2},
	}
	for _, v := range tc {
		op := perimeter(v.length, v.breadth)
		if op != v.peri {
			t.Errorf("perometer(%d,%d)==%d, want = %d", v.length, v.breadth, op, v.peri)
		}
	}
}

func TestArea(t *testing.T) {
	tc := []struct {
		length, breadth, area int
	}{
		{2, 4, 8},
		{3, 4, 12},
		{4, 5, 20},
		{2, 0, 0},
		{2, -1, -2},
	}
	for _, v := range tc {
		op := area(v.length, v.breadth)
		if op != v.area {
			t.Errorf("perometer(%d,%d)==%d, want = %d", v.length, v.breadth, op, v.area)
		}
	}
}
