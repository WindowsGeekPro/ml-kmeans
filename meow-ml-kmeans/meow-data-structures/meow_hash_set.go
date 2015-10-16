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