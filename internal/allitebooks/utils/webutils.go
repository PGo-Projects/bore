package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func GetFile(url string) (content []byte, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func GetFileString(url string) (content string, err error) {
	bytes, err := GetFile(url)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func GetGoqueryDocument(url string) (document *goquery.Document, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Status code error: %d %s", resp.StatusCode, resp.Status)
	}
	return goquery.NewDocumentFromReader(resp.Body)
}
