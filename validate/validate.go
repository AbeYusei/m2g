package validate

import (
	"os"

	"github.com/AbeYusei/m2g"
)

// Exists checks path is exists
func exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// Arg2Path arguments to struct object
func Arg2Path(args []string) (*m2g.Path, error) {

	if len(args) == 0 || len(args) < 2 {
		return nil, m2g.Err(m2g.InvalidArgLength)
	}

	rawPath := args[1]
	path := m2g.Convert2Path(rawPath)

	if path.Raw != path.Mov {
		return nil, m2g.Err(m2g.InvalidArgFile)
	}
	if !exists(path.Mov) {
		return nil, m2g.Err(m2g.NotExistsInputFile)
	}
	if exists(path.Gif) {
		return nil, m2g.Err(m2g.AlraedyExistsOutputFile)
	}

	return path, nil
}
