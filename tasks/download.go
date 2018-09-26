package tasks

import model "stash.mtvi.com/scm/ms/hls-packager-service/core/model"

// FileDownload : download image from NASA!
type FileDownload struct {
	url              string
	downloadLocation string
}

// Execute : Implement Task interface
func (nd *FileDownload) Execute(downloader model.Downloader) {
	downloader.Download(nd.url, nd.downloadLocation)
}
