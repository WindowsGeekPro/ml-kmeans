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

// if empty
func (my *meowHashSet) meowEmpty() bool {
	return my.meowLen() == 0
}

// meowAdd
func (my *meowHashSet) meowAdd(objects ...interface{}) {
	my.meowLock.Lock()
	defer my.meowLock.Unlock()
	for o := range objects {
		my.stuffs[o] = true
	}
}

// meowFetch
func (my *meowHashSet) meowFetch(k interface{}) interface{} {
	my.meowLock.Lock()
	defer my.meowLock.Unlock()
	return my.stuffs[k]
}