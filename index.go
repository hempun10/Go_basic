package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
	// fmt.Println()
	// punisherSign()
	// fmt.Println(one,two,three,four)
	printName()
	punisherSign()

}



// Variables are containers for storing data values
// In go there are two ways to declare a variable
// 1. Using var keyword
var name string = "Himal"

func printName()  {
	fmt.Println(name)
}


//2. := sign
func punisherSign(){
	surName := "pun"
	fmt.Println(surName)
	
}

//Go Multiple Variable Declare
var one,two,three,four int = 1,2,3,4;




