package controls

import (
	"os"

	"github.com/nsf/termbox-go"

	"github.com/goshathebusiness/termengi/pkg/models"
)

const (
	MoveStep = 1
)

func HandleKeyboard(rect *models.Rect) {
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			rect.Lock()

			switch ev.Key {
			case termbox.KeyArrowUp:
				rect.Y -= MoveStep
			case termbox.KeyArrowDown:
				rect.Y += MoveStep
			case termbox.KeyArrowLeft:
				rect.X -= MoveStep
			case termbox.KeyArrowRight:
				rect.X += MoveStep

			case termbox.KeyCtrlC:
				os.Exit(0)
			}

			rect.Unlock()
		case termbox.EventError:
			panic(ev.Err)
		}
	}
}
