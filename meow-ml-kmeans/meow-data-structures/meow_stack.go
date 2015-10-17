package meow-data-structures

import "sync"

type meowStackNode struct {
	meowData interface{}
	meowNxt *meowStackNode
}

type meowStack struct {
	meowHead *meowStackNode
	meowCount int
	meowLock *sync.Mutex
}

func meowNewStack() *meowStack {
	s := &meowStack{}
	s.meowLock = &sync.Mutex{}
	return s
}