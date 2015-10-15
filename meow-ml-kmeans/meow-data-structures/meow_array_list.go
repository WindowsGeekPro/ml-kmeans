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
	defer meowLock.Unlock()

	for o := range objects {
		my.meowAdd(o)
	}
}
func (my *meowArrayList) meow_add(o interface{}) {
	my.stuffs[my.meowCount] = o
	my.meowCount++
	my.resize()
}