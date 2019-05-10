package main

import (
	"os"

	"../../validate"
	"../../ffmpeg"
	"../../error"
)

func main() {
	ffmpeg.Mov2gif(validate.arg2Path(os.Args))
}
