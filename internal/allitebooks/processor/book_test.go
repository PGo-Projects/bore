package processor

import (
	"fmt"
	"os"
	"syscall"
	"testing"

	"github.com/PGo-Projects/bore/internal/allitebooks/scraper"
	"github.com/PGo-Projects/signalhandler/pkg/signalhandler"
)

func TestProcessBookValidCase(t *testing.T) {
	handlerFunc := func() {
		fmt.Println("Inside test handler")
	}
	handler := signalhandler.New(handlerFunc, os.Interrupt, syscall.SIGTERM)
	title, pdfLink, category, summary, err := scraper.GetBookInfo("http://www.allitebooks.com/virtual-augmented-reality-for-dummies/")
	if err != nil {
		t.Errorf("Cannot retrieve information; %s", err)
	}
	err = ProcessBook(handler, title, pdfLink, category, summary)
	if err != nil {
		os.RemoveAll(category)
		t.Error("Should not have errored!")
	}
	os.RemoveAll(category)
}

func TestProcessBookInvalidCase(t *testing.T) {
	handlerFunc := func() {
		fmt.Println("Inside test handler")
	}
	handler := signalhandler.New(handlerFunc, os.Interrupt, syscall.SIGTERM)
	title, _, category, summary, err := scraper.GetBookInfo("http://www.allitebooks.com/virtual-augmented-reality-for-dummies/")
	if err != nil {
		t.Error(err)
	}
	err = ProcessBook(handler, title, "https://google.pdf", category, summary)
	if err == nil {
		os.RemoveAll(category)
		t.Error("Should have errored")
	}
	os.RemoveAll(category)
}
