package main

import "fmt"

func divOfTwoNum(a, b float64) float64 {
	if b == 0 {
		return 0
	}
	return a / b
}

func main() {
	fmt.Println("Start Error Handling Code")
	res := divOfTwoNum(10, 0)
	fmt.Println(res)
}
