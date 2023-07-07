package main

import (
	"path/filepath"
	"time"

	"gorm.io/gorm"
)

type Item struct {
	Name  string `gorm:"type:varchar(100);not null"`
	Price int
	URL   string `gorm:"type:varchar(100);uniqueIndex"`
}

type LatestItem struct {
	Item
	CreatedAt time.Time
}

type ItemMaster struct {
	gorm.Model
	Item
	Description         string
	ImageURL            string
	ImageLastModifiedAt time.Time
	ImageDownloadPath   string
	PdfURL              string
	PdfLastModifiedAt   time.Time
	PdfDownloadPath     string
}

// テーブル名がitem_mastersになるのを防ぐため
func (ItemMaster) TableName() string {
	return "item_master"
}

func (i ItemMaster) ImageFileName() string {
	return filepath.Base(i.ImageURL)
}

func (i ItemMaster) PdfFileName() string {
	return filepath.Base(i.PdfURL)
}

func (i ItemMaster) equals(target ItemMaster) bool {
	return i.Description == target.Description &&
		i.ImageURL == target.ImageURL &&
		i.ImageLastModifiedAt == target.ImageLastModifiedAt &&
		i.ImageDownloadPath == target.ImageDownloadPath &&
		i.PdfURL == target.PdfURL &&
		i.PdfLastModifiedAt == target.PdfLastModifiedAt &&
		i.PdfDownloadPath == target.PdfDownloadPath
}

type HistoricalItem struct {
	Item
	CreatedAt time.Time
}
