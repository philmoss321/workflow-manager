package module

import (
	"log"
)

// CreateAssetDirectories : task
type CreateAssetDirectories struct {
	Name string
}

// Execute : Execute task
func (cad *CreateAssetDirectories) Execute(adapter interface{}) error {
	log.Println("CREATE!")
	return nil
}
