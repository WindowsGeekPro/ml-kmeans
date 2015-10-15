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

func meowNewKmeans(num_of_clusters int) *meowKmeans {
	my := &meowKmeans{}
	my.dots = meow-data-structures.meowNewArrayList()
	my.num_of_clusters = num_of_clusters
	my.meowta = 0.001
	return my
}