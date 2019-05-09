package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

type Path struct {
	raw string
	mov string
	gif string
}

func exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func convert2path(rawPath string) *Path {
	ext := filepath.Ext(rawPath)
	base := rawPath[0 : len(rawPath)-len(ext)]
	return &Path{rawPath, base + ".mov", base + ".gif"}
}

func arg2Path(args []string) *Path {

	if len(args) < 2 {
		handleError("引数に変換対象の.movファイルを指定してください。")
	}

	rawPath := args[1]
	path := convert2path(rawPath)

	if path.raw != path.mov {
		handleError(".movファイルを指定してください。")
	}
	if exists(path.gif) {
		handleError("ファイル " + path.gif + " がすでに存在します。")
	}

	return path
}

func handleError(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

func mov2gif(path *Path) {
	out, err := exec.Command("ffmpeg", "-i", path.mov, "-r", "10", path.gif).CombinedOutput()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(out))
}

func main() {
	mov2gif(arg2Path(os.Args))
}
