package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	tb "github.com/nsf/termbox-go"
)

func main() {
	input, err := os.Open("config.txt") // Открытие файла
	if err != nil {
		println(err)
		log.Fatal("Файла нет")
	}
	defer input.Close()
	configFile := bufio.NewScanner(input) // Инициализация сканера.
	configFile.Scan()
	deskSizeVert, err := strconv.Atoi(configFile.Text())
	if err != nil {
		log.Fatal("Хопа! А я не могу прочитать, что написано в конфиг файле")
	}
	configFile.Scan()
	deskSizeHoriz, err := strconv.Atoi(configFile.Text())
	if err != nil {
		log.Fatal("Хопа! А я не могу прочитать, что написано в конфиг файле")
	}
	configFile.Scan()
	frameSpeed, errSpeed := strconv.ParseFloat(configFile.Text(), 32)
	if errSpeed != nil {
		log.Fatal("Хопа! А я не могу прочитать, что написано в конфиг файле")
	}
	// Создаём двумерный слайс
	playground := make([][]string, deskSizeVert)
	for i := range playground {
		playground[i] = make([]string, deskSizeHoriz)
	}
	appleSymbol := "🔴 "
	spaceSymbol := "🟢"
	snakeSymbol := "🟣"
	snakeHeadSymbol := "🟣"
	appleCord := make([]int, 2)
	snakeDirectionHorizontal := 1
	snakeDirectionVertical := 0
	gameOver := false
	snakeLength := 2
	score := 0
	appleScoreAdd := 100
	// Заполняем двумерный слайс, сначала пробелы
	for i := 0; i < deskSizeVert; i++ {
		for j := 0; j < deskSizeHoriz; j++ {
			playground[i][j] = spaceSymbol
		}
	}
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
	snakeCord[0][0] = deskSizeHoriz / 2
	snakeCord[0][1] = deskSizeVert / 2
	appleCord[0] = rand.Intn(deskSizeVert-1) + 0
	appleCord[1] = rand.Intn(deskSizeHoriz-1) + 0
	// Если координаты головы змеи совпадают с яблоком, то перемещаем яблоко
	for appleCord[0] == snakeCord[0][0] && appleCord[1] == snakeCord[0][1] {
		appleCord[0] = rand.Intn(deskSizeVert-1) + 0
		appleCord[1] = rand.Intn(deskSizeHoriz-1) + 0
	}
	playground[snakeCord[0][1]][snakeCord[0][0]] = snakeHeadSymbol
	playground[snakeCord[1][1]][snakeCord[1][0]] = snakeSymbol
	playground[snakeCord[2][1]][snakeCord[2][0]] = snakeSymbol
	playground[appleCord[0]][appleCord[1]] = appleSymbol
	go readKey(&snakeDirectionHorizontal, &snakeDirectionVertical)
	for { // for {} == while True. Постоянный цикл
		// Координаты каждой клетки змейки кроме первой приравниваем к предыдущей
		// Первую клетку двигаем вперёд
		// Отрисовываем каждую клетку
		for i := 0; i < snakeLength; i++ {
			snakeCord[snakeLength-i][0], snakeCord[snakeLength-i][1] = snakeCord[snakeLength-i-1][0], snakeCord[snakeLength-i-1][1]
			playground[snakeCord[snakeLength-i][1]][snakeCord[snakeLength-i][0]] = snakeSymbol
		}
		// Проверочка, чтобы в частных случаях иконка яблока не пропадала
		if playground[snakeCord[snakeLength][1]][snakeCord[snakeLength][0]] != appleSymbol {
			playground[snakeCord[snakeLength][1]][snakeCord[snakeLength][0]] = spaceSymbol
		}
		// Смотрим, выходит ли змейка за рамки
		if snakeCord[0][1]+snakeDirectionVertical == -1 || snakeCord[0][1]+snakeDirectionVertical == deskSizeVert || snakeCord[0][0]+snakeDirectionHorizontal == -1 || snakeCord[0][0]+snakeDirectionHorizontal == deskSizeHoriz {
			// gameOver = true
			switch snakeDirectionVertical {
			case 1:
				snakeCord[0][1] = -1
			case -1:
				snakeCord[0][1] = deskSizeVert
			default:
				switch snakeDirectionHorizontal { // Такого уродства нет даже в погребе у Сатаны
				case 1:
					snakeCord[0][0] = -1
				case -1:
					snakeCord[0][0] = deskSizeHoriz
				}
			}
		}
		if !gameOver {
			snakeCord[0][1], snakeCord[0][0] = snakeCord[0][1]+snakeDirectionVertical, snakeCord[0][0]+snakeDirectionHorizontal
		}
		// Смотрим, врезается ли змейка или нет
		if playground[snakeCord[0][1]][snakeCord[0][0]] == snakeSymbol {
			gameOver = true
		}
		playground[snakeCord[0][1]][snakeCord[0][0]] = snakeHeadSymbol
		// Захавал яблоко. Делаем новое
		if snakeCord[0][1] == appleCord[0] && snakeCord[0][0] == appleCord[1] {
			snakeLength = snakeLength + 1
			score = score + appleScoreAdd
			snakeCordAdd := []int{snakeCord[snakeLength-1][1] - snakeDirectionVertical, snakeCord[snakeLength-1][0] - snakeDirectionHorizontal}
			snakeCord = append(snakeCord, snakeCordAdd)
			appleCord[0] = rand.Intn(deskSizeVert-1) + 0
			appleCord[1] = rand.Intn(deskSizeHoriz-1) + 0
			for i := 0; i < snakeLength; i++ {
				for appleCord[1] == snakeCord[i][0] && appleCord[0] == snakeCord[i][1] {
					// Если новые координаты яблока совпадают с телом змеи, то яблоко нужно пересоздать
					appleCord[0] = rand.Intn(deskSizeVert-1) + 0
					appleCord[1] = rand.Intn(deskSizeHoriz-1) + 0
				}
			}
			playground[appleCord[0]][appleCord[1]] = appleSymbol
		}
		if gameOver {
			for k := 0; k < deskSizeVert+1; k++ {
				fmt.Printf("\033[1A\033[K")
			}
			fmt.Println("Game Over")
			break
		} else {
			fmt.Println(appleCord[1], appleCord[0])
			render(&deskSizeVert, &deskSizeHoriz, &score, &frameSpeed, &playground)
		}
	}
}

func readKey(horizAddress *int, vertAddress *int) { // Чтение инпута с клавиатуры. Ненавижу
	for {
		event := tb.PollEvent()
		switch {
		case event.Ch == 'a':
			if *horizAddress == 0 {
				*horizAddress = -1
				*vertAddress = 0
			}
		case event.Ch == 's':
			if *vertAddress == 0 {
				*horizAddress = 0
				*vertAddress = 1
			}
		case event.Ch == 'd':
			if *horizAddress == 0 {
				*horizAddress = 1
				*vertAddress = 0
			}
		case event.Ch == 'w':
			if *vertAddress == 0 {
				*horizAddress = 0
				*vertAddress = -1
			}
		}
	}
}

func render(deskLinkVert *int, deskLinkHoriz *int, scoreLink *int, speedLink *float64, playgroundLink *[][]string) {
	for k := 0; k < *deskLinkVert; k++ { // Вывод матрицы в терминал
		for l := 0; l < *deskLinkHoriz; l++ {
			fmt.Print((*playgroundLink)[k][l])
		}
		fmt.Println()
	}
	fmt.Println("_____________________")
	fmt.Println("Score: ", *scoreLink)
	time.Sleep(time.Duration(*speedLink) * time.Millisecond)
	for k := 0; k < *deskLinkVert+3; k++ {
		fmt.Printf("\033[1A\033[K")
	}
}
