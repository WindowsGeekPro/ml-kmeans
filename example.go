package main

import (
	"fmt"
	"ashumeow/meow-ml-kmeans/meow-data-structures"
	"ashumeow/meow-ml-kmeans/meow-ml"
	"ashumeow/meow-ml-kmeans"
)

func main() {
	fmt.Println("---Queue---")
	meowQueue := meow-data-structures.meowNewQueue()
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

	// still more to code...
}