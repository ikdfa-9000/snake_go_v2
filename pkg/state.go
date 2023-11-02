package pkg

type State struct {
	isGameOver bool
	apple      Symbol
	space      Symbol
	snake      Snake
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
