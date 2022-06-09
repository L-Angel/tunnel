package worker

import "time"

type Position struct {
	Name       string
	Pos        uint64
	Force      bool
	UpdateTime time.Time
}

func NewPosition(name string, pos uint64, force bool) *Position {
	return &Position{Name: name, Pos: pos, UpdateTime: time.Now()}
}

func NewPositionWithForce(name string, pos uint64) *Position {
	return &Position{Name: name, Pos: pos, Force: true, UpdateTime: time.Now()}
}
