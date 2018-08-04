package allitebooks

import (
	"fmt"
	"os"
	"strings"
	"syscall"

	"github.com/PGo-Projects/bore/internal/allitebooks/processor"
	"github.com/PGo-Projects/bore/internal/allitebooks/scraper"
	"github.com/PGo-Projects/signalhandler/pkg/signalhandler"
	"github.com/schollz/progressbar"
	"github.com/spf13/viper"

	tm "github.com/buger/goterm"
	jww "github.com/spf13/jwalterweatherman"
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
	bar := progressbar.NewOptions(a.startPage)
	urlFormat := "http://www.allitebooks.com/page/%d"
	sighandler := signalhandler.New(saveProgress, os.Interrupt, syscall.SIGTERM)
	foundStartURL := false

	tm.Clear()
	drawProgressBarSetup()
	bar.RenderBlank()

	logFile, err := os.OpenFile("errors.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		displayMessage("Unable to create error log", tm.RED)
		os.Exit(1)
	}
	jww.SetLogThreshold(jww.LevelInfo)
	jww.SetLogOutput(logFile)

	for pageNum := a.startPage; pageNum > 0; pageNum-- {
		viper.Set("allitebooks-startpage", pageNum)
		pageURL := fmt.Sprintf(urlFormat, pageNum)
		booklist, err := scraper.GetBookList(pageURL)
		if err != nil {
			message := fmt.Sprintf("There was an error retrieving booklist from %s", pageURL)
			displayMessage(message, tm.RED)
			jww.INFO.Println(message)
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
				message := fmt.Sprintf("There was an error retrieving info from %s", bookURL)
				displayMessage(message, tm.RED)
				jww.INFO.Println(message)
			}
			displayMessage(fmt.Sprintf("Processing %s", title), tm.WHITE)
			err = processor.ProcessBook(sighandler, processor.BookInfo{
				Title:    title,
				PdfLink:  pdfLink,
				Category: category,
				Summary:  summary,
			})
			if err != nil {
				message := fmt.Sprintf("There was an error processing %s", title)
				displayMessage(message, tm.RED)
				jww.INFO.Println(message)
			}
		}
		drawProgressBarSetup()
		bar.Add(1)
	}
}

func drawProgressBarSetup() {
	tm.MoveCursor(1, 1)
	tm.Flush()
}

func displayMessage(message string, color int) {
	tm.MoveCursor(1, 4)
	tm.Printf("%s\r", strings.Repeat(" ", tm.Width()))
	tm.MoveCursor(1, 4)
	tm.Print("   ", tm.Color(message, color))
	tm.Flush()
}

func saveProgress() {
	err := viper.WriteConfig()
	if err != nil {
		displayMessage("There was an error saving configuration...", tm.RED)
	}
	displayMessage("Exited!", tm.GREEN)
	os.Exit(0)
}

func (a *allitebooks) GetStartPage() int {
	return a.startPage
}

func (a *allitebooks) GetStartURL() string {
	return a.startURL
}
