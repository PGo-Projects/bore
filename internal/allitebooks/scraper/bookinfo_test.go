package scraper

import (
	"fmt"
	"testing"
)

func TestGetBookInfo(t *testing.T) {
	title, pdfLink, category, summary, err := GetBookInfo("http://www.allitebooks.com/virtual-augmented-reality-for-dummies/")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(title)
	fmt.Println(pdfLink)
	fmt.Println(category)
	fmt.Println(summary)
}
