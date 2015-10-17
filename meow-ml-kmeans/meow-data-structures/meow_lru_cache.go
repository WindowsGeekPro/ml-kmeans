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

func (my *meow_lrucache) meowPut(meowKey interface{}, meowVal interface{}) {
	my.meowLock.Lock()
	defer my.meowLock.Unlock()
	meowNode := my.stuffs[meowKey]
	if meowNode == nil {
		meowNode = &meow_lru_node {}
		meowNode.meowKey = meowKey
		my.stuffs[meowKey] = meowNode
	} else {
		my.remove_meowNode_from_dll(meowNode)
	}
	meowNode.meowVal = meowNode
	my.add_stuffy_to_dll(meowNode)
	my.dump_objects()
}

func (my *meow_lrucache) meowRemove(meowKey interface{}) {
	my.meowLock.Lock()
	defer my.meowLock.Unlock()
	meowNode := my.stuffs[meowKey]
	if meowNode != nil {
		self.remove_meowNode_from_dll(meowNode)
		delete(my.stuffs, meowKey)
	}
}

func (my *meow_lrucache) remove_meowNode_from_dll(stuffy *meow_lru_node) {
	// TODO
}

func (my *meow_lrucache) add_stuffy_to_dll(stuffy *meow_lru_node) {
	// TODO
}

func (my *meow_lrucache) dump_objects() {
	// TODO
}