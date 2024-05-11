package main

import (
	"fmt"
)


func main() {

	fmt.Println(one)

	//Playing with go variable rules:
	// 12num := 4 // non-name cannot be used a variable 
	// fmt.Print(12num)

	// first name := "Himal"; //Cannot used space 
	// fmt.Println(first name)


	// itCannotChange = "Let's break rule" //cannot assign to itCannotChange (neither addressable nor a map index expression)
	fmt.Println(itCannotChange)

	checkPasswordStrength("HimalPun123")
	checkPasswordStrength("Himal")
	checkPasswordStrength("Pun")

	fmt.Println("For Loop")
	loop()

	fmt.Println("Continue Statement")
	continueLoop()

	fmt.Println("Break Statement")
	breakLoop()

	fmt.Println("Nested Loop")
	// nestedLoop()
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


//Const Keyword -> Value can't be changed
const itCannotChange = "This variable is constant"





//Conditions -> If else

func checkPasswordStrength (password string ) {
	if len(password) > 7 {
		fmt.Println("It is strong Password")
	}else if len(password) <= 3{
		fmt.Println("It is weak password")
	}else {
		fmt.Println("It is a moderate password")
	}
}




//Loop in Go

// Sudo Syntax 
// for initialExpression; condition;increment {<<code>>}

//1. For Loop
func loop(){
	for  i:=0; i < 5; i++{
		fmt.Println(i)
	}
}


//2. Continue Statement 
//The continue statement is used to skip one or more irerations in loop, it then continous with the next ireteration in the loop

func continueLoop(){
	for i:=0; i<5;i++{
		if(i == 3){
			continue
		}
		fmt.Println(i)
	}
}

//3. Break Statement 

func breakLoop(){
	for i:= 0; i <5; i++{
		if (i==2){
			break
		}
		fmt.Println(i)
	}
}


//4. Nested Loop

func nestedLoop(){
	for i:=0; i< 20; i++{
		fmt.Print("---Outer-----")
		for j:=0; j<10;j++{
			fmt.Println("---Inner----")
		}
	}
}