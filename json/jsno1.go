package main

import (
	"encoding/json"
	"fmt"
)

//type response1 struct {
//	Page   int
//	Fruits []string
//}

type response2 struct {
	Page   int      `json:"p"`
	Fruits []string `json:"f"`
}

func main() {
	//ex, _ := json.Marshal(true)
	//fmt.Println(ex)
	//fmt.Println(string(ex))

	//lst := []string{"hi", "hello", "gophers"}
	//b, _ := json.Marshal(lst)
	//fmt.Println(b)
	//fmt.Println(string(b))

	//res1 := &response1{
	//	Page:   1,
	//	Fruits: []string{"apple", "orange", "mango"},
	//}

	res2 := &response2{
		Page:   1,
		Fruits: []string{"apple", "orange", "mango"},
	}

	//b1, _ := json.Marshal(res1)
	//fmt.Println(b1)
	//fmt.Println(string(b1))

	//fmt.Println()

	b2, _ := json.Marshal(res2)
	fmt.Println(b2)
	fmt.Println(string(b2))
	fmt.Println()

	//result1 := response1{
	//	Page:   1,
	//	Fruits: []string{"apple", "orange", "mango"},
	//}

	//var result2, b3 response2
	////json.Unmarshal(b1, result1)
	////fmt.Println(result1)
	//b3 := response2{1, {"apple", "orange", "mango"}}

	//json.Unmarshal(b3, result2)
	//fmt.Println(result2)
	//fmt.Println(result2.Page)

}
