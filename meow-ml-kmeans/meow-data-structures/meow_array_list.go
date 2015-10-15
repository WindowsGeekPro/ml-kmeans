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