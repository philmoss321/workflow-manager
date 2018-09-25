package model

// Message : incoming message POST format
type Message struct {
	URL            string         `json:"url,omitempty" valid:"required~Invalid URL provided,url~Invalid URL provided"`
	Namespace      string         `json:"namespace,omitempty" valid:"required~namespace value is missing"`
	PackageType    []string       `json:"packageType,omitempty" valid:"-"`
	LanguageCode   string         `json:"languageCode,omitempty" valid:"-"`
	CuePointList   []int          `json:"cuePointList,omitempty" valid:"-"`
	ReferenceVideo ReferenceVideo `json:"referenceVideo,omitempty" valid:"-"`
	ForcePackage   bool           `json:"forcePackage,omitempty" valid:"-"`
}

// ReferenceVideo : Video source reference for audio asset
type ReferenceVideo struct {
	URL          string `json:"url,omitempty" valid:"required~Reference video URL required in audio and caption workflows,url~Invalid URL provided in reference video"`
	Namespace    string `json:"namespace,omitempty" valid:"required~Reference video namespace required in audio and caption workflows"`
	LanguageCode string `json:"languageCode,omitempty" valid:"-"`
}
