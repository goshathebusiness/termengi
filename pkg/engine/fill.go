package engine

import (
	"fmt"
)

const (
	BackgroundSymbol = "."
)

func BackgroundFill(w, h int) {
	fmt.Print("\033[1000;1000H")
	for i := 0; i <= h; i++ {
		for j := 0; j <= w; j++ {
			ReplaceSymbolAtPosition(i, j, BackgroundSymbol)
		}
	}
}

func ReplaceSymbolAtPosition(row, column int, symbol string) {
	fmt.Printf("\033[%d;%df%s", row, column, symbol)
}
