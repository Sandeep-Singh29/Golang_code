package main

import "fmt"

func divOfTwoNum(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("denominator zero nahi hona chaiya")
	}
	return a / b, nil
}

func main() {
	fmt.Println("Start Error Handling Code")

	// res, err := divOfTwoNum(4, 2)
	// if err != nil {
	// 	fmt.Println("Error Handling")
	// }
	// fmt.Println("Division of two Number is :", res)

	res, _ := divOfTwoNum(10, 2) // under scope is used to avoid handing error || we have handle the error in previous code
	fmt.Println("Division of two Number is :", res)

}
