package scraper

import (
	"fmt"
	"strings"

	"github.com/PGo-Projects/bore/internal/allitebooks/utils"
)

func cleanStringForPath(s string) string {
	withoutAmpersand := strings.Replace(s, "&", "and", -1)
	withoutSpaces := strings.Replace(withoutAmpersand, " ", "-", -1)
	return withoutSpaces
}

func GetBookInfo(url string) (title string, pdfLink string, category string, summary string, err error) {
	doc, err := utils.GetGoqueryDocument(url)
	if err != nil {
		return "", "", "", "", err
	}
	rawTitle := doc.Find("#main-content h1.single-title").Text()
	title = cleanStringForPath(rawTitle)
	pdfLink, ok := doc.Find("span.download-links a").Attr("href")
	if !ok {
		return "", "", "", "", fmt.Errorf("Unable to retrieve pdf link for %s", url)
	}
	rawCategory := doc.Find("dt:contains(Category)").Next().Find("a").Text()
	category = cleanStringForPath(rawCategory)
	summary = doc.Find("div.entry-content p").Text()
	return title, pdfLink, category, summary, nil
}
