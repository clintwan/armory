package qm

import (
	"os"
	"path/filepath"
)

// AppPath AppPath
func AppPath(subPath *string) *string {
	rootPath, _ := os.Executable()
	s := filepath.Join(filepath.Dir(rootPath), *subPath)
	return &s
}
