package entity

import (
	"fmt"
	"image"
	"image/color"
	"strings"
	"sync"
)

type Artist struct {
	ramp   string
	scaleX int
	scaleY int
	canvas [][]string
}

// NewArtist is a constructor for the Artist entity.
// @return 		Artist			Artist entity instance.
func NewArtist() *Artist {
	/** TODO - In a future version, I'd change the attributes into parameters
	 * 		  or env variables, depending on how flexible the team wants to be.
	 **/
	dali := new(Artist)
	dali.ramp = "@#+=." // "$@B%8&WM#*oahkbdpqwmZO0QLCJUYXzcvunxrjft/\|()1{}[]?-_+~<>i!lI;:,"^`'."
	dali.scaleX = 32
	dali.scaleY = 16
	return dali
}

func (dali *Artist) Paint(muse image.Image) string {
	println(fmt.Sprintf("Content: %+v", muse))
	max := muse.Bounds().Max
	dali.canvas = dali.getBlankCanvas(max)
	var wg sync.WaitGroup
	for y := 0; y < max.Y; y += dali.scaleX {
		for x := 0; x < max.X; x += dali.scaleY {
			wg.Add(1)
			go func(x, y, iterX, iterY int) {
				dali.canvas[y][x] = string(dali.ramp[len(dali.ramp)*dali.avgPixel(muse, x, y)/65536])
				wg.Done()
			}(x, y, x/dali.scaleY, y/dali.scaleX)
		}
	}
	wg.Wait()
	return dali.show()
}

func (dali Artist) getBlankCanvas(max image.Point) [][]string {
	cnvs := [][]string{}
	for len(cnvs) < max.Y {
		cnvs = append(cnvs, make([]string, max.X))
	}
	return cnvs
}

func (dali *Artist) show() string {
	rows := []string{}
	for _, row := range dali.canvas {
		rows = append(rows, strings.Join(row, ""))
	}
	return strings.Join(rows, "\n")
}

func (dali Artist) avgPixel(muse image.Image, x, y int) int {
	cnt, sum, max := 0, 0, muse.Bounds().Max
	for i := x; i < x+dali.scaleX && i < max.X; i++ {
		for j := y; j < y+dali.scaleY && j < max.Y; j++ {
			sum += grayscale(muse.At(i, j))
			cnt++
		}
	}
	return sum / cnt
}

func grayscale(c color.Color) int {
	r, g, b, _ := c.RGBA()
	return int(0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b))
}
