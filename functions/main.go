package main

import "fmt"

func myFirstPrintFun() {
	fmt.Println("This is My First Function Without parameter and without return")
}

func addTwoNumber(a, b int) {
	fmt.Println("Add to Number", a, "and", b, "= ", a+b)
}

func multipleOfTwoNumber(a, b int) int { // if same type parameter than declare datatype in last
	print("Multilpe of ", a, " and ", b, " = ")
	return a * b
}

func concatString(str string) string { // return type write in after the parameter bracker (a,b int) (int) >> this last is return type
	return str + "concatString"
}

func divideTwoNumber(a int, b int) float32 { // we have declare a (a int ,b int) if same type variable u can also write (a,b int)
	fmt.Println("Divison of ", a, " and ", b, " = ") // if both variable us differnt type(a string , b int)
	return float32(a) / float32(b)
}

func sumOfTwoNumbers(a, b int) (result int) { // if you want to return a variable to store a ans than we have first declare it ans return
	fmt.Print("Sum Of ", a, " - ", b, " =  ")
	result = a - b
	return
}

func main() {
	fmt.Println("---------------Learing a Function or Method using Go---------------")
	myFirstPrintFun()
	fmt.Println("Parameter Method Add Two Number")
	addTwoNumber(4, 4)
	fmt.Println("Return 2 Number Multiple")
	multiple := multipleOfTwoNumber(2, 2)
	fmt.Println(multiple)
	fmt.Println()
	concatstr := concatString("sandeep")
	fmt.Println(concatstr)
	result := divideTwoNumber(10, 5)
	fmt.Println("Division result:", result)
	fmt.Println()
	sumAns := sumOfTwoNumbers(10, 5)
	fmt.Println(sumAns)

}
