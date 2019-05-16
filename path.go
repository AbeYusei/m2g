package m2g

import (
	"path/filepath"
)

// Path file path wrapper
type Path struct {
	Raw string
	Dir string
	Mov string
	Gif string
}

// Convert2Path convert rawpath to struct object
func Convert2Path(rawPath string) *Path {
	ext := filepath.Ext(rawPath)
	base := rawPath[0 : len(rawPath)-len(ext)]
	return &Path{rawPath, base, base + ".mov", base + ".gif"}
}
