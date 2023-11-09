package pkg

import (
	"example.com/snake_go/pkg/themes"
	"fmt"
	"time"
)

// TODO: перенести в themes.colors
const ()

func render(deskLinkVert *int, deskLinkHoriz *int, scoreLink *int, speedLink *float64, playgroundLink *[][]themes.Symbol) {
	for k := 0; k < *deskLinkVert; k++ { // Вывод матрицы в терминал
		for l := 0; l < *deskLinkHoriz; l++ {
			fmt.Print((*playgroundLink)[k][l])
		}
		fmt.Println()
	}
	fmt.Println("------------------")
	fmt.Println(themes.ColorGreen, "Очки: ", *scoreLink)
	fmt.Print(themes.ColorReset)
	time.Sleep(time.Duration(*speedLink) * time.Millisecond)
	for k := 0; k < *deskLinkVert+3; k++ {
		fmt.Printf("\033[1A\033[K")
	}
}

func renderMenu(textPlay string, textOptions string, textExit string, score int) {
	fmt.Println("\033[33m", "Змейка! :D")
	fmt.Println(themes.ColorReset)
	fmt.Println(textPlay)
	fmt.Println(textOptions)
	fmt.Println(textExit)
	fmt.Println("-----------------")
	fmt.Println(themes.ColorGreen, "Последнее кол-во очков: ", score)
	fmt.Print(themes.ColorReset)
	// TODO: Вывести рекорд
	time.Sleep(time.Duration(100) * time.Millisecond)
	for i := 0; i < 7; i++ {
		clearRender()
	}
}

func clearRender() {
	fmt.Printf("\033[1A\033[K")
}

func GetPlaygroudRow(columns int, symbol themes.Symbol) []themes.Symbol {
	var row []themes.Symbol
	for i := 0; i <= columns; i++ {
		row = append(row, symbol)
	}

	return row
}
