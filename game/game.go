package game

// Player implements the Play method
type Player interface {
	Play(chan [][]bool)
}

type Simple struct {
	world [][]bool
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

func (s Simple) Play(c chan [][]bool) {
	world := s.world

	for {
		world = iterate(world)
		c <- world

	}
}

func New(w [][]bool) Player {
	return Simple{w}
}
