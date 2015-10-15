package meow-ml

import "ashumeow/ml-kmeans/meow-data-structures"

type meowCluster struct {
	blackhole *meowPoint
	dots *meow-data-structures.meowArrayList
}

// new meowCluster
func meowNewCluster(blackhole *meowPoint) *meowCluster {
	my := &meowCluster{}
	my.blackhole = blackhole
	my.dots = meow-data-structures.meowNewArrayList()
	return my
}