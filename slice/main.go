package main

import (
	"fmt"
)

func main() {
	fmt.Println("Learn Slice This is Same as List there are no Fixed Size")
	number := []int{1, 2, 3, 4, 5, 6} // this is same as list no size define
	fmt.Println(number)
	number = append(number, 7, 8, 9, 10, 22) // add more number
	fmt.Println(number)
	num := number[len(number)-1] // get last number
	fmt.Println("last number in Slice : ", num)
	fmt.Println("length of Slice : ", len(number))

	fmt.Println("Create a Slice empty")
	name := []string{}
	fmt.Println("Empty Splice ", name)
	fmt.Println("Make a Slices USing a Make()")
	prices := make([]int, 3, 6) // Slice create Using make()function weh have define a length and capaity of slice
	fmt.Println("Prices Slice", prices)
	fmt.Println("Length of prices is ", len(prices))
	fmt.Println("Capacity of Slice is", cap(prices))
	prices = append(prices, 4)
	prices = append(prices, 5, 6, 7)
	fmt.Println("Prices Slice after append", prices)
	fmt.Println("Length prices after ", len(prices))
	fmt.Println("Capacity Slice after", cap(prices))

}
