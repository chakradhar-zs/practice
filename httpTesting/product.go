package main

import (
	"fmt"
	"net/http"
	"time"
)

func ProductHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("content-type", "application/json")
	if r.Method == http.MethodPost {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func HttpClient(url string) (int, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return 0, err
	}
	c := http.Client{Timeout: time.Minute}

	res, err := c.Do(req)
	if err != nil {
		return 0, err
	}
	return res.StatusCode, nil
}

func main() {
	fmt.Println("Server Listening at port 3000")
	http.ListenAndServe(":3000", http.HandlerFunc(ProductHandler))

}
