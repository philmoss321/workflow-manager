package tasks

import model "stash.mtvi.com/scm/ms/hls-packager-service/core/model"

// FileDownload : download image from NASA!
type FileDownload struct {
	url              string
	downloadLocation string
	Downloader       model.Downloader
}

// Execute : Implement Task interface
func (nd *FileDownload) Execute() {
	nd.Downloader.Download(nd.url, nd.downloadLocation)
}
