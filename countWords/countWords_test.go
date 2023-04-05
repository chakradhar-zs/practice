package main

import "testing"

func Test_CountWords(t *testing.T) {
	tc := []struct {
		in1  string
		want int
	}{
		{"Hello Gophers", 2},
		{"Welcome to Go Language", 4},
		{"It is a sample message to you", 7},
		{"", 0},
		{"Hi", 1},
	}

	for _, v := range tc {
		op := CountWords(v.in1)
		if op != v.want {
			t.Errorf("GreetEmployee(%s)==%d, want=%d", v.in1, op, v.want)
		}

	}
}
