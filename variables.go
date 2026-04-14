package main

import "fmt"

func main() {
	const fName = "Cedric"
	const lName = "Robitaille"

	s := fmt.Sprintf("Hello %s %s\n", fName, lName)

	fmt.Print(s)
}
