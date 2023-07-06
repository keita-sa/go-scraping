package main

import (
	"fmt"
	"os"
	"path/filepath"
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

	err = createLatestItems(indexItems, db)
	if err != nil {
		panic(err)
	}

	err = updateItemMaster(db)
	if err != nil {
		panic(err)
	}

	var updateChkItems []ItemMaster
	updateChkItems, err = findItemMaster(db)
	if err != nil {
		panic(err)
	}

	var updatedItems []ItemMaster
	currentDir, _ := os.Getwd()
	downloadBasePath := filepath.Join(currentDir, "work", "downloadFiles")
	updatedItems, err = fetchDetails(updateChkItems, downloadBasePath)
	if err != nil {
		panic(err)
	}

	if err = createDetails(updatedItems, db); err != nil {
		panic(err)
	}
}
