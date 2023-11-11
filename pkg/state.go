package pkg

import "example.com/snake_go/pkg/themes"

type StateStatus int
type MenuStatus int

const (
	gameExitStatus   StateStatus = -2
	gameOverStatus               = -1
	menuActiveStatus             = 0
	gameActiveStatus             = 1
	gamePauseStatus              = 2
)

const (
	menuIdPlay    MenuStatus = 0
	menuIdOptions            = 1
	menuIdExit               = 2
)

const (
	noDirection         int = 0
	directionHorizLeft      = -1
	directionHorizRight     = 1
	directionVertUp         = -1
	directionVertDown       = 1
)

type State struct {
	status       StateStatus
	menuStatusId MenuStatus
	canChangeDir bool
	// TODO
	menuStrings []string

	apple themes.Symbol
	space themes.Symbol
	snake Snake

	score int
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

func (s *State) menuInitialize() {
	switch s.menuStatusId {
	case menuIdPlay:
		s.status = gameActiveStatus
	case menuIdOptions:
	case menuIdExit:
		s.status = gameExitStatus
	}
}

func (s *State) ResetScore() {
	s.score = 0
}

func (s *State) AddScore(add int) {
	s.score += add
}

// TODO
//func (s *State) menuSummon() {
//	s.isMenuActive = true
//	if s.isGameOver {
//		println("Ты проиграл!")
//	}
//}

type Snake struct {
	length     int
	symbol     themes.Symbol
	headSymbol themes.Symbol
}
