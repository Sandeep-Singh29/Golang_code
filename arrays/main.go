package main

import "fmt"

func main() {
	fmt.Println("Learn Array in Golang")

	var name [5]string                        // way 1 to assign the array
	name[0] = "sandeep"                       // pass the value in index 0
	name[1] = "kuldeep"                       // pass the value in index 1
	name[2] = "rishabh"                       // pass the value in index 2
	name[3] = "aman"                          // pass the value in index 3
	name[4] = "vivek"                         // pass the value in index 4
	fmt.Println("NAme of Person is : ", name) // Print the array

	var number = [5]int{1, 2, 3, 4, 5} // 2nd way to insilized a array we have assign element in one time
	fmt.Println("Number is : ", number)
	fmt.Println("Length of name Array is : ", len(name)) // find the length of array using len() method
	fmt.Println("Name in 2nd Index is : ", name[1])      // Element Acces using index

	var price [5]string                     // Define Array for 5 size for int type
	fmt.Println("Price of Array : ", price) // if you print initial all index is default 0 value
	fmt.Printf("price array is q means(quotes)  %q\n", price)

	arr := [5]int{1, 2, 3, 4, 5} // way 3 to initilized a Array
	fmt.Println("arr print ", arr)
	fmt.Println("Array with Automatic Length")
	arrAuto := [...]int{1, 2, 3, 4, 5, 6, 7, 8} // Array with Automatic Length no need to define a size of array
	fmt.Println(arrAuto)
	fmt.Println("Length of arrAuto is : ", len(arrAuto))

}
