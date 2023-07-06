package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"net/url"
	"path/filepath"
	"strconv"
	"strings"
)

func parseList(resp *http.Response) ([]Item, error) {
	// レスポンスボディを取得
	body := resp.Body

	// レスポンスに含まれているリクエスト情報からリクエスト先のURLを取得
	requestURL := *resp.Request.URL

	var items []Item

	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		return nil, fmt.Errorf("get document error: %w", err)
	}

	// Find関数でテーブルの行要素を取得
	tr := doc.Find("table tr")
	notFoundMessage := "ページが存在しません"
	if strings.Contains(doc.Text(), notFoundMessage) || tr.Size() == 0 {
		return nil, nil
	}

	tr.Each(func(_ int, s *goquery.Selection) {
		item := Item{}

		// Find関数を使用して商品の各要素を取得
		item.Name = s.Find("td:nth-of-type(2) a").Text()
		item.Price, _ = strconv.Atoi(strings.ReplaceAll(strings.ReplaceAll(s.Find("td:nth-of-type(3)").Text(), ",", ""), "円", ""))
		itemURL, exists := s.Find("td:nth-of-type(2) a").Attr("href")
		refURL, parseErr := url.Parse(itemURL)

		if exists && parseErr == nil {
			// requestURLとrefURLを結合して絶対URLを取得
			item.URL = (*requestURL.ResolveReference(refURL)).String()
		}

		if item.Name != "" {
			items = append(items, item)
		}
	})

	return items, nil
}

func parseDetail(response *http.Response, item ItemMaster, downloadBasePath string) (ItemMaster, error) {
	body := response.Body
	requestURL := *response.Request.URL
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		return ItemMaster{}, fmt.Errorf("get detail page document body error %w", err)
	}

	item.Description = doc.Find("table tr:nth-of-type(2) td:nth-of-type(2)").Text()

	// 以下追加
	// Image
	href, exists := doc.Find("table tr:nth-of-type(1) td:nth-of-type(1) img").Attr("src")
	refURL, parseErr := url.Parse(href)
	if exists && parseErr == nil {
		imageURL := (*requestURL.ResolveReference(refURL)).String()
		// checkFileUpdated関数はこの後で実装
		isUpdated, currentLastModified := checkFileUpdated(imageURL, item.ImageLastModifiedAt)
		if isUpdated {
			item.ImageURL = imageURL
			item.ImageLastModifiedAt = currentLastModified

			imageDownloadPath := filepath.Join(downloadBasePath, "img", strconv.Itoa(int(item.ID)), item.ImageFileName())
			err := downloadFile(imageURL, imageDownloadPath)
			if err != nil {
				return ItemMaster{}, fmt.Errorf("download image error: %w", err)
			}
			item.ImageDownloadPath = imageDownloadPath
		}
	}

	// PDF
	href, exists = doc.Find("table tr:nth-of-type(3) td:nth-of-type(2) a").Attr("href")
	refURL, parseErr = url.Parse(href)
	if exists && parseErr == nil {
		pdfURL := (*requestURL.ResolveReference(refURL)).String()
		// checkFileUpdated関数はこの後で実装
		isUpdated, currentLastModified := checkFileUpdated(pdfURL, item.PdfLastModifiedAt)
		if isUpdated {
			item.PdfURL = pdfURL
			item.PdfLastModifiedAt = currentLastModified

			pdfDownloadPath := filepath.Join(downloadBasePath, "pdf", strconv.Itoa(int(item.ID)), item.PdfFileName())
			err := downloadFile(pdfURL, pdfDownloadPath)
			if err != nil {
				return ItemMaster{}, fmt.Errorf("download pdf error: %w", err)
			}
			item.PdfDownloadPath = pdfDownloadPath
		}
	}
	return item, nil
}
