package main

import (
	"fmt"
)

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

	// _ をdbに変更
	db, err := connectDB()
	if err != nil {
		panic(err)
	}

	err = migrateDB(db)
	if err != nil {
		panic(err)
	}

	err = createLatestItems(items, db)
	if err != nil {
		panic(err)
	}

	err = updateItemMaster(db)
	if err != nil {
		panic(err)
	}
}
