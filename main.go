package main

import "log"

func main() {

	// declare some test variables
	var isTrue bool
	isTrue = false
	// Test if isTrue is true or false
	if isTrue {
		log.Println("It is true", isTrue)
	} else {
		log.Println("It is false", isTrue)
	}
	// Make a decision with strings
	cat := "dog"
	if cat == "cat" {
		log.Println("What pet is in the pet shot? ", cat)
	} else {
		log.Println("It is not a cat! ", cat)
	}
	// Use many else if statemates to accomidate multiple outcomes
	myNum := 100
	isTrue2 := false
	// The ! symboly denotes not true
	if myNum > 99 && !isTrue2 {
		log.Println("myNum is true and isTrue2 is false")
	} else if myNum < 100 && isTrue2 {
		log.Println("Dag Dabbit")
	} else if myNum == 101 || isTrue2 {
		log.Println("Rag Babbit")
	} else if myNum > 1000 && isTrue2 == false {
		log.Println("Tag Kabbit")
	}

	// Switch statements are similar to select case

	mySwitchVar := "rat"
	switch mySwitchVar {
	case "cat":
		log.Println("We found a cat!")
	case "dog":
		log.Println("We found a dog!")
	case "rat":
		log.Println("We found a rat!")
	default:
		log.Println("No animals were found")
	}

}
