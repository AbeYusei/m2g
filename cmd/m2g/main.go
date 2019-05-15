package main

import (
	"fmt"
	"log"
	"os"

	"github.com/AbeYusei/m2g/ffmpeg"
	"github.com/AbeYusei/m2g/validate"
)

var exit = os.Exit

func main() {
	p, err := validate.Arg2Path(os.Args)
	if err != nil {
		os.Exit(e(err))
	}

	o, err := ffmpeg.Mov2Gif(p)
	if err != nil {
		os.Exit(e(err))
	}

	fmt.Println(o)
}

// error handling with logging proxy
func e(err error) int {
	log.Println(err)
	fmt.Println(err)
	return 1
}
