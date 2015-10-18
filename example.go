package main

import (
	"fmt"
	"ashumeow/meow_ml_kmeans/meow_data_structures"
	"ashumeow/meow_ml_kmeans/meow_ml"
	"ashumeow/meow_ml_kmeans"
)

func main() {
	fmt.Println("---Queue---")
	meowQueue := meow_data_structures.meowNewQueue()
	meowQueue.meowPush("Testing... ")
	meowQueue.meowPush(4)
	meowQueue.meowPush(5)
	meowQueue.meowPush(8)
	meowQueue.meowPush(6)
	meowQueue.meowPush("Hold on!")
	fmt.Println("Total stuffs before distribution: ", meowQueue.meowLen())
	fmt.Println("Power before distribution: ", meowQueue.meowErect())
	fmt.Println(meowQueue.meowDistro(), 
				meowQueue.meowDistro(), 
				meowQueue.meowDistro(), 
				meowQueue.meowDistro(), 
				meowQueue.meowDistro(), 
				meowQueue.meowDistro(), 
				meowQueue.meowDistro()
				)
	fmt.Println("Total stuffs after distribution: ", meowQueue.meowLen())
	fmt.Println("Power: ", meowQueue.meowErect())

	fmt.Println("\n")

	fmt.Println("---Stack---")
	meowStack := meow_data_structures.meowNewStack()
	meowStack.meowPush("Hold on... ")
	meowStack.meowPush(4)
	meowStack.meowPush(5)
	meowStack.meowPush(8)
	meowStack.meowPush(6)
	meowStack.meowPush("Testing!")
	fmt.Println("Total stuffs before jump event: ", meowStack.meowLen())
	fmt.Println("Power before jump event: ", meowStack.meowErect())
	fmt.Println(meowStack.meowJump(), 
				meowStack.meowJump(), 
				meowStack.meowJump(), 
				meowStack.meowJump(), 
				meowStack.meowJump(), 
				meowStack.meowJump(), 
				meowStack.meowJump()
				)
	fmt.Println("Total stuffs after jump event: ", meowStack.meowLen())
	fmt.Println("Power: ", meowStack.meowErect())

	fmt.Println("\n")

	fmt.Println("---Array list---")
	meowArray := meow_data_structures.meowNewArrayList()
	fmt.Println("Listing out random stuff in an empty array: ", meowArray.meowSample())
	meowArray.meowAdd("Hello")
	meowArray.meowAdd("Ashu")
	meowArray.meowAdd("Meow")
	fmt.Println("Total stuffs in an array: ", meowArray.meowLen())
	fmt.Println("Array: ", meowArray)
	fmt.Println("Testing arrays now... \n")
	fmt.Println("Array[1]: ", meowArray.meowFetch(1))
	fmt.Println("Getting index of Hello: ", meowArray.IndexOf("Hello"))
	fmt.Println("Fetching registers: ", meowArray.meowRegisters("Hello"), meowArray.meowRegisters("random"))
	fmt.Println("Listing out random stuff in an array: ", meowArray.meowSample())
	fmt.Println("Listing out random stuff in an array: ", meowArray.meowSample())
	for x := 0; x < meowArray.meowLen(); x++ {
		fmt.Println("stuff(" , x, "):", meowArray.meowFetch(x))
	}
	fmt.Println("Array removal: ", meowArray.meowRemove("Ashu"), meowArray)
	fmt.Println("Array removal: ", meowArray.meowRemove("Meow"), meowArray)
	fmt.Println("Empty array: ", 
				meowArray.IsEmpty(), 
				meowArray.meowRemove("Hello"),
				meowArray.IsEmpty(),
				meowArray.meowLen()
				)
	for x := 0; x < 30; x++ {
		meowArray.meowAdd(x)
	}
	fmt.Println("Array indexing: ", meowArray.IndexOf(5), meowArray.meowLen())
	meowArray.meowReset()
	meowArray.meowAdd("un", "deux", "trois")
	meowArray_new2 := meow_data_structures.meowNewArrayList()
	meowArray_new2.meowAdd("quatre", "cinq", "six")
	meowArray.meowAddFromList(meowArray_new2)
	fmt.Println("Array: ", meowArray, meowArray.meowLen())
	fmt.Println("Source-Destination arrays: ", meowArray.meowSource(), meowArray.meowDest())
	fmt.Println("Sliced array: ", meowArray.meowSlice())

	fmt.Println("\n")

	fmt.Println("---Hash Set---")
	meowSet := meow_data_structures.meowNewHashSet()
	meowSet.meowAdd("hello", "meow", "Hello")
	fmt.Println("List out total stuffs in the set: ", meowSet.meowLen(), meowSet)
	meowSlicingSet := meowSet.meowSlice()
	fmt.Println("Sliced set: ", meowSlicingSet, len(meowSlicingSet))

	fmt.Println("\n")

	fmt.Println("---Clustering---")
	meowCluster := meow_ml_kmeans.meowNewCluster(4)
	meowCluster.meowAddSliceDots([]float64{48.2641334571,86.4516903905})
	meowCluster.meowAddSliceDots([]float64{0.114004262656,35.8368597414})
	meowCluster.meowAddSliceDots([]float64{97.4319168245,92.8009240744})
	meowCluster.meowAddSliceDots([]float64{24.4614031388,18.3292584382})
	meowCluster.meowAddSliceDots([]float64{36.2367675367,32.8294024271})
	meowCluster.meowAddSliceDots([]float64{75.5836860736,68.30729977})
	meowCluster.meowAddSliceDots([]float64{38.6577034445,25.7701728584})
	meowCluster.meowAddSliceDots([]float64{28.2607136287,64.4493377817})
	meowCluster.meowAddSliceDots([]float64{61.5358486771,61.2195232194})
	meowCluster.meowAddSliceDots([]float64{1.52352224798,38.5083779618})
	meowCluster.meowAddSliceDots([]float64{11.6392182793,68.2369021579})
	meowCluster.meowAddSliceDots([]float64{53.9486870607,53.9136556533})
	meowCluster.meowAddSliceDots([]float64{14.6671651772,26.0132534731})
	meowCluster.meowAddSliceDots([]float64{65.9506725878,82.5639317581})
	meowCluster.meowAddSliceDots([]float64{58.3682872339,51.6414580337})
	meowCluster.meowAddSliceDots([]float64{12.6918921252,2.28888447759})
	meowCluster.meowAddSliceDots([]float64{31.7587852231,18.1368234166})
	meowCluster.meowAddSliceDots([]float64{63.6631115204,24.933301389})
	meowCluster.meowAddSliceDots([]float64{29.1652289905,34.456759171})
	meowCluster.meowAddSliceDots([]float64{44.3830953085,70.4813875779})
	meowCluster.meowAddSliceDots([]float64{47.0571691145,65.3507625811})
	meowCluster.meowAddSliceDots([]float64{74.0584537502,98.2271944247})
	meowCluster.meowAddSliceDots([]float64{55.8929146157,86.6196265477})
	meowCluster.meowAddSliceDots([]float64{20.4744253473,12.0025149302})
	meowCluster.meowAddSliceDots([]float64{14.2867767281,40.2850440995})
	meowCluster.meowAddSliceDots([]float64{40.43551369,94.5410407116})
	meowCluster.meowAddSliceDots([]float64{87.6178871195,12.4700151639})
	meowCluster.meowAddSliceDots([]float64{47.2703048197,93.0636237124})
	meowCluster.meowAddSliceDots([]float64{59.7895104175,69.2621288413})
	meowCluster.meowAddSliceDots([]float64{80.8612333922,42.9183411179})
	meowCluster.meowAddSliceDots([]float64{31.1271795535,55.6669044656})
	meowCluster.meowAddSliceDots([]float64{78.9671049353,65.833739365})
	meowCluster.meowAddSliceDots([]float64{39.8324533414,63.0343115139})
	meowCluster.meowAddSliceDots([]float64{79.126343548,14.9128874133})
	meowCluster.meowAddSliceDots([]float64{65.8152400306,77.5202358013})
	meowCluster.meowAddSliceDots([]float64{75.2762752704,42.4858435609})
	meowCluster.meowAddSliceDots([]float64{29.6475948493,61.2068411763})
	meowCluster.meowAddSliceDots([]float64{67.421857106,54.8955604259})
	meowCluster.meowAddSliceDots([]float64{10.4652931501,29.7954139372})
	meowCluster.meowAddSliceDots([]float64{32.0272462745,99.5422900971})
	meowCluster.meowAddSliceDots([]float64{80.1520927001,84.2710379142})
	meowCluster.meowAddSliceDots([]float64{2.27240208403,41.2138854089})
	meowCluster.meowAddSliceDots([]float64{44.4601509555,1.72563901513})
	meowCluster.meowAddSliceDots([]float64{16.8676021068,35.3415636277})
	meowCluster.meowAddSliceDots([]float64{58.1977544121,29.2752085455})
	meowCluster.meowAddSliceDots([]float64{24.6119080085,39.9440735137})
	meowCluster.meowAddSliceDots([]float64{63.0759798755,60.9841014448})
	meowCluster.meowAddSliceDots([]float64{30.9289119657,95.0173219502})
	meowCluster.meowAddSliceDots([]float64{8.54972950047,41.7384441737})
	meowCluster.meowAddSliceDots([]float64{61.2606910793,4.06738902059})
	meowCluster.meowAddSliceDots([]float64{83.2302091964,11.6373312879})
	meowCluster.meowAddSliceDots([]float64{89.4443065362,42.5694882801})
	meowCluster.meowAddSliceDots([]float64{24.5619318152,97.7947977804})
	meowCluster.meowAddSliceDots([]float64{50.3134024475,40.6429336223})
	meowCluster.meowAddSliceDots([]float64{58.1422402033,36.1112632557})
	meowCluster.meowAddSliceDots([]float64{32.0668520827,29.9924151435})
	meowCluster.meowAddSliceDots([]float64{89.6057447137,84.9532177777})
	meowCluster.meowAddSliceDots([]float64{9.8876440816,18.2540486261})
	meowCluster.meowAddSliceDots([]float64{17.9670383961,47.596032257})
	meowCluster.meowAddSliceDots([]float64{50.2977668282,93.6851189223})
	meowCluster.meowAddSliceDots([]float64{98.0700386253,86.5816924579})
	meowCluster.meowAddSliceDots([]float64{10.8175290981,26.4344732252})
	meowCluster.meowAddSliceDots([]float64{34.7463851288,24.4154447141})
	meowCluster.meowAddSliceDots([]float64{92.5470100593,17.3595513748})
	meowCluster.meowAddSliceDots([]float64{79.0426629356,4.59850018907})
	meowCluster.meowAddSliceDots([]float64{89.9791366918,29.523946842})
	meowCluster.meowAddSliceDots([]float64{3.89920214563,91.3650215111})
	meowCluster.meowAddSliceDots([]float64{35.4669861576,62.1865368798})
	meowCluster.meowAddSliceDots([]float64{2.78150918086,24.5280230552})
	meowCluster.meowAddSliceDots([]float64{50.0390951889,57.0414421682})
	meowCluster.meowAddSliceDots([]float64{64.4521660758,48.4962172448})
	meowCluster.meowAddSliceDots([]float64{94.4915452316,56.6508179406})
	meowCluster.meowAddSliceDots([]float64{47.1655534769,15.8292055671})
	meowCluster.meowAddSliceDots([]float64{94.2027011374,45.6802385454})
	meowCluster.meowAddSliceDots([]float64{30.5846324871,54.783635876})
	meowCluster.meowAddSliceDots([]float64{57.7043252948,0.286661610381})
	meowCluster.meowAddSliceDots([]float64{41.7908674949,14.7206014023})
	meowCluster.meowAddSliceDots([]float64{59.6689465934,64.8849831965})
	meowCluster.meowAddSliceDots([]float64{92.2553335495,55.9096460272})
	meowCluster.meowAddSliceDots([]float64{48.493467262,69.4766837809})
	meowCluster.meowAddSliceDots([]float64{23.1837859581,71.4406867443})
	meowCluster.meowAddSliceDots([]float64{29.0737623652,66.9391416961})
	meowCluster.meowAddSliceDots([]float64{95.7442323112,89.4677505059})
	meowCluster.meowAddSliceDots([]float64{68.7707275828,40.9900140055})
	meowCluster.meowAddSliceDots([]float64{84.5445737133,32.1707309618})
	meowCluster.meowAddSliceDots([]float64{67.4126251988,56.6710579117})
	meowCluster.meowAddSliceDots([]float64{10.688352016,28.1745892928})
	meowCluster.meowAddSliceDots([]float64{56.7620324155,18.3034334207})
	meowCluster.meowAddSliceDots([]float64{50.6751320678,86.6916908032})
	meowCluster.meowAddSliceDots([]float64{74.6185482896,34.022483532})
	meowCluster.meowAddSliceDots([]float64{20.7011996002,32.855295357})
	meowCluster.meowAddSliceDots([]float64{11.479054664,1.59204297586})
	meowCluster.meowAddSliceDots([]float64{51.6805387648,25.4063026358})
	meowCluster.meowAddSliceDots([]float64{84.4109522357,47.237632645})
	meowCluster.meowAddSliceDots([]float64{90.6395051745,57.7917166935})
	meowCluster.meowAddSliceDots([]float64{58.6159601042,84.1226173848})
	meowCluster.meowAddSliceDots([]float64{46.2184509277,28.559934585})
	meowCluster.meowAddSliceDots([]float64{97.0302485783,41.3135022812})
	meowCluster.meowAddSliceDots([]float64{31.3144587058,87.2459910122})
	meowCluster.meowAddSliceDots([]float64{5.93357833962,95.6812831872})
	meow_clusters := meowCluster.meow_cluster()

	for x := 0; x < meow_clusters.meowLen(); x++ {
		meow_cluster := meow_clusters.meowFetch(x).(*meow_ml.meowCluster)
		fmt.Println("Clustering: ", meow_cluster.blackhole().stuffs().meowSlice())
		for xx := 0; xx < meow_cluster.dots().meowLen(); xx++ {
			dot := meow_cluster.dots().meowFetch(xx).(*meow_ml.meowPoint)
			fmt.Println("---", dot.stuffs().meowSlice())
		}
	}
}