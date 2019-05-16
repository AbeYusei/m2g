package ffmpeg

import (
	"log"

	"github.com/AbeYusei/m2g"
)

// Mov2Gif convert .mov to .gif
func Mov2Gif(p *m2g.Path) (string, error) {
	m := p.Mov
	g := p.Gif

	log.Println("output gif file to: " + g)

	args := []string{"-i", m, "-r", "10", g}

	log.Print(p)
	err := Exec(p.Dir, "ffmpeg", args)

	return "ffmpeg done.", err
}
