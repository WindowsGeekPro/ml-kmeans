package meow-data-structures

import "sync"

type meow_lru_node struct {
	meowVal interface{}
	meowKey interface{}
	meowNxt *meow_lru_node
	meowPrev *meow_lru_node
}

type meow_lrucache struct {
	meowLock *sync.Mutex
	meowHead *meow_lru_node
	meowTail *meow_lru_node
	meowPower int
	stuffs map[interface{}]*meow_lru_node
}

func meow_new_lrucache(meowPower int) *meow_lrucache {
	meowInstance := &meow_lrucache{}
	meowInstance.meowLock = &sync.Mutex{}
	meowInstance.stuffs = make(map[interface{}]*meow_lru_node)
	meowInstance.meowPower = meowPower
	return meowInstance
}

func (my *meow_lrucache) meowLen() int {
	my.meowLock.Lock()
	defer my.meowLock.Unlock()
	return len(my.stuffs)
}