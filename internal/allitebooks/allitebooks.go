package allitebooks

import (
	"fmt"
	"os"
	"syscall"

	"github.com/PGo-Projects/bore/internal/allitebooks/processor"
	"github.com/PGo-Projects/bore/internal/allitebooks/scraper"
	"github.com/PGo-Projects/signalhandler/pkg/signalhandler"
	"github.com/spf13/viper"
)

type Allitebooks interface {
	GetAll()

	GetStartPage() int
	GetStartURL() string
}

type allitebooks struct {
	startPage int
	startURL  string
}

func (a *allitebooks) GetAll() {
	urlFormat := "http://www.allitebooks.com/page/%d"
	sighandler := signalhandler.New(saveProgress, os.Interrupt, syscall.SIGTERM)
	foundStartURL := false
	for pageNum := a.startPage; pageNum > 0; pageNum-- {
		viper.Set("allitebooks-startpage", pageNum)
		pageURL := fmt.Sprintf(urlFormat, pageNum)
		booklist, err := scraper.GetBookList(pageURL)
		if err != nil {
			fmt.Printf("There was an error retrieving booklist from %s", pageURL)
		}
		for index := len(booklist) - 1; index >= 0; index-- {
			bookURL := booklist[index]
			if bookURL == viper.GetString("allitebooks-starturl") {
				foundStartURL = true
			}
			if !foundStartURL {
				continue
			}
			viper.Set("allitebooks-starturl", bookURL)
			title, pdfLink, category, summary, err := scraper.GetBookInfo(bookURL)
			if err != nil {
				fmt.Printf("There was an error retrieving info from %s\n", bookURL)
			}
			fmt.Printf("- Processing %s\n", title)
			err = processor.ProcessBook(sighandler, title, pdfLink, category, summary)
			if err != nil {
				fmt.Printf("There was an error processing %s\n", title)
			}
		}
	}
}

func saveProgress() {
	err := viper.WriteConfig()
	if err != nil {
		fmt.Println("Was not able to save the config")
	}
	fmt.Println("\nExited!")
	os.Exit(0)
}

func (a *allitebooks) GetStartPage() int {
	return a.startPage
}

func (a *allitebooks) GetStartURL() string {
	return a.startURL
}
