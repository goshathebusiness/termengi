package engine

import (
	"github.com/goshathebusiness/termengi/pkg/terminal"
)

type Engine struct {
	items []*interface{}
}

func NewEngine() *Engine {
	return &Engine{items: []*interface{}{}}
}

func (e *Engine) Draw() error {
	w, h, err := terminal.GetSize()
	if err != nil {
		panic(err)
	}
	BackgroundFill(w, h)

	return nil
}
