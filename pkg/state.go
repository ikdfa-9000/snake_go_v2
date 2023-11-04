package pkg

type State struct {
	isGameOver   bool
	isGamePaused bool
	isMenuActive bool
	isGameNeeded bool

	menuStrings  []string
	menuStringID int

	apple Symbol
	space Symbol
	snake Snake
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

func (s *State) menuInitialize() {
	switch s.menuStringID {
	case 1:
		s.isMenuActive = false
	case 3:
		s.isGameNeeded = false
		// TODO: Меню опций
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
