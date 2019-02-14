package armory

import (
	"os"
	"path/filepath"
)

type pilot struct{}

var Pilot *pilot

// AppPath AppPath
func (p *pilot) AppPath(subPath *string) *string {
	rootPath, _ := os.Executable()
	s := filepath.Join(filepath.Dir(rootPath), *subPath)
	return &s
}
