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

type BookInfo struct {
	Title    string
	PdfLink  string
	Category string
	Summary  string
}

func ProcessBook(h *signalhandler.SignalHandler, basepath string, bookInfo BookInfo) (err error) {
	workFunc := func() error {
		return processBook(basepath, bookInfo)
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

func processBook(basepath string, bookInfo BookInfo) error {
	fileDirectory := path.Join(basepath, bookInfo.Category)
	filename := path.Join(fileDirectory, bookInfo.Title+".pdf")
	txtFilename := path.Join(fileDirectory, bookInfo.Title+".txt")
	err := utils.CreateDirIfNotExist(fileDirectory)
	if err != nil {
		return err
	}

	pdf, err := utils.GetFile(bookInfo.PdfLink)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, pdf, 0644)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(txtFilename, []byte(bookInfo.Summary), 0644)
	if err != nil {
		os.Remove(filename)
	}
	return err
}
