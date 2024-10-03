package main

import "fmt"

func main() {
	fmt.Println("Learn Switch Case")

	day := 10
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Please Insert a Valid day")
	}

	months := "July"

	switch months {
	case "January", "Febuary", "March":
		fmt.Println("Winter Time")
	case "April", "May", "June":
		fmt.Println("Summer Session")
	case "July", "August":
		fmt.Println("Rainy")
	case "September", "October", "November", "December":
		fmt.Println("Normal Session")
	default:
		fmt.Println("Unknown Month")
	}

	temperature := 30
	switch {
	case temperature < 0:
		fmt.Println("Freezing")
	case temperature > 5 && temperature < 15:
		fmt.Println("Winter")
	case temperature > 20 && temperature < 30:
		fmt.Println("Warm")
	case temperature > 30 && temperature < 45:
		fmt.Println("Hot")
	default:
		fmt.Println("NONE")
	}

}
