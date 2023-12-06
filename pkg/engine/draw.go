package engine

import (
	"github.com/goshathebusiness/termengi/pkg/models"
	"github.com/goshathebusiness/termengi/pkg/terminal"
)

type Engine struct {
	Items []interface{}
}

func NewEngine() *Engine {
	return &Engine{Items: []interface{}{}}
}

type Pixel struct {
	X, Y   int
	Symbol string
}

func (e *Engine) Draw() error {
	w, h, err := terminal.GetSize()
	if err != nil {
		panic(err)
	}
	BackgroundFill(w, h)
	var pixels []*Pixel
	for _, item := range e.Items {
		switch object := item.(type) {
		case *models.Rect:
			object.Lock()
			for i := object.X; i < (object.X + object.W); i++ {
				for j := object.Y; j < (object.Y + object.H); j++ {
					pixels = append(pixels, &Pixel{X: i, Y: j, Symbol: "#"})
				}
			}
			object.Unlock()
		}
	}

	err = drawPixels(pixels)
	if err != nil {
		return err
	}

	return nil
}

func drawPixels(pixels []*Pixel) error {
	for _, pixel := range pixels {
		ReplaceSymbolAtPosition(pixel.Y, pixel.X, pixel.Symbol)
	}

	return nil
}
