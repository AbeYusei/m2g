package validate

import (
	"os"
	"path/filepath"
)

// Exists checks path is exists
func Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// Convert2Path convert rawpath to struct object
func Convert2Path(rawPath string) *Path {
	ext := filepath.Ext(rawPath)
	base := rawPath[0 : len(rawPath)-len(ext)]
	return &Path{rawPath, base + ".mov", base + ".gif"}
}

// Arg2Path arguments to struct object
func Arg2Path(args []string) *Path {

	if len(args) < 2 {
		handleError("引数に変換対象の.movファイルを指定してください。")
	}

	rawPath := args[1]
	path := Convert2Path(rawPath)

	if path.raw != path.mov {
		handleError(".movファイルを指定してください。")
	}
	if Exists(path.gif) {
		handleError("ファイル " + path.gif + " がすでに存在します。")
	}

	return path
}
