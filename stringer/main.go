package main

import "fmt"

type Person struct {
	name string
	age  int
}

//This method implements Stringer interface and use String method to print name and age in customized format
func (p Person) String() string {
	if p.name == "" || p.age == 0 {
		return ""
	}
	return fmt.Sprintf("%v, %v", p.name, p.age)
}

func main() {
	p := Person{"Chakradhar", 20}
	fmt.Println(p)
}
