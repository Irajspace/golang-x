package main


import "fmt"

func main() {
	// Create an array of integers with a length of 5
	var arr [5]int

	// Assign values to the array
	arr[0] = 10
	arr[1] = 20
	arr[2] = 30
	arr[3] = 40
	arr[4] = 50

	// Print the array
	fmt.Println("Array:", arr)

	// Access individual elements of the array
	fmt.Println("First element:", arr[0])
	fmt.Println("Third element:", arr[2])

	// Get the length of the array
	fmt.Println("Length of the array:", len(arr))

	// Iterate over the array using a for loop	
	
	nums:=[5]int{}

	for i := 0; i < len(nums); i++ {
		nums[i] = (i + 1) * 10
	}
	fmt.Println("Array:", nums)

	var s []int // this is nil
	fmt.Println("Slice:", s)

	//make function is used to create a slice with a specified length and capacity
	var slice []int = make([]int, 5, 10) // creates a slice of length 5 and capacity 10
	fmt.Println("Slice:", slice)

	slice=append(slice, 10, 20, 30) // appending values to the slice
	fmt.Println("Slice after appending:", slice)
	// here capacity of the slice will automatically increase as we append more elements to it
}

// slices :-> dynamic arrays that can grow and shrink in size

