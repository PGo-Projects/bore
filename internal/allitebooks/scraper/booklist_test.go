package scraper

import (
	"regexp"
	"testing"
)

func TestGetBookList(t *testing.T) {
	list, err := GetBookList("http://www.allitebooks.com")
	if err != nil {
		t.Error(err)
	}
	pattern := regexp.MustCompile(`http://www.allitebooks.com/[\w-]+/?`)
	for _, url := range list {
		if !pattern.MatchString(url) {
			t.Errorf("The url %s does not match the normal pattern", url)
		}
	}
}
