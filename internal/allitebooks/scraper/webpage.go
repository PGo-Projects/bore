package scraper

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"

	"github.com/PGo-Projects/bore/internal/allitebooks/utils"
)

func GetLastURLForPage(homepage string, pageNum int) (url string, err error) {
	lastPageURL := fmt.Sprintf("%s/page/%d", homepage, pageNum)
	doc, err := utils.GetGoqueryDocument(lastPageURL)
	if err != nil {
		return "", err
	}

	url, ok := doc.Find("h2.entry-title a").Last().Attr("href")
	if ok {
		return url, nil
	}
	return "", errors.New("Unable to find last url")
}

func GetTotalPages(homepage string) (pageCount int, err error) {
	doc, err := utils.GetGoqueryDocument(homepage)
	if err != nil {
		return 0, err
	}
	pageCountText := doc.Find("span.pages").Text()
	pattern := regexp.MustCompile(`\d+ / (\d+)`)
	matches := pattern.FindStringSubmatch(pageCountText)
	if len(matches) != 2 {
		return 0, errors.New("Unable to find page count")
	}
	pageCount, _ = strconv.Atoi(matches[1])
	return pageCount, nil
}
