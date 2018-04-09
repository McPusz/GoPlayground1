package main

import "fmt"

func main() {
	fmt.Println("Hello world")
	var age int = 24
	var favNum float64 = 1.23

	fmt.Println(favNum, age)

	var numOne = 1.00
	var num99 = 99.00

	fmt.Println(numOne - num99)

	const pi float64 = 3.14159265

	var(
		varA = 2
		varB = 3
	)

	fmt.Println("Var A + B = ", varA+varB)
	var name string = "Magda"
	fmt.Println("Name length = ", len(name))
	fmt.Println(name + " is a robot")

	fmt.Printf("%.2f \n", pi)
	fmt.Printf("%T is type of favNum\n", favNum)
	var isCatSoft bool = true
	fmt.Printf("%t is value of softCat\n", isCatSoft)
}