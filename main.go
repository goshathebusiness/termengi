package main

import (
	"os/exec"
	"time"

	"github.com/goshathebusiness/termengi/pkg/engine"
	"github.com/goshathebusiness/termengi/pkg/models"
)

const (
	FPS = 24
)

func main() {

	cmd := exec.Command("tput", "civis")
	cmd.Stdout = nil
	cmd.Run()

	e := engine.NewEngine()
	var items []interface{}
	items = append(items, &models.Rect{
		X:        20,
		Y:        10,
		W:        10,
		H:        10,
		Rotation: 0,
	})
	for {
		err := e.Draw()
		if err != nil {
			panic(err)
		}

		time.Sleep(time.Second / FPS)
	}
}
