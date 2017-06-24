package grid

import "fmt"

type Grid [13]byte

func Diff(g1 *Grid, g2 *Grid) *Grid {
	result := &Grid{}
	for i := uint(0); i < 13; i++ {
		result[i] = g1[i] - (g1[i] & g2[i])
	}
	return result
}

func (g *Grid) Set(x, y uint) {
	pos := 10*y + x
	byteIdx := pos / 8
	bitIdx := pos % 8
	g[byteIdx] = g[byteIdx] | (1 << bitIdx)
}

func (g *Grid) Get(x, y uint) bool {
	pos := 10*y + x
	byteIdx := pos / 8
	bitIdx := pos % 8
	return g[byteIdx]&(1<<bitIdx) != 0
}

func (g *Grid) IsEmpty() bool {
	for _, byte := range g {
		if byte != 0 {
			return false
		}
	}
	return true
}

func (g *Grid) Print() {
	for x := uint(0); x < 10; x++ {
		for y := uint(0); y < 10; y++ {
			if g.Get(x, y) {
				fmt.Printf(" ◉ ")
			} else {
				fmt.Printf(" ○ ")
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n\n")
}
