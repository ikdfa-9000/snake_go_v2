package pkg

import (
	"fmt"
	"time"
)

const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorPurple = "\033[35m"
	colorCyan   = "\033[36m"
	colorWhite  = "\033[37m"
)

func render(deskLinkVert *int, deskLinkHoriz *int, scoreLink *int, speedLink *float64, playgroundLink *[][]Symbol) {
	for k := 0; k < *deskLinkVert; k++ { // Вывод матрицы в терминал
		for l := 0; l < *deskLinkHoriz; l++ {
			fmt.Print((*playgroundLink)[k][l])
		}
		fmt.Println()
	}
	fmt.Println("------------------")
	fmt.Println(colorGreen, "Очки: ", *scoreLink)
	fmt.Print(colorReset)
	time.Sleep(time.Duration(*speedLink) * time.Millisecond)
	for k := 0; k < *deskLinkVert+3; k++ {
		fmt.Printf("\033[1A\033[K")
	}
}

func renderMenu(textPlay string, textOptions string, textExit string, score int) {
	fmt.Println("\033[33m", "Змейка! :D")
	fmt.Println(colorReset)
	fmt.Println(textPlay)
	fmt.Println(textOptions)
	fmt.Println(textExit)
	fmt.Println("-----------------")
	fmt.Println(colorGreen, "Последнее кол-во очков: ", score)
	fmt.Print(colorReset)
	// TODO: Вывести рекорд
	time.Sleep(time.Duration(100) * time.Millisecond)
	for i := 0; i < 7; i++ {
		clearRender()
	}
}

func clearRender() {
	fmt.Printf("\033[1A\033[K")
}
