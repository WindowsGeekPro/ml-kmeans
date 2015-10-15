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

// blackhole -- center
func (my *meowCluster) blackhole() *meowPoint {
	return my.blackhole
}

// blackhole in deep -- re-centering
func(my *meowCluster) blackholeInDeep() float64 {
	totalDots := my.dots.meowLen()
	sourceDot := my.dots.meowFetch(0).(*meowPoint)
	totalCoordinates := sourceDot.stuffs.meowLen()

	// still more to code
}