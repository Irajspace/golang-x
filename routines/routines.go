package main

import (
	"fmt"
	"sync"
	"time"
)

// We pass a pointer to the WaitGroup so the function can tell it when it's done
func makeToast(wg *sync.WaitGroup) {
	defer wg.Done() // Subtracts 1 from the WaitGroup counter when the function exits
	fmt.Println("Started toasting bread...")
	time.Sleep(2 * time.Second)
	fmt.Println("Bread is done!")
}

func fryEggs(wg *sync.WaitGroup) {
	defer wg.Done() 
	fmt.Println("Started frying eggs...")
	time.Sleep(3 * time.Second)
	fmt.Println("Eggs are done!")
}

func main() {
	var wg sync.WaitGroup

	// We have 2 tasks to do, so we add 2 to the counter
	wg.Add(2)

	// Launch the goroutines and pass the WaitGroup to them
	go makeToast(&wg)
	go fryEggs(&wg)

	// Tell the main function to pause here until the counter hits 0
	wg.Wait()
	
	fmt.Println("Breakfast is ready!")
}