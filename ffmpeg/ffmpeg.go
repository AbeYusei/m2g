package ffmpeg

import (
	"fmt"
	"os/exec"

	"github.com/AbeYusei/ffmpeg_custom/common/util"
)

// Mov2Gif convert .mov to .gif
func Mov2Gif(path *Path) {
	out, err := exec.Command("ffmpeg", "-i", path.mov, "-r", "10", path.gif).CombinedOutput()
	if err != nil {
		util.Handle(err)
	}
	fmt.Println(string(out))
}
