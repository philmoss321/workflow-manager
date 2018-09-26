package adapter

import (
	"io"
	"net/http"
	"os"

	"gopkg.in/cheggaaa/pb.v1"
)

// HTTPRequest : http adapter
type HTTPRequest struct {
	Name string
}

// Download : Fulfill Downloader interface
func (h *HTTPRequest) Download(url string, downloadLocation string) error {

	// Create the file
	out, err := os.Create(downloadLocation)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	contentLength := resp.ContentLength
	bar := pb.New(int(contentLength)).SetUnits(pb.U_BYTES)
	bar.Start()
	reader := bar.NewProxyReader(resp.Body)

	// Write the body to file
	_, err = io.Copy(out, reader)
	if err != nil {
		return err
	}

	return nil
}
