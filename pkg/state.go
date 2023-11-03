package pkg

type State struct {
	isGameOver bool
	apple      Symbol
	space      Symbol
	snake      Snake
}

func (s *State) gameOver() {
	s.isGameOver = true
}

type Snake struct {
	length     int
	symbol     Symbol
	headSymbol Symbol
}

type AppleCoord struct {
	x int
	y int
}
