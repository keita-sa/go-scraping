package main

import "fmt"

func main() {
	conf, err := loadConfig()
	if err != nil {
		panic(err)
	}

	// baseURLの宣言を削除
	// baseURL := "http://localhost:8000"
	resp, err := fetch(conf.BaseURL)
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
	db, err := connectDB(conf)
	if err != nil {
		panic(err)
	}

	err = migrateDB(db)
	if err != nil {
		panic(err)
	}

	indexItems, err = fetchMultiPages(conf.BaseURL)
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
	// currentDirとdownloadBasePathの宣言を削除
	// currentDir, _ := os.Getwd()
	// downloadBasePath := filepath.Join(currentDir, "work", "downloadFiles")
	// 引数でconf.DownloadBasePathを与えるように変更
	updatedItems, err = fetchDetails(updateChkItems, conf.DownloadBasePath)
	if err != nil {
		panic(err)
	}

	if err = createDetails(updatedItems, db); err != nil {
		panic(err)
	}
}
