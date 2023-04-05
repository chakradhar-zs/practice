package main

import "testing"

func TestString(t *testing.T) {
	testcases := []struct {
		input  Person
		output string
	}{
		{Person{"John", 25}, "John, 25"},
		{Person{"Chakradhar", 20}, "Chakradhar, 20"},
		{Person{}, ""},
		{Person{"Shubham", 21}, "Shubham, 21"},
		{Person{"", 20}, ""},
		{Person{"Raj", 0}, ""},
	}

	for _, val := range testcases {
		op := val.input.String()

		if op != val.output {
			t.Errorf("Unexpected Output %s , want %s", op, val.output)
		}
	}
}
