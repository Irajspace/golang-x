package main

import "fmt"

func add(nums... int)int{
	sum:=0
	for _,num:=range nums{
		sum+=num
	}
	return sum
}

func main(){


	num:=add(1,2,3,4,5)
	fmt.Println(num)
}