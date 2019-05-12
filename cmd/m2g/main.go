package main

import (
	"os"

	"github.com/AbeYusei/m2g/ffmpeg"
	"github.com/AbeYusei/m2g/validate"
)

func main() {
	ffmpeg.Mov2Gif(validate.Arg2Path(os.Args))
}
