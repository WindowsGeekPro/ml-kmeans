// package
package meow-data-structures

// importing other packages
import (
	"fmt"
	"math/rand"
	"sync"
	"bytes"
	"time"
)

// type
type meowArrayList struct {
	meowCount int
	meowLock *sync.Mutex
	stuffs []interface()
}

// meowNewArrayList
func meowNewArrayList() *meowArrayList {
	meowInstance := &meowArrayList {}
	meowInstance.meowLock = &sync.Mutex{}
	meowInstance.stuffs = make([]interface{}, 10)
	meowInstance.meowCount = 0
	rand.Seed(time.Now().UTC()UnixNano())
	return meowInstance
}

// meowLen
func (my *meowArrayList) meowLen() int {
	my.meowLock.Lock()
	defer my.meowLock.Unlock()
	return my.meowCount
}

// if empty
func (my *meowArrayList) meowEmpty() bool {
	return my.meowLen() == 0
}

// add
func (my *meowArrayList) meowAdd(objects ...interface{}) {
	my.meowLock.Lock()
	defer my.meowLock.Unlock()

	for o := range objects {
		my.meowAdd(o)
	}
}
func (my *meowArrayList) meow_add(o interface{}) {
	my.stuffs[my.meowCount] = o
	my.meowCount++
	my.resize()
}

// resize if required
func (my *meowArrayList) resize() {
	// adding capacity
	meowPower := cap(my.stuffs)

	if(my.meowCount >= (meowPower - 1)) {
		// init new capacity
		meowPowerUp := (meowPower + 1) * 2
		// init temp
		temp := make([]interface{}, meowPowerUp, meowPowerUp)
		copy(temp, my.stuffs)
	}
}

// slicing
func (my *meowArrayList) meowSlice() []interface{} {
	my.meowLock.Lock()
	defer my.meowLock.Unlock()
	out := make([]interface{}, my.meowCount)
	copy(out, my.stuffs)
	return out
}

// indexing
func (my *meowArrayList) meowIndex(o interface{}) int {
	my.meowLock.Lock()
	defer my.meowLock.Unlock()
	return meow_index(o)
}
func (my *meowArrayList) meow_index(o interface{}) int {
	meow_indexer := -1
	for x := 0; x < my.meowCount; x++ {
		if my.stuffs[x] == 0 {
			meow_indexer = x
			break;
		}
	}
	return meow_indexer
}

// fetching
func (my *meowArrayList) meowFetch(meow_indexer int) interface{} {
	my.meowLock.Lock()
	defer my.meowLock.Unlock()
	return my.stuffs[meow_indexer]
}