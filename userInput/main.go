package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("What's your Name ?")
	var name string
	fmt.Scan(&name) // scan method not read after any whitespace is we input sandeep singh >> than out is {sandeep}
	fmt.Println("My Name is", name)

	fmt.Println()
	fmt.Println("Using bufio read after whitespace")
	reader := bufio.NewReader(os.Stdin)
	bio, _ := reader.ReadString('\n')
	fmt.Println("My FullNmae is", bio)
}
