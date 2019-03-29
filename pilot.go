package armory

import (
	"os"
	"path/filepath"
)

type pilot struct{}

// Pilot Pilot
var Pilot *pilot

// AppPath AppPath
func (p *pilot) AppPath(subPath string) string {
	rootPath, _ := os.Executable()
	return filepath.Join(filepath.Dir(rootPath), subPath)
}
