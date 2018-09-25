package model

// DBMultiReturn : Return array of workflows
type DBMultiReturn struct {
	Workflows []DBReturn
}

// DBReturn : Return single workflow
type DBReturn struct {
	Workflow WorkflowStatus
}

// WorkflowStatus : Returned workflow
type WorkflowStatus struct {
	Status string       `json:"status,omitempty"`
	Data   WorkflowData `json:"data,omitempty"`
}

// Config : General application configuration
type Config struct {
	Downloader downloadConfig
	Uploader   uploadConfig
	Encoder    EncodeConfig
	Database   map[string]databaseConfig
}

type downloadConfig struct {
	HTTPDownload bool `toml:"http_download"`
	Count        int
}

type uploadConfig struct {
	Count int
}

// EncodeConfig : encoding configuration options
type EncodeConfig struct {
	EncryptionMode       string `toml:"encryption_mode"`
	SegmentDuration      int    `toml:"segment_duration"`
	OutputSingleFile     bool   `toml:"output_single_file"`
	HlsVersion           int    `toml:"hls_version"`
	MasterFilename       string `toml:"master_filename"`
	IframeFilename       string `toml:"iframe_filename"`
	StreamLevelFilename  string `toml:"stream_level_filename"`
	TsFilename           string `toml:"ts_filename"`
	BinaryLocation       string `toml:"binary_location"`
	InitializationVector string `toml:"initialization_vector"`
	TempIframeMP4        string `toml:"temp_iframe_mp4"`
}

type databaseConfig struct {
	ReplicaSet     []string `toml:"replica_set"`
	Region         string
	User           string
	Pass           string
	TableName      string `toml:"table_name"`
	DBName         string `toml:"db_name"`
	Authentication string
	ReplicasetName string `toml:"replicaset_name"`
	AuthDB         string `toml:"auth_db"`
	UseAuth        bool   `toml:"use_auth"`
	UseLocal       bool   `toml:"use_local"`
	LocalURL       string `toml:"local_url"`
}

// WorkflowData : returned workflow data
type WorkflowData struct {
	IframeLocation string `json:"iframeLocation,omitempty"`
	UploadPath     string `json:"uploadPath,omitempty"`
	AssetName      string `json:"assetName,omitempty"`
	Namespace      string `json:"namespace,omitempty"`
}

// StatusReturn : json for status returns
type StatusReturn struct {
	Complete   bool        `json:"complete"`
	Renditions []Rendition `json:"renditions,omitempty"`
}

// Rendition : HLS rendition data
type Rendition struct {
	Ref    string `json:"ref,omitempty"`
	Path   string `json:"path,omitempty"`
	Iframe string `json:"iframe,omitempty"`
}
