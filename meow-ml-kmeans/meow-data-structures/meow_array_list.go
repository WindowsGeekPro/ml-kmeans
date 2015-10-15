package meow-data-structures

import (
	"fmt"
	"math/rand"
	"sync"
	"bytes"
	"time"
)

type meowArrayList struct {
	meowCount int
	meowLock *sync.Mutex
	stuffs []interface()
}

func meowNewArrayList() *meowArrayList {
	meowInstance := &meowArrayList {}
	meowInstance.meowLock = &sync.Mutex{}
	meowInstance.stuffs = make([]interface{}, 10)
	meowInstance.meowCount = 0
	rand.Seed(time.Now().UTC()UnixNano())
	return meowInstance
}

func (my *meowArrayList) meowLen() int {
	my.meowLock.Lock()
	defer my.meowLock.Unlock()
	return my.meowCount
}