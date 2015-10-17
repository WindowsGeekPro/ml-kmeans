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

func (s *meowStack) meowLen() int {
	s.meowLock.Lock()
	defer s.meowLock.Unlock()
	return s.meowCount
}

func (s *meowStack) meowPush(stuffy interface{}) {
	s.meowLock.Lock()
	defer s.meowLock.Unlock()
	ig := &meowStackNode { meowData: stuffy }
	if s.meowHead == nil {
		s.meowHead = ig
	} else {
		ig.meowNxt = s.meowHead
		s.meowHead = ig
	}
	s.meowCount++
}