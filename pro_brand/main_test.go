package main

import (
	"reflect"
	"testing"
)

func TestCreateProduct(t *testing.T) {
	testcases := []struct {
		input1 Product
		input2 string
		output int64
	}{

		{Product{3, "sneaker shoes", "stylish", 1000, 3, "shoes", 4, "Available"}, "Nike", 1},
		{Product{4, "Rolex", "useful", 50000, 1, "wristwatch", 5, "Discontinued"}, "Titan", 1},
	}

	for _, v := range testcases {
		op := CreateProduct(v.input1, v.input2)

		if !reflect.DeepEqual(op, v.output) {
			t.Error("Inserted Data incorrectly")
		}
	}
}

func TestRetrieveProducts(t *testing.T) {
	testcases := []struct {
		input   int
		output1 Product
		outpu2  string
	}{

		{3, Product{3, "sneaker shoes", "stylish", 1000, 3, "shoes", 4, "Available"}, "Nike"},
		{4, Product{4, "Rolex", "useful", 50000, 1, "wristwatch", 5, "Discontinued"}, "Titan"},
	}

	for _, v := range testcases {
		op1, op2 := RetrieveProducts(v.input)

		if !reflect.DeepEqual(op1, v.output1) {
			t.Errorf("Retrieved Product Data is incorrect got %v, expected %v", op1, v.output1)
		}

		if op2 != v.outpu2 {
			t.Errorf("Retrieved  Brand name is incorrect got %v,%v, expected %v,%v", op1, op2, v.output1, v.outpu2)
		}
	}
}
