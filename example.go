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

	// still more to code...
}