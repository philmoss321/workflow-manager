package module

import (
	"errors"
	"log"

	model "stash.mtvi.com/scm/ms/hls-packager-service/core/model"
)

// DownloadAsset : task
type DownloadAsset struct {
	Name             string
	downloadURL      string
	downloadLocation string
}

// Execute : Execute task
func (da *DownloadAsset) Execute(adapter interface{}) error {
	log.Println(adapter)
	downloader, ok := adapter.(model.Downloader)
	if !ok {
		return errors.New("Downloader adapter required")
	}
	da.downloadURL = "http://a10.akadl.mtvnservices.com/534/mtvnorigin/gsp.alias/mediabus/vh1.com/2014/04/18/02/43/600901/5064020_600901_20140418144320323_1280x720_3500_h32.mp4?__gda__=1538296439_13d369c3cc377859649c96478299b059"
	da.downloadLocation = "/Users/philmoss/Desktop/shit/testmyshit.mp4"
	err := downloader.Download(da.downloadURL, da.downloadLocation)
	if err != nil {
		return err
	}
	log.Println("DOWNLOAD")
	return nil
}
