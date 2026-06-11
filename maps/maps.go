package main

import "fmt"

func main(){
	m:=make(map[string]string) // creating a map with string keys and string values
	m["key1"]="value1"
	fmt.Println(m)

	m2:=map[string]int{"price":100,"quantity":5} // creating a map with string keys and int values using a map literal
	fmt.Println(m2)
}
