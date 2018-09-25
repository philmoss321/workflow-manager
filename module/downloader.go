package module

// DownloadAsset : task
type DownloadAsset struct {
	Name string
}

// Execute : Execute task
func (da *DownloadAsset) Execute() error {
	return nil
}
