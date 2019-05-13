package ffmpeg

import (
	"os/exec"

	"github.com/AbeYusei/m2g"
)

// Mov2Gif convert .mov to .gif
func Mov2Gif(p *m2g.Path) (string, error) {
	m := p.Mov
	g := p.Gif
	out, err := exec.Command("ffmpeg", "-i", m, "-r", "10", g).CombinedOutput()
	if err != nil {
		return "error", m2g.Err(m2g.InternalProcess)
	}

	return string(out), nil
}
