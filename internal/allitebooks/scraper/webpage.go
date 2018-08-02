package scraper

import (
	"errors"
	"regexp"
	"strconv"

	"github.com/PGo-Projects/bore/internal/allitebooks/utils"
)

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
