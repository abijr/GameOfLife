package game

import "testing"

var testWorld = [][]bool{
	[]bool{true, true, false, false, false},
	[]bool{true, false, false, true, false},
	[]bool{false, false, false, true, false},
	[]bool{false, false, true, true, true},
	[]bool{false, false, true, true, false},
}

func Test_countNeighbors(t *testing.T) {
	type neighNum struct {
		c   cell
		num int // number of neighbors
	}

	tests := []neighNum{
		neighNum{cell{1, 1}, 3},
		neighNum{cell{0, 0}, 2},
		neighNum{cell{3, 3}, 5},
	}

	for index, item := range tests {
		if countNeighbors(item.c, testWorld) != item.num {
			t.Errorf("Number of neighbors incorrect in test #%d.\n", index)
		}
	}
}

func Test_isAlive(t *testing.T) {
	type alive struct {
		c       cell
		isalive bool
	}

	tests := []alive{
		alive{cell{1, 1}, true},
		alive{cell{3, 3}, false},
	}

	for index, item := range tests {
		if isAlive(item.c, testWorld) != item.isalive {
			t.Errorf("Number of neighbors incorrect in test #%d.\n", index)
		}
	}

}
