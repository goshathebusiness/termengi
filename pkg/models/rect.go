package models

import "sync"

type object struct {
	sync.Mutex
}

type Rect struct {
	object

	X, Y     int
	W, H     int
	Rotation int
}
