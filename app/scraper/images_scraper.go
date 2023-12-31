package scraper

import (
	"io"
	"net/http"
	"os"
)

type imagesScraper struct{}

// Download the image given an URL and the filename (with image path)
func (s *imagesScraper) DownloadImage(url string, filename string) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	return nil
}

func NewImageScraper() *imagesScraper {
	return &imagesScraper{}
}
