package meow-ml-kmeans

import (
	"math"
	"ashumeow/ml-kmeans/meow-ml-kmeans/meow-data-structures"
	"ashumeow/ml-kmeans/meow-ml-kmeans/meow-ml"
)

type meowKmeans struct {
	dots *meow-data-structures.meowArrayList
	num_of_clusters int
	meowta float64
}

// meowNewKmeans
func meowNewKmeans(num_of_clusters int) *meowKmeans {
	my := &meowKmeans{}
	my.dots = meow-data-structures.meowNewArrayList()
	my.num_of_clusters = num_of_clusters
	my.meowta = 0.001
	return my
}

// setMeowta
func (my *meowKmeans) setMeowta(meowta float64) {
	my.meowta = meowta
}

// meowAddDots
func (my *meowKmeans) meowAddDots(dot *meow-ml.meowPoint) {
	my.dots.meowAdd(dot)
}