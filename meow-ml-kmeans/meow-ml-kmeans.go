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