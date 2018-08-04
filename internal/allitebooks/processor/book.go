package processor

import (
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/PGo-Projects/bore/internal/allitebooks/utils"
	"github.com/PGo-Projects/signalhandler/pkg/signalhandler"
	tm "github.com/buger/goterm"
)

func ProcessBook(h *signalhandler.SignalHandler, title string, pdfLink string, category string, summary string) (err error) {
	workFunc := func() error {
		return processBook(title, pdfLink, category, summary)
	}
	messageFunc := func() {
		tm.MoveCursor(1, 5)
		tm.Printf("%s\r", strings.Repeat(" ", tm.Width()))
		tm.Flush()
		tm.MoveCursor(1, 4)
		tm.Printf("%s\r", strings.Repeat(" ", tm.Width()))
		tm.MoveCursor(1, 4)
		tm.Print("   ", tm.Color("Finishing up!", tm.BLUE))
		tm.Flush()
	}
	return h.WithSignalBlockedAndSignalMessageFunc(workFunc, messageFunc)
}

func processBook(title string, pdfLink string, category string, summary string) error {
	filename := path.Join(category, title+".pdf")
	txtFilename := path.Join(category, title+".txt")
	err := utils.CreateDirIfNotExist(category)
	if err != nil {
		return err
	}

	pdf, err := utils.GetFile(pdfLink)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, pdf, 0644)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(txtFilename, []byte(summary), 0644)
	if err != nil {
		os.Remove(filename)
	}
	return err
}
