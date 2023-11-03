package pkg

type State struct {
	isGameOver   bool
	isGamePaused bool
	isMenuActive bool
	apple        Symbol
	space        Symbol
	snake        Snake
}

func (s *State) gameOver() {
	s.isGameOver = true
}

func (s *State) togglePause() {
	s.isGamePaused = !(s.isGamePaused)
}

func (s *State) menuSummon() {
	s.isMenuActive = true
	if s.isGameOver {
		println("Ты проиграл!")
	}
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
