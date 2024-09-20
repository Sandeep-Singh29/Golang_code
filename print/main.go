package main

import (
	"fmt"
)

func main() {
	age := 26
	name := "sandeep"
	height := 5.113213
	//	fmt.Println("age:", age, "name:", name, "height:", height)
	//	fmt.Println("Hello World..!!") //Println same as Java mean print and new line

	fmt.Println("-------Use Formatter Method-------")
	fmt.Printf("name is %s\n", name)
	fmt.Printf("My age is %d\n", age)
	fmt.Printf("My Height is %.2f\n", height)
	fmt.Printf("Find the datatype of height is ::  %T\n", height)
	fmt.Printf("Find the datatype of name is ::  %T\n", name)
	fmt.Printf("Find the datatype of age is ::  %T\n", age)
}
