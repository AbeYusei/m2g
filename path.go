package m2g

import (
	"path/filepath"
)

// Path file path wrapper
type Path struct {
	Raw string
	Mov string
	Gif string
}

// Covert2Path convert rawpath to struct object
func Covert2Path(rawPath string) *Path {
	ext := filepath.Ext(rawPath)
	base := rawPath[0 : len(rawPath)-len(ext)]
	return &Path{rawPath, base + ".mov", base + ".gif"}
}
