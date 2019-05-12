package validate

import (
	"os"

	"github.com/AbeYusei/m2g"
	"github.com/AbeYusei/m2g/error"
)

// Exists checks path is exists
func exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// Arg2Path arguments to struct object
func Arg2Path(args []string) *m2g.Path {

	if len(args) < 2 {
		error.Handle("引数に変換対象の.movファイルを指定してください。")
	}

	rawPath := args[1]
	path := m2g.Convert2Path(rawPath)

	if path.Raw != path.Mov {
		error.Handle(".movファイルを指定してください。")
	}
	if !exists(path.Mov) {
		error.Handle("指定された.movファイルが存在しません。")
	}
	if exists(path.Gif) {
		error.Handle("ファイル " + path.Gif + " がすでに存在します。")
	}

	return path
}
