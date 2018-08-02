package processor

import (
	"io/ioutil"
	"os"
	"path"

	"github.com/PGo-Projects/bore/internal/allitebooks/utils"
	"github.com/PGo-Projects/signalhandler/pkg/signalhandler"
)

func ProcessBook(h *signalhandler.SignalHandler, title string, pdfLink string, category string, summary string) (err error) {
	return h.WithSignalBlocked(func() error {
		filename := path.Join(category, title+".pdf")
		txtFilename := path.Join(category, title+".txt")
		err = utils.CreateDirIfNotExist(category)
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
	})
}
