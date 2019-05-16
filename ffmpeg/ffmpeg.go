package ffmpeg

import (
	"log"
	"os/exec"

	"github.com/AbeYusei/m2g"
)

// Mov2Gif convert .mov to .gif
func Mov2Gif(p *m2g.Path) (string, error) {
	m := p.Mov
	g := p.Gif

	log.Println("output gif file to: " + g)

	args := []string{"-i", m, "-r", "10", g}

	out, err := exec.Command("ffmpeg", args...).CombinedOutput()
	if err != nil {
		return "error", m2g.Err(m2g.InternalProcess)
	}

	return string(out), nil
}
