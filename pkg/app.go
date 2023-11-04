package pkg

import (
	"fmt"
	tb "github.com/nsf/termbox-go"
	"math/rand"
)

type Symbol string

const (
	red    Symbol = "🔴"
	green         = "🟢"
	purple        = "🟣"
)

func Run(config Config) {
	state := initState()
	state.menuStringID = 1
	// Всю остальную часть кода заключаем в for и проверяем isGameNeeded, isGamePaused, isMenuPaused, isGameOver.
	// Вот не очень хочется впихивать в for кучу if else, подумаю и чё-нибудь замучу... блинб

	rows := config.deskRows
	columns := config.deskColumns
	frameSpeed := float64(config.deskFrameSpeed)
	// Создаём двумерный слайс
	playground := make([][]Symbol, rows)
	// TODO: доделать
	plRow := []Symbol{green, green, green, green, green, green, green, green}
	for i := range playground {
		row := make([]Symbol, columns)
		copy(row, plRow)
		playground[i] = row
	}

	appleCoord := make([]int, 2)
	snakeDirectionHorizontal := 1
	snakeDirectionVertical := 0
	score := 0
	appleScoreAdd := 100
	snakeCord := make([][]int, 3) // Двумерный слайс змейки, каждый слайс содержит вертикальную и горизонтальную координату
	for i := range snakeCord {
		snakeCord[i] = make([]int, 2)
	}
	keyboardErr := tb.Init()
	if keyboardErr != nil {
		panic(keyboardErr)
	}
	defer tb.Close()
	// Задаём координаты голове змеи и яблока
	snakeCord[0][0] = columns / 2
	snakeCord[0][1] = rows / 2
	appleCoord[0] = rand.Intn(rows-1) + 0
	appleCoord[1] = rand.Intn(columns-1) + 0
	// Если координаты головы змеи совпадают с яблоком, то перемещаем яблоко
	for appleCoord[0] == snakeCord[0][0] && appleCoord[1] == snakeCord[0][1] {
		appleCoord[0] = rand.Intn(rows-1) + 0
		appleCoord[1] = rand.Intn(columns-1) + 0
	}
	playground[snakeCord[0][1]][snakeCord[0][0]] = state.snake.headSymbol
	playground[snakeCord[1][1]][snakeCord[1][0]] = state.snake.symbol
	playground[snakeCord[2][1]][snakeCord[2][0]] = state.snake.symbol
	playground[appleCoord[0]][appleCoord[1]] = state.apple
	go readKey(&snakeDirectionHorizontal, &snakeDirectionVertical, &state)
	for { // for {} == while True. Постоянный цикл
		if state.status == gamePauseStatus {
			fmt.Println(appleCoord[1], appleCoord[0])
			render(&rows, &columns, &score, &frameSpeed, &playground)
			if state.status == gameOverStatus {
				for k := 0; k < rows+1; k++ {
					clearRender()
				}
				fmt.Println("Game Over")
				break
			}
		} else {
			// Координаты каждой клетки змейки кроме первой приравниваем к предыдущей
			// Первую клетку двигаем вперёд
			// Отрисовываем каждую клетку
			for i := 0; i < state.snake.length; i++ {
				snakeCord[state.snake.length-i][0], snakeCord[state.snake.length-i][1] = snakeCord[state.snake.length-i-1][0], snakeCord[state.snake.length-i-1][1]
				playground[snakeCord[state.snake.length-i][1]][snakeCord[state.snake.length-i][0]] = state.snake.symbol
			}
			// Проверочка, чтобы в частных случаях иконка яблока не пропадала
			if playground[snakeCord[state.snake.length][1]][snakeCord[state.snake.length][0]] != state.apple {
				playground[snakeCord[state.snake.length][1]][snakeCord[state.snake.length][0]] = state.space
			}
			// Смотрим, выходит ли змейка за рамки
			if snakeCord[0][1]+snakeDirectionVertical == -1 || snakeCord[0][1]+snakeDirectionVertical == rows || snakeCord[0][0]+snakeDirectionHorizontal == -1 || snakeCord[0][0]+snakeDirectionHorizontal == columns {
				// gameOver = true
				switch snakeDirectionVertical {
				case 1:
					snakeCord[0][1] = -1
				case -1:
					snakeCord[0][1] = rows
				default:
					switch snakeDirectionHorizontal { // Такого уродства нет даже в погребе у Сатаны
					case 1:
						snakeCord[0][0] = -1
					case -1:
						snakeCord[0][0] = columns
					}
				}
			}
			if state.status == gameActiveStatus {
				snakeCord[0][1], snakeCord[0][0] = snakeCord[0][1]+snakeDirectionVertical, snakeCord[0][0]+snakeDirectionHorizontal
			}
			// Смотрим, врезается ли змейка или нет
			if playground[snakeCord[0][1]][snakeCord[0][0]] == state.snake.symbol {
				state.gameOver()
			}
			playground[snakeCord[0][1]][snakeCord[0][0]] = state.snake.headSymbol
			// Захавал яблоко. Делаем новое
			if snakeCord[0][1] == appleCoord[0] && snakeCord[0][0] == appleCoord[1] {
				state.snake.length = state.snake.length + 1
				score = score + appleScoreAdd
				snakeCordAdd := []int{snakeCord[state.snake.length-1][1] - snakeDirectionVertical, snakeCord[state.snake.length-1][0] - snakeDirectionHorizontal}
				snakeCord = append(snakeCord, snakeCordAdd)
				appleCoord[0] = rand.Intn(rows-1) + 0
				appleCoord[1] = rand.Intn(columns-1) + 0
				for i := 0; i < state.snake.length; i++ {
					for appleCoord[1] == snakeCord[i][0] && appleCoord[0] == snakeCord[i][1] {
						// Если новые координаты яблока совпадают с телом змеи, то яблоко нужно пересоздать
						// TODO: Исправить баг, спавнящий яблоко на змейке
						appleCoord[0] = rand.Intn(rows-1) + 0
						appleCoord[1] = rand.Intn(columns-1) + 0
					}
				}
				playground[appleCoord[0]][appleCoord[1]] = state.apple
			}
			if state.status == gameOverStatus {
				for k := 0; k < rows+1; k++ {
					clearRender()
				}
				fmt.Println("Game Over")
				break
			} else {
				fmt.Println(appleCoord[1], appleCoord[0])
				render(&rows, &columns, &score, &frameSpeed, &playground)
			}
		}
	}
}

func initState() State {
	return State{
		status:       gameActiveStatus, // TODO: temp
		apple:        red,
		space:        green,
		menuStrings:  []string{" Играть", " Опции", " Выход"},
		menuStringID: 1,
		snake: Snake{
			length:     2,
			symbol:     purple,
			headSymbol: purple,
		},
	}
}

// Считывает нажатия с клавиатуры и изменяет направление змейки.
// Завершает игру, если пользователь хочет выйти
func readKey(horizAddress *int, vertAddress *int, state *State) {
	// Не знаю, насколько целесообразно использовать сразу для трёх (а с Options для целых четырех) стейтов лишь одну
	// функцию для чтения инпута с клавиатуры.

	for {
		event := tb.PollEvent()

		switch state.status {
		case gameExitStatus:
		case gamePauseStatus:
			// Во время паузы игрок не должен иметь возможность менять направление змейки
			switch {
			case event.Ch == 'p':
				state.togglePause()
			case event.Key == tb.KeyCtrlC || event.Key == tb.KeyEsc:
				state.gameOver()
			}

		//case gameOverStatus:
		//case menuActiveStatus:
		//	switch {
		//	case event.Key == tb.KeyArrowUp:
		//		if state.menuStringID == 1 {
		//			state.menuStringID = 3
		//		}
		//	case event.Key == tb.KeyArrowDown:
		//		if state.menuStringID == 3 {
		//			state.menuStringID = 1
		//		}
		//	case event.Key == tb.KeyEnter:
		//		// Цикл for почему-то помечает дальнейший код как недостижимый... хз
		//		fmt.Printf("\033[1A\033[K")
		//		fmt.Printf("\033[1A\033[K")
		//		fmt.Printf("\033[1A\033[K")
		//		fmt.Printf("\033[1A\033[K")
		//		fmt.Printf("\033[1A\033[K")
		//		fmt.Printf("\033[1A\033[K")
		//		fmt.Printf("\033[1A\033[K")
		//		state.menuInitialize()
		//	}
		case gameActiveStatus:
			switch {
			case event.Ch == 'p':
				state.togglePause()
			case event.Key == tb.KeyCtrlC || event.Key == tb.KeyEsc:
				state.gameOver()

			case event.Ch == 'a' || event.Key == tb.KeyArrowLeft:
				if *horizAddress == 0 {
					*horizAddress = -1
					*vertAddress = 0
				}
			case event.Ch == 's' || event.Key == tb.KeyArrowDown:
				if *vertAddress == 0 {
					*horizAddress = 0
					*vertAddress = 1
				}
			case event.Ch == 'd' || event.Key == tb.KeyArrowRight:
				if *horizAddress == 0 {
					*horizAddress = 1
					*vertAddress = 0
				}
			case event.Ch == 'w' || event.Key == tb.KeyArrowUp:
				if *vertAddress == 0 {
					*horizAddress = 0
					*vertAddress = -1
				}
			}
		}
	}
}
