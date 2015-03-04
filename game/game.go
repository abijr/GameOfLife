package game

import (
	"fmt"
	"strings"
)

// Player implements the Play method
type Player interface {
	Play(chan struct{}) chan Player
}

// A simple game of life 'Player'
type Simple struct {
	world [][]bool
}

func (s Simple) Play(done chan struct{}) chan Player {
	world := iterate(s.world)
	generations := make(chan Player)

	go func() {
		defer fmt.Println("Exiting goroutine.")
		for {
			select {
			case generations <- Simple{world}:
				world = iterate(world)
			case _, ok := <-done:
				if !ok {
					return
				}
			}
		}
	}()
	return generations
}

func (s Simple) String() string {
	rows := make([]string, len(s.world))
	for i, row := range s.world {
		rows[i] = fmt.Sprint(row)
	}

	return strings.Join(rows, "\n")

}

func New(w [][]bool) Player {
	return Simple{w}
}

func iterate(current [][]bool) [][]bool {
	xSize, ySize := len(current[0]), len(current)
	nextgen := make([][]bool, ySize)

	for y, row := range current {
		nextgen[y] = make([]bool, xSize)
		for x := range row {
			nextgen[y][x] = isAlive(cell{x, y}, current)
		}
	}

	return nextgen
}
