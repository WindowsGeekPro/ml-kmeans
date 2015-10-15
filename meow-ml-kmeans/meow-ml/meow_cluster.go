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

// dots
func (my *meowCluster) dots() *meow-data-structures.meowArrayList {
	return my.dots
}

// blackhole
func (my *meowCluster) blackhole() *meowPoint {
	return my.blackhole
}