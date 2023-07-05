package main

import "fmt"

func main() {
	baseURL := "http://localhost:8000"
	resp, err := fetch(baseURL)
	if err != nil {
		panic(err)
	}

	indexItems, err := parseList(resp)
	if err != nil {
		panic(err)
	}

	for _, item := range indexItems {
		fmt.Println(item)
	}
}
