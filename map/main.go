package main

import (
	"fmt"
)

func main() {
	fmt.Println("Learn Map In  ?? Java we also use map")
	fmt.Println()
	//We have create a Map String as key and Value is Integer
	// so map is stored a  student data on String and Integer type

	student := make(map[string]int)

	student["sandeep"] = 80
	student["kuldeep"] = 95
	student["Risahbh"] = 99
	student["Lal"] = 80
	fmt.Println("Student Data iterate : ") // All Data Show
	fmt.Println(student)                   // All Data print for Student

	fmt.Println("sandeep Marks is : ", student["sandeep"]) // get a Data by key >> In go no concept of get() we have direct pass the key
	student["sandeep"] = 91                                // update the map and set new value
	fmt.Println("Update Sandeep Marks is : ", student["sandeep"])

	delete(student, "Lal") // Delete a Key in map and remove data into map
	student["Aman"] = 75   // new Data add in Map
	fmt.Println("Student Data :  ", student)
	// if we get a any one using key but this key is not present in map but they give me default value of Datatype
	//Check Key is present or Not
	grades, exixts := student["NahiHA"]
	fmt.Println("Marks of student :  ", grades)     // If present the value is set in value || if not presnt show defaut value of Datatype
	fmt.Println("Student is Exists of :  ", exixts) // return False if not exist in map || true if present

	//Iterate Map Using a For Loop Range

	for key, value := range student {
		fmt.Printf("Student {key} is :  %s <and> {Value} is :  %d\n", key, value)
	}

	//Insert data When we Initilized a MAP

	person := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
		"D": 4,
	}

	fmt.Println(person)
	for k, v := range person {
		fmt.Printf("Key is %s and value is %d\n", k, v)
	}

}
