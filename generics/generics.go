package main
import "fmt"
func printSlice[T int|string](items []T){
	for _, item := range items {
		fmt.Println(item)
	}
}

func main(){
	intSlice := []int{1, 2, 3, 4, 5}
	stringSlice := []string{"Hello", "Generics", "in", "Go"}

	fmt.Println("Integer Slice:")
	printSlice(intSlice)

	fmt.Println("\nString Slice:")
	printSlice(stringSlice)
}
