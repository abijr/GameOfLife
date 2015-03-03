package game

type cell struct {
	X, Y int
}

// neighbors find the live neighbor count of the given coordinates
func countNeighbors(c cell, w [][]bool) int {

	neighbors := 9 // represents the 3x3 block encompasing the given coordinates

	maxX, maxY := len(w[0])-1, len(w)-1 // The coordinates are used as w[y][x]

	neighborCount := 0

	for i := 0; i < neighbors; i++ {
		// Skip the center (current cell) of the 3x3 square
		if i == 4 {
			continue
		}

		// Calculate current loop coordinates
		x, y := (i%3-1)+c.X, (i/3-1)+c.Y

		// Skip cells out of bounds
		if x < 0 || y < 0 || x > maxX || y > maxY {
			continue
		}

		if w[y][x] {
			neighborCount++
		}
	}
	return neighborCount
}

// isAlive determines the status of the current cell
// for the next tick, i.e. given the current neighbors of the cell
// and the cells current status, is it going to be alive for next tick?
func isAlive(c cell, w [][]bool) bool {
	nc := countNeighbors(c, w) // neighbor count
	cs := w[c.Y][c.X]          // cell status

	// live cell with fewer than two live neighbours dies
	// live cell with two or three live neighbours lives
	// live cell with more than three live neighbours dies
	// dead cell with exactly three live neighbours becomes a live cell
	switch {
	case cs && (nc == 2):
		return true
	case nc == 3:
		return true
	default:
		return false
	}
}
