package scraper

import "testing"

func TestGetLastURL(t *testing.T) {
	lastURL, err := GetLastURL("http://www.allitebooks.com")
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
