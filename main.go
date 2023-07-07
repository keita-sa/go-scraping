package main

func main() {
	conf, err := loadConfig()
	if err != nil {
		panic(err)
	}

	db, err := connectDB(conf)
	if err != nil {
		panic(err)
	}

	err = migrateDB(db)
	if err != nil {
		panic(err)
	}

	items, err := fetchMultiPages(conf.BaseURL)
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

	var updateChkItems []ItemMaster
	updateChkItems, err = findItemMaster(db)
	if err != nil {
		panic(err)
	}

	var updatedItems []ItemMaster
	updatedItems, err = fetchDetails(updateChkItems, conf.DownloadBasePath)
	if err != nil {
		panic(err)
	}

	if err = createDetails(updatedItems, db); err != nil {
		panic(err)
	}
}
