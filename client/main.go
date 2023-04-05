package main

import (
	"fmt"
	"io"
	"net/http"
)

func getRespBody(url string) string {
	response, err := http.Get(url)
	if err != nil {
		return ""
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {

		scanner, err := io.ReadAll(response.Body)

		if err != nil {
			return ""
		}
		return string(scanner)
	}
	return ""
}

func main() {
	fmt.Println(getRespBody("http://localhost:8080/ping"))
}
