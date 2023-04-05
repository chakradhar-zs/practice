package main

import "testing"

func Test_GreetEmployee(t *testing.T) {
	tc := []struct {
		f, l string
		want string
	}{
		{"Chakradhar", "Ghali", "Hello, Chakradhar Ghali"},
		{"Shubham", "Munjani", "Hello, Shubham Munjani"},
		{"Piyush", "Singh", "Hello, Piyush Singh"},
		{"Maneesh", "", "Hello, Maneesh"},
	}

	for _, v := range tc {
		op := GreetEmployee(v.f, v.l)
		if op != v.want {
			t.Errorf("GreetEmployee(%s,%s)==%s, want=%s", v.f, v.l, op, v.want)
		}

	}
}

func Test_CalculateAge(t *testing.T) {
	tc := []struct {
		dob  string
		want int
	}{
		{"15-04-2002", 21},
		{"12-06-2005", 18},
		{"21-8-1998", 25},
		{"30-1-2020", 3},
		{"15-12-2000", 23},
	}
	for _, v := range tc {
		op := CalculateAge(v.dob)
		if op != v.want {
			t.Errorf("CalculateAge(%s)==%d, want=%d", v.dob, op, v.want)
		}

	}
}
