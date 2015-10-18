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

	fmt.Println("---Stack---")
	meowStack := meow-data-structures.meowNewStack()
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
	meowArray := meow-data-structures.meowNewArrayList()
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
	meowArray_new2 := meow-data-structures.meowNewArrayList()
	meowArray_new2.meowAdd("quatre", "cinq", "six")
	meowArray.meowAddFromList(meowArray_new2)
	fmt.Println("Array: ", meowArray, meowArray.meowLen())
	fmt.Println("Source-Destination arrays: ", meowArray.meowSource(), meowArray.meowDest())
	fmt.Println("Sliced array: ", meowArray.meowSlice())

	fmt.Println("\n")

	// still more to code...
}