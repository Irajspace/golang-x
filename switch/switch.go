package main

import "fmt"

func main(){
	whoAmI :=func(i interface{}){
		switch t:= i.(type){
		case int:
			fmt.Printf("I am an integer and my value is %d\n",t)
		case string:
			fmt.Printf("I am a string and my value is %s\n",t)
		default:
			fmt.Printf("I am of type %T and my value is %v\n",t,t)
		}
	}
	whoAmI(42)
	whoAmI("Hello, World!")
	whoAmI(3.14)
}
