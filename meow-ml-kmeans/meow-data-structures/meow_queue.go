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

// creating a new pointer to a new meowQueue
func meowNewQueue() *meowQueue {
	q := &meowQueue{}
	q.meowLock = &sync.Mutex{}
	return q
}