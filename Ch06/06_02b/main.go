package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://jsonplaceholder.typicode.com/posts"

func main() {
	fmt.Println("HTTP Request")

	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	checkError(err)

	req.Header.Set("User-Agent", "Go HTTP Client")

	resp, err := client.Do(req)
	checkError(err)

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	checkError(err)
	content := string(body)

	fmt.Println(content)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}