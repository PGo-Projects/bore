package allitebooks

import (
	"fmt"
	"os"
	"syscall"

	"github.com/PGo-Projects/bore/internal/allitebooks/processor"
	"github.com/PGo-Projects/bore/internal/allitebooks/scraper"
	"github.com/PGo-Projects/signalhandler/pkg/signalhandler"
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
	for pageNum := a.startPage; pageNum > 0; pageNum-- {
		booklist, err := scraper.GetBookList(fmt.Sprintf(urlFormat, pageNum))
		if err != nil {
			fmt.Printf("There was an error retrieving booklist from %s", fmt.Sprintf(urlFormat, pageNum))
		}
		for index := len(booklist) - 1; index >= 0; index-- {
			bookURL := booklist[index]
			title, pdfLink, category, summary, err := scraper.GetBookInfo(bookURL)
			if err != nil {
				fmt.Printf("There was an error retrieving info from %s\n", bookURL)
			}
			err = processor.ProcessBook(sighandler, title, pdfLink, category, summary)
			if err != nil {
				fmt.Printf("There was an error processing %s\n", title)
			}
		}
	}
}

func saveProgress() {

}

func (a *allitebooks) GetStartPage() int {
	return a.startPage
}

func (a *allitebooks) GetStartURL() string {
	return a.startURL
}
