package main

import (
	"os/exec"
	"time"

	termbox "github.com/nsf/termbox-go"

	"github.com/goshathebusiness/termengi/pkg/controls"
	"github.com/goshathebusiness/termengi/pkg/engine"
	"github.com/goshathebusiness/termengi/pkg/models"
)

const (
	FPS = 12
)

func main() {

	cmd := exec.Command("tput", "civis")
	cmd.Stdout = nil
	cmd.Run()

	if err := termbox.Init(); err != nil {
		panic(err)
	}
	defer termbox.Close()

	e := engine.NewEngine()
	rect := models.Rect{
		X:        20,
		Y:        10,
		W:        10,
		H:        10,
		Rotation: 0,
	}
	e.Items = append(e.Items, &rect)

	// Polling for keyboard events
	controls.HandleKeyboard(&rect)

	for {
		err := e.Draw()
		if err != nil {
			panic(err)
		}

		time.Sleep(time.Second / FPS)
	}
}
