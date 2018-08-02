package scraper

import "testing"

func TestGetTotalPages(t *testing.T) {
	pages, err := GetTotalPages("http://www.allitebooks.com")
	if err != nil {
		t.Error(err)
	}
	if pages != 769 {
		t.Error("The scraped page count does not matched!")
	}
}
