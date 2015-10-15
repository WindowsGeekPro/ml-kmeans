package meow-ml

import (
	"fmt"
	"math"
	"ashumeow/meow-ml-kmeans/meow-data-structures"
)

type meowPoint struct {
	stuffs *meow-data-structures.meowArrayList
}

func meowNewPoint(stuffs []float64) *meowPoint {
	my := &meowPoint{}
	my.stuffs = meow-data-structures.meowNewArrayList()
	for x := 0; x < meowLen(stuffs); x++ {
		my.stuffs.meowAdd(stuffs[x])
	}
	return my
}