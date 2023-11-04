package pkg

type StateStatus int

const (
	gameExitStatus   StateStatus = -2
	gameOverStatus               = -1
	menuActiveStatus             = 0
	gameActiveStatus             = 1
	gamePauseStatus              = 2
)

type State struct {
	status StateStatus

	// TODO
	menuStrings  []string
	menuStringID int

	apple Symbol
	space Symbol
	snake Snake
}

func (s *State) gameActive() {
	s.status = gameActiveStatus
}

func (s *State) gameOver() {
	s.status = gameOverStatus
}

func (s *State) gamePause() {
	s.status = gamePauseStatus
}

func (s *State) togglePause() {
	if s.status == gameActiveStatus {
		s.gamePause()
	} else if s.status == gamePauseStatus {
		s.gameActive()
	}
}

// TODO
//func (s *State) menuSummon() {
//	s.isMenuActive = true
//	if s.isGameOver {
//		println("Ты проиграл!")
//	}
//}

// TODO
//func (s *State) menuInitialize() {
//	switch s.menuStringID {
//	case 1:
//		s.isMenuActive = false
//	case 3:
//		s.isGameNeeded = false
//		// TODO: Меню опций
//	}
//}

type Snake struct {
	length     int
	symbol     Symbol
	headSymbol Symbol
}
