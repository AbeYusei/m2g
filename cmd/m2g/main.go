package main

import (
	"os"

	"github.com/AbeYusei/m2g/ffmpeg"
)

func main() {
	ffmpeg.Mov2gif(validate.arg2Path(os.Args))
}
