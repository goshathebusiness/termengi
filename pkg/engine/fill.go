package engine

import (
	"fmt"
	"os"
	"os/exec"
)

const (
	BackgroundSymbol = " "
	ClearTerminal    = "\033[2J"
)

func BackgroundFill(w, h int) {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()

	for i := 0; i <= h; i++ {
		for j := 0; j <= w; j++ {
			ReplaceSymbolAtPosition(i, j, BackgroundSymbol)
		}
	}
}

func ReplaceSymbolAtPosition(row, column int, symbol string) {
	fmt.Printf("\033[%d;%df%s", row, column, symbol)
}
