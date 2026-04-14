package main

import (
	"fmt"
	"unicode/utf8"
)

func stringInterpolation() {
	/*
		For string interpolation:
		- %v any value in default format ("i am %v years old", 23 OR "too many")
		- %s string values ("i am %s years old", "too many")
		- %d integer values ("i am %d years old", 23)
		- %f float values ("i am %.2f years old", 23.756)
			- %.2f denotes decimal points.
	*/

	const name = "Cedric"
	const openRate = 5.2

	msg := fmt.Sprintf("Hi %s, your open rate is %.1f percent\n", name, openRate)
	fmt.Print(msg)
}

func stringEncoding() {
	/*
		In Go, strings are a sequence of bytes that can hold any data.
		Using the 'rune' type, it breaks the string into individual characters (where each can be more than 1 byte long)
		Therefore, we can include various Unicode characters / chinese characters / whathaveyou.
	*/

	const name = "🚀"
	fmt.Printf("constant 'name' byte length: %d\n", len(name))
	fmt.Printf("constant 'name' rune length: %d\n", utf8.RuneCountInString(name))
	fmt.Println("=========")
	fmt.Printf("Hello %s\n", name)
}

func formatPractice() {
	fname := "Dalinar"
	lname := "Kholin"
	age := 45
	messageRate := 0.5
	isSubscribed := false
	message := "Sometimes a hypocrite is nothing more than a man in the process of changing."

	userLog := fmt.Sprintf("Name: %s %s, Age: %d, Rate: %.1f, Is Subscribed: %t, Message: %s\n", fname, lname, age, messageRate, isSubscribed, message)

	fmt.Printf(userLog)
}

func main() {
	// stringInterpolation()
	// stringEncoding()
	formatPractice()
}
