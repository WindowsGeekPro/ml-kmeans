package meow-data-structures

import "sync"

type meowStackNode struct {
	meowData interface{}
	meowNxt *meowStackNode
}