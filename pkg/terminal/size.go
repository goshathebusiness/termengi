package terminal

import "golang.org/x/term"

const (
	defaultWidth  = 20
	defaultHeight = 20
)

func GetSize() (int, int, error) {
	w, h, err := term.GetSize(0)
	if err != nil {
		return defaultWidth, defaultHeight, err
	}

	return w, h, nil
}
