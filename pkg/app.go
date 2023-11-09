package pkg

import (
	"fmt"
	"math/rand"
	"os"
	"strings"

	"example.com/snake_go/pkg/cfg"
	"example.com/snake_go/pkg/themes"

	tb "github.com/nsf/termbox-go"
)

func Run(config cfg.Config) {
	state := initState()
	snakeDirectionHorizontal := 1
	snakeDirectionVertical := 0
	//score := 0
	keyboardErr := tb.Init()
	if keyboardErr != nil {
		panic(keyboardErr)
	}
	defer tb.Close()
	go readKey(&snakeDirectionHorizontal, &snakeDirectionVertical, &state)
	for {
		switch state.status {
		case menuActiveStatus, gameOverStatus:
			renderMenu(state.menuStrings[0], state.menuStrings[1], state.menuStrings[2], state.score)
		case gameActiveStatus, gamePauseStatus:
			rows := config.DeskRows
			columns := config.DeskColumns
			state.snake.length = 2
			frameSpeed := float64(config.DeskFrameSpeed)
			// Создаём двумерный слайс
			playground := make([][]themes.Symbol, rows)

			playgroundRow := GetPlaygroudRow(columns, state.space)
			for i := range playground {
				row := make([]themes.Symbol, columns)
				copy(row, playgroundRow)
				playground[i] = row
			}

			appleCoord := make([]int, 2)
			state.ResetScore()
			snakeDirectionHorizontal = 1
			snakeDirectionVertical = 0
			appleScoreAdd := 100
			snakeCord := make([][]int, 3) // Двумерный слайс змейки, каждый слайс содержит вертикальную и горизонтальную координату
			for i := range snakeCord {
				snakeCord[i] = make([]int, 2)
			}
			// Задаём координаты голове змеи и яблока
			snakeCord[0][0] = columns / 2
			snakeCord[0][1] = rows / 2
			appleCoord[0] = rand.Intn(columns-1) + 0
			appleCoord[1] = rand.Intn(rows-1) + 0
			// Если координаты головы змеи совпадают с яблоком, то перемещаем яблоко
			for playground[appleCoord[1]][appleCoord[0]] != state.space {
				appleCoord[0] = rand.Intn(columns-1) + 0
				appleCoord[1] = rand.Intn(rows-1) + 0
			}
			playground[snakeCord[0][1]][snakeCord[0][0]] = state.snake.headSymbol
			playground[snakeCord[1][1]][snakeCord[1][0]] = state.snake.symbol
			playground[snakeCord[2][1]][snakeCord[2][0]] = state.snake.symbol
			playground[appleCoord[1]][appleCoord[0]] = state.apple
			for { // for {} == while True. Постоянный цикл
				if state.status == gamePauseStatus {
					render(&rows, &columns, &state.score, &frameSpeed, &playground)
					if state.status == gameOverStatus {
						for k := 0; k < rows+1; k++ {
							clearRender()
						}
						fmt.Println(colorRed, "Игра окончена")
						fmt.Print(colorReset)
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
					// Чёта тут я проверял... сейчас вроде нормально
					// if playground[snakeCord[state.snake.length][1]][snakeCord[state.snake.length][0]] != state.apple {
					playground[snakeCord[state.snake.length][1]][snakeCord[state.snake.length][0]] = state.space
					// }
					// Смотрим, выходит ли змейка за рамки
					if snakeCord[0][1]+snakeDirectionVertical == -1 || snakeCord[0][1]+snakeDirectionVertical == rows || snakeCord[0][0]+snakeDirectionHorizontal == -1 || snakeCord[0][0]+snakeDirectionHorizontal == columns {
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
					if snakeCord[0][0] == appleCoord[0] && snakeCord[0][1] == appleCoord[1] {
						state.snake.length = state.snake.length + 1
						state.AddScore(appleScoreAdd)
						snakeCordAdd := []int{snakeCord[state.snake.length-1][1] - snakeDirectionVertical, snakeCord[state.snake.length-1][0] - snakeDirectionHorizontal}
						snakeCord = append(snakeCord, snakeCordAdd)
						appleCoord[0] = rand.Intn(columns-1) + 0
						appleCoord[1] = rand.Intn(rows-1) + 0
						for i := 0; i < len(snakeCord); i++ {
							for playground[appleCoord[1]][appleCoord[0]] != state.space {
								// Если новые координаты яблока совпадают с телом змеи, то яблоко нужно пересоздать
								appleCoord[0] = rand.Intn(columns-1) + 0
								appleCoord[1] = rand.Intn(rows-1) + 0
							}
						}
						playground[appleCoord[1]][appleCoord[0]] = state.apple
					}
					if state.status == gameOverStatus {
						for k := 0; k < rows+1; k++ {
							clearRender()
						}
						fmt.Println(colorRed, "Игра окончена")
						fmt.Print(colorReset)
						state.status = menuActiveStatus
						break
					} else {
						fmt.Println(appleCoord[1], appleCoord[0])
						render(&rows, &columns, &state.score, &frameSpeed, &playground)
					}
				}
			}
		case gameExitStatus:
			clearRender()
			os.Exit(0)
		}
	}
}

func initState() State {
	return State{
		menuStatusId: menuIdPlay,
		menuStrings:  []string{">Играть", " Опции", " Выход"},
		status:       menuActiveStatus, // TODO: temp
		apple:        themes.SymbolRed,
		space:        themes.SymbolGreen,
		score:        0,
		snake: Snake{
			length:     2,
			symbol:     themes.SymbolPurple,
			headSymbol: themes.SymbolWhiteCircle,
		},
	}
}

func readKey(horizAddress *int, vertAddress *int, state *State) {
	for {
		if state.status == gameExitStatus {
			break
		}
		event := tb.PollEvent()

		switch state.status {
		case gamePauseStatus:
			// Во время паузы игрок не должен иметь возможность менять направление змейки
			switch {
			case event.Ch == 'p':
				state.togglePause()
			case event.Key == tb.KeyCtrlC || event.Key == tb.KeyEsc:
				state.gameOver()
			}
		case menuActiveStatus:
			switch {
			case event.Key == tb.KeyArrowUp:
				if state.menuStatusId == menuIdPlay {
					state.menuStatusId = menuIdExit
					state.menuStrings[menuIdPlay] = strings.Replace(state.menuStrings[menuIdPlay], ">", " ", -1)
				} else {
					state.menuStatusId--
					state.menuStrings[state.menuStatusId+1] = strings.Replace((state.menuStrings[state.menuStatusId+1]), ">", " ", -1)
				}
				state.menuStrings[state.menuStatusId] = strings.Replace(state.menuStrings[state.menuStatusId], " ", ">", -1)
			case event.Key == tb.KeyArrowDown:
				if state.menuStatusId == menuIdExit {
					state.menuStatusId = menuIdPlay
					state.menuStrings[menuIdExit] = strings.Replace(state.menuStrings[menuIdExit], ">", " ", -1)
				} else {
					state.menuStatusId++
					state.menuStrings[state.menuStatusId-1] = strings.Replace(state.menuStrings[state.menuStatusId-1], ">", " ", -1)
				}
				state.menuStrings[state.menuStatusId] = strings.Replace(state.menuStrings[state.menuStatusId], " ", ">", -1)
			case event.Key == tb.KeyEnter:
				// Цикл for для clearRender() почему-то помечает дальнейший код как недостижимый... хз
				clearRender()
				clearRender()
				clearRender()
				clearRender()
				clearRender()
				clearRender()
				clearRender()
				clearRender()
				state.menuInitialize()
			case event.Key == tb.KeyEsc || event.Key == tb.KeyCtrlC:
				state.status = gameExitStatus
			}
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
