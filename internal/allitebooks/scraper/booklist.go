package scraper

import (
	"github.com/PGo-Projects/bore/internal/allitebooks/utils"
	"github.com/PuerkitoBio/goquery"
)

func GetBookList(url string) (list []string, err error) {
	doc, err := utils.GetGoqueryDocument(url)
	if err != nil {
		return nil, err
	}
	list = make([]string, 0)
	doc.Find("h2.entry-title a").Each(func(i int, s *goquery.Selection) {
		url, ok := s.Attr("href")
		if ok {
			list = append(list, url)
		}
	})
	return list, nil
}
