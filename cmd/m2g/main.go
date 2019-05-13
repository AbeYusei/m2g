package main

import (
	"fmt"
	"os"

	"github.com/AbeYusei/m2g/ffmpeg"
	"github.com/AbeYusei/m2g/validate"
)

func main() {
	p, err := validate.Arg2Path(os.Args)
	if err != nil {
		handleError(err)
	}

	o, err := ffmpeg.Mov2Gif(p)
	if err != nil {
		handleError(err)
	}

	fmt.Println(o)
}

func handleError(err error) {
	fmt.Println(err)
	os.Exit(1)
}
