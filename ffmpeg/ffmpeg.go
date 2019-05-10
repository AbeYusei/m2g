package ffmpeg

import (
	"fmt"
	"os/exec"

	"github.com/AbeYusei/m2g"
	"github.com/AbeYusei/m2g/error"
)

// Mov2Gif convert .mov to .gif
func Mov2Gif(p *m2g.Path) {
	m := p.Mov
	g := p.Gif
	out, err := exec.Command("ffmpeg", "-i", m, "-r", "10", g).CombinedOutput()
	if err != nil {
		error.HandleError(err)
	}
	fmt.Println(string(out))
}
