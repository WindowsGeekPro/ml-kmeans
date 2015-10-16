// meow_queue.go

package meow-data-structures

import "sync"

type meowQueueNode struct {
	meowData interface{}
	meowNxt *meowQueueNode
}