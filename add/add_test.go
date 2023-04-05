package main

import "testing"

func TestAdd(t *testing.T) {

	tc := []struct {
		in1, in2, exout int
	}{
		{1, 2, 3},
		{2, 3, 5},
		{3, -1, 2},
	}

	for _, v := range tc {
		op := add(v.in1, v.in2)
		if op != v.exout {
			t.Errorf("add(%d)==%d but we want %d", v.in1, v.in2, v.exout)
		}
	}
}

func TestDivide(t *testing.T) {

	tc := []struct {
		in1, in2, exout int
	}{
		{4, 2, 2},
		{5, 2, 2},
		{3, 1, 3},
	}

	for _, v := range tc {
		op := divide(v.in1, v.in2)
		if op != v.exout {
			t.Errorf("divide(%d)==%d but we want %d", v.in1, v.in2, v.exout)
		}
	}
}
