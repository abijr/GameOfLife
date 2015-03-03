package game

// Player implements the Play method
type Player interface {
	Play(chan [][]bool)
}

type Naive struct {
	world [][]bool
}

func New([][]bool) Player {

}
