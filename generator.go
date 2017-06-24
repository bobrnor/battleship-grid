package grid

import (
	"math/rand"
	"time"

	"github.com/pkg/errors"
)

var (
	sizes = []int{4, 3, 3, 2, 2, 2, 1, 1, 1, 1}
)

type generator struct {
	ships []*ship
	sizes []int
}

func Generate() (*Grid, error) {
	rand.Seed(time.Now().UnixNano())

	g := generator{}

	g.randomizeSizes()
	if err := g.generateArrangement(); err != nil {
		return nil, err
	}

	return g.grid(), nil
}

func (g *generator) randomizeSizes() {
	g.sizes = []int{}
	randIndexes := rand.Perm(10)
	for _, index := range randIndexes {
		g.sizes = append(g.sizes, sizes[index])
	}
}

func (g *generator) generateArrangement() error {
	g.initArrangement()
	if err := g.findValidArrangement(); err != nil {
		return err
	}
	return nil
}

func (g *generator) initArrangement() {
	g.ships = []*ship{}
	for i := 0; i < 10; i++ {
		size := g.sizes[i]
		g.ships = append(g.ships, newShip(size))
	}
}

func (g *generator) findValidArrangement() error {
	for i := 0; i < 10; {
		if g.isSubArrangementValid(i) {
			i++
		} else {
			for !g.findNextPosition(i) {
				if i > 0 {
					i--
				} else {
					return errors.Errorf("Can't find valid arrangement:\n%+v", g)
				}
			}
		}
	}
	return nil
}

func (g *generator) isSubArrangementValid(lastShipIndex int) bool {
	ship := g.ships[lastShipIndex]

	if !ship.isValid() {
		return false
	}

	for i := lastShipIndex - 1; i >= 0; i-- {
		s := g.ships[i]
		if s.intersect(ship) {
			return false
		}
	}

	return true
}

func (g *generator) findNextPosition(shipIndex int) bool {
	ship := g.ships[shipIndex]
	return ship.findNextPosition()
}

func (g *generator) grid() *Grid {
	var grid Grid
	for _, ship := range g.ships {
		for x := uint(ship.Left); x <= uint(ship.Right); x++ {
			for y := uint(ship.Top); y <= uint(ship.Bottom); y++ {
				grid.Set(x, y)
			}
		}
	}
	return &grid
}
