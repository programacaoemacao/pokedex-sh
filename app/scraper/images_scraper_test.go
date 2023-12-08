package scraper

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestImagesScraper_DownloadImage(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			imageContent := []byte{0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A}
			rw.Write(imageContent)
		}))
		defer server.Close()

		scraper := &imagesScraper{}

		tempFile, err := os.CreateTemp("", "test_image_*.png")
		if err != nil {
			t.Fatal(err)
		}
		defer os.Remove(tempFile.Name())

		err = scraper.DownloadImage(server.URL, tempFile.Name())
		require.NoError(t, err)

		fileInfo, err := os.Stat(tempFile.Name())
		require.NoError(t, err)
		require.NotEqual(t, 0, fileInfo.Size(), "The downloaded file is empty")
	})
}
