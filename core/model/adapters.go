package model

// Downloader : interface for Download adapter
type Downloader interface {
	Download(url string, downloadLocation string) error
}

// Uploader : interface for Upload adapter
type Uploader interface {
	Upload(url string, localPath string) error
}

// Encoder : interface for encoding adapter
type Encoder interface {
}

// Segmenter : interface for segmenting adapter
type Segmenter interface {
	GenerateHLS(*EncodeConfig) error
}

// FileSystem : interface for filesystem adapter
type FileSystem interface {
	CreateDir(dirName string) error
	DeleteDir(dirName string) error
}
