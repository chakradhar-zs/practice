package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func sendRequest() {

	client := http.Client{}

	req, err := http.NewRequest("DELETE", "http://www.example.com/bucket/sample", nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	// Read Response Body
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("response Status : ", resp.Status)
	fmt.Println("response Headers : ", resp.Header)
	fmt.Println("response Body : ", string(respBody))
}

func main() {
	sendRequest()
}
