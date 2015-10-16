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

func meowNewHashSet() *meowHashSet {
	meowInstance := &meowHashSet{}
	meowInstance.meowLock = &sync.Mutex{}
	meowInstance.stuffs = make(map[interface{}]interface{})
	return meowInstance
}