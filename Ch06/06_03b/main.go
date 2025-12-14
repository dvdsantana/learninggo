package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const url = "https://jsonplaceholder.typicode.com/posts"

func main() {
	fmt.Println("HTTP Request Parser")

	content := readHTTPContent()
	fmt.Println(content)

	tours := PostsFromJSON(content)
	for _, tour := range tours {
		fmt.Printf("%v: ($%v)\n", tour.Title, tour.Body)
	}
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func readHTTPContent() string {
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

	return content
}

func PostsFromJSON(content string) []Post {
	posts := make([]Post, 0)
	decoder := json.NewDecoder(strings.NewReader(content))
	_, err := decoder.Token()
	checkError(err)

	var post Post
	for decoder.More() {
		err := decoder.Decode(&post)
		checkError(err)
		posts = append(posts, post)
	}
	return posts
}

type Post struct {
	UserId int `json:"userId"`
	Id     int `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}