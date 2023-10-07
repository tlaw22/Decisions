package main

import "log"

func main() {

	var isTrue bool
	isTrue = false
	// Test if isTrue is true or false
	if isTrue {
		log.Println("It is true", isTrue)
	} else {
		log.Println("It is false", isTrue)
	}

	cat := "dog"
	if cat == "cat" {
		log.Println("What pet is in the pet shot? ", cat)
	} else {
		log.Println("It is not a cat! ", cat)
	}
}
