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

// meowAddSliceDots
func (my *meowKmeans) meowAddSliceDots(stuffs []float64) {
	my.dots.meowAdd(meow-ml.meowNewKmeans(stuffs))
}

// clustering
func (my *meowKmeans) meow_cluster() *meow-data-structures.meowArrayList {
	if(my.num_of_clusters == 1) {
		panic("Pick more than one cluster ;)")
	}
	meow_clusters := meow-data-structures.meowArrayList()
	uniqueBlackholes := meow-data-structures.meowNewHashSet()
	for x := 0; x < my.num_of_clusters; x++ {
		randomBlackhole := my.dots.meowSample().(*meow-ml.meowPoint)
		for uniqueBlackholes.meowRegisters(randomBlackhole) {
			randomBlackhole = my.dots.meowSample().(*meow-ml.meowPoint)
		}
		uniqueBlackholes.meowAdd(randomBlackhole)
		meow_cluster := meow-ml.meowNewCluster(randomBlackhole)
		meow_clusters.meowAdd(meow_cluster)
	}

	// determining nearest meow_cluster for assigning a dot
	for x := 0; x < my.dots.meowLen(); x++ {
		meowSmallDist := math.MaxFloat64
		var meowNearCluster *meow-ml.meowCluster
	}
}