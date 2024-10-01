package main

import "fmt"

func main() {
	fmt.Println("Learn if else condition")
	a := 5
	b := 5
	if a > b {
		fmt.Println("A is greater")
	} else if a < b {
		fmt.Println("B is Greater")
	} else {
		fmt.Println("Both are equal")
	}

	x := 10
	y := 20
	if x < 20 && y > 15 {
		fmt.Println("Both are in Range x and y")
	}

	if x > 15 || y > 10 {
		fmt.Println("One consition is Passed")
	}

	if x > 8 && (x > 5 || y > 10) {
		fmt.Println("AND condition Passed with Bracket")
	}

	fmt.Println("Check Days")
	day := 7

	if day == 1 {
		fmt.Println("Monday : ", day)
	} else if day == 2 {
		fmt.Println("Tuesday : ", day)
	} else if day == 3 {
		fmt.Println("Wednesday : ", day)
	} else if day == 4 {
		fmt.Println("Thursday : ", day)
	} else if day == 5 {
		fmt.Println("Friday : ", day)
	} else if day == 6 {
		fmt.Println("Saturday : ", day)
	} else if day == 7 {
		fmt.Println("Sunday : ", day)
	} else {
		fmt.Println("UnKnown Day Pass", day)
	}

}
