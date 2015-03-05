package game

import "fmt"

// Player implements the Play method
type Player interface {
	Play(chan struct{}) chan Player
	World() [][]bool
}

// Simple is a simple game of life 'Player'
type Simple struct {
	world [][]bool
}

func (s Simple) World() [][]bool {
	return s.world
}

// Play plays the given 'Simple' game of life in
// a concurrent manner. Close 'done' channel to
// terminate goroutine.
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
	return fmt.Sprint(s.world)

}

// New returns a new Game Of Life 'Player'
// using the given world as a starting point
func New(world [][]bool) Player {
	return Simple{world}
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
