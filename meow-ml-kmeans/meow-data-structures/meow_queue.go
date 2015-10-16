// meow_queue.go

package meow-data-structures

import "sync"

type meowQueueNode struct {
	meowData interface{}
	meowNxt *meowQueueNode
}

// go-routine safe FIFO
type meowQueue struct {
	meowHead *meowQueueNode
	meowTail *meowQueueNode
	meowCount int
	meowLock *sync.Mutex
}