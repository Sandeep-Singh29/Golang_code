package main

import (
	"fmt"
	// "mylearning/myutils"
)

func main() {
	fmt.Println("Learn Go Language")
	// myutils.PrintMessage("This is my test")
	var firstName string = "Sandeep" // nit mandatory to define a dataType but if we give datatype than compiler check in this refernce only string is passed
	var middleName = "Singh"         // In this line we have not define any datatype so runtime compiler check which datatype you have passed in variable
	var lastName string = "Kushwaha"
	fmt.Println("First Name :  ", firstName)
	fmt.Println("Middle Name :  ", middleName)
	fmt.Println("Last Name :  " + lastName) // You can Also use for concate + or , for concate a string

	var age int = 27 // we have not define any datatype same as string this will work as a same compiler checked in runtime
	fmt.Println("My age is : ", age)
	var salary = 100029
	fmt.Println("Salary is : ", salary)
	var userIsActive bool = true
	fmt.Println("User isActive ?  : ", userIsActive)
	fmt.Println("--------------------------")
	const fatherName string = "Brijendra Singh" // if we declare a const(Constant in java final) we have not update it its will fix
	fmt.Println("FatherName is Constant : ", fatherName)
	var address = "Delhi" // In address we have assign a address
	fmt.Println("Before without Change Address : ", address)
	address = "Banglore" // after the address we have update a address bcz we have not used a const in this placed
	fmt.Println("After Change Address : ", address)

	fmt.Println("---------Other Way to declare Variable---------")
	godName := "Mahadev" //Other way to define
	fmt.Println(godName)
}
