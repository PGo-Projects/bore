package scraper

import "testing"

func TestGetLastURLForPage(t *testing.T) {
	lastPage, err := GetTotalPages("http://www.allitebooks.com")
	if err != nil {
		t.Error(err)
	}
	lastURL, err := GetLastURLForPage("http://www.allitebooks.com", lastPage)
	if err != nil {
		t.Error(err)
	}
	if lastURL != "http://www.allitebooks.com/powershell-for-sql-server-essentials/" {
		t.Error("The scraped last url does not match!")
	}
}

func TestGetTotalPages(t *testing.T) {
	pages, err := GetTotalPages("http://www.allitebooks.com")
	if err != nil {
		t.Error(err)
	}
	if pages != 770 {
		t.Error("The scraped page count does not match!")
	}
}
