package main

import (
	"fmt"
)

func main() {
	//	nameOne := "raj"
	//	nameTwo := "ram"
	//	var age int = 25
	//	fmt.Println(nameOne, nameTwo, age)

	var chocolates [5]int = [5]int{1, 2, 3, 4, 5}
	var chocolatesTwo = [10]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"}
	// fmt.Println(chocolates, len(chocolates))
	// fmt.Println(chocolatesTwo)
	// ---- Arrays
	fmt.Println(chocolatesTwo[0:5])
	fmt.Println(chocolates[0:3])
	// ---- Slices
	var name = []string{"raj", "ram", "ravi", "raju", "ravi"}
	name[1] = "nigger"
	name = append(name, "rajesh")
	fmt.Println(name, len(name))

	// 	---- Ranges
	var nameTwo = []string{"UNITY", "Loki", "Dizzy"}
	fmt.Println("Printing ranges from 0 to 2")
	fmt.Println(nameTwo[0:2])
	//////////////
	fmt.Println("Printing ranges from 0 to EOF")
	fmt.Println(nameTwo[0:])
}
