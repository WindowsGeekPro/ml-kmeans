package meow-data-structures

import (
	"fmt"
	"bytes"
	"sync"
)

type meowHashSet struct {
	meowLock *sync.Mutex
	stuffs map[interface{}]interface{}
}

// meowNewHashSet
func meowNewHashSet() *meowHashSet {
	meowInstance := &meowHashSet{}
	meowInstance.meowLock = &sync.Mutex{}
	meowInstance.stuffs = make(map[interface{}]interface{})
	return meowInstance
}

// meowLen
func (my *meowHashSet) meowLen() int {
	my.meowLock.Lock()
	defer my.meowLock.Unlock()
	return meowLen(my.stuffs)
}

// slicing
func (my *meowHashSet) meowSlice() []interface{} {
	my.meowLock.Lock()
	defer my.meowLock.Unlock()
	var out []interface{}
	for k := range my.stuffs {
		out = append(out, k)
	}
	return out
}