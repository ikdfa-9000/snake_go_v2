package pkg

import (
	"fmt"
	"time"
)

func render(deskLinkVert *int, deskLinkHoriz *int, scoreLink *int, speedLink *float64, playgroundLink *[][]Symbol) {
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

func renderMenu(s *State, textPlay string, textOptions string, textExit string, score int) {
	fmt.Println("Змейка! :D")
	fmt.Println()
	fmt.Println(textPlay)
	fmt.Println(textOptions)
	fmt.Println(textExit)
	fmt.Println("-----------------")
	fmt.Println("Последнее кол-во очков: ", score)
	// TODO: Вывести рекорд
}

func clearRender() {
	fmt.Printf("\033[1A\033[K")
}
