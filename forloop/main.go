package main

import "fmt"

func main() {
	fmt.Println("Learn Golang For Loops")

	for i := 1; i <= 2; i++ {
		fmt.Println("Number is : ", i)
	}

	fmt.Println("Infinite Loop")

	count := 1

	for { // infinte loop we have not pass any condition ut we wirte case when the case is match the loop is terminate worked as while
		fmt.Println("Number of Count is : ", count)
		count++
		if count == 3 {
			break
		}
	}

	fmt.Println("Learn Range in For Loop")
	number := []int{30, 12, 65, 8, 43, 1}
	fmt.Println(number)
	fmt.Println("Size of Number Slice is ", len(number))
	for i, v := range number {
		fmt.Println("Index is : ", i, "  &  Value is : ", v)
	}

	name := "Sandeep Singh"

	for i, v := range name { // Using Range we have get the index and the index value
		fmt.Printf("Index of Data is %d, and Value is %c\n", i, v)
	}

}
