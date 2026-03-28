package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to CLI calculator!")
	for {
		PresentOptions()
		var choice int
		fmt.Scan(&choice)

		if choice == 1 {
			var num1 float64
			var num2 float64
			fmt.Print("Input first number: ")
			_, err1 := fmt.Scanln(&num1)
			if err1 != nil {
				fmt.Println("Error: please input a valid number")
				continue
			}
			fmt.Print("Input Second number: ")
			_, err2 := fmt.Scanln(&num2)
			if err2 != nil {
				fmt.Println("Error: please input a valid number")
				continue
			}
			result := Addition(num1, num2)
			fmt.Println("Result: ", result)
		}
		if choice == 2 {
			var num1 float64
			var num2 float64
			fmt.Print("Input first number: ")
			_, err1 := fmt.Scanln(&num1)
			if err1 != nil {
				fmt.Println("Error: please input a valid number")
				continue
			}
			fmt.Print("Input Second number: ")
			_, err2 := fmt.Scanln(&num2)
			if err2 != nil {
				fmt.Println("Error: please input a valid number")
				continue
			}
			result := Subtraction(num1, num2)
			fmt.Println("Result: ", result)
		}
		if choice == 3 {
			var num1 float64
			var num2 float64
			fmt.Print("Input first number: ")
			_, err1 := fmt.Scanln(&num1)
			if err1 != nil {
				fmt.Println("Error: please input a valid number")
				continue
			}
			fmt.Print("Input Second number: ")
			_, err2 := fmt.Scanln(&num2)
			if err2 != nil {
				fmt.Println("Error: please input a valid number")
				continue
			}
			result := Multiplication(num1, num2)
			fmt.Println("Result: ", result)
		}
		if choice == 4 {
			var num1 float64
			var num2 float64
			fmt.Print("Input first number: ")
			_, err1 := fmt.Scanln(&num1)
			if err1 != nil {
				fmt.Println("Error: please input a valid number")
				continue
			}
			fmt.Print("Input Second number: ")
			_, err2 := fmt.Scanln(&num2)
			if err2 != nil {
				fmt.Println("Error: please input a valid number")
				continue
			}
			result, err := Division(num1, num2)
			if err != nil {
				fmt.Println("Error: ", err)
				continue
			}
			fmt.Println("Result: ", result)
		}
		if choice > 5 || choice < 1 {
			fmt.Println("Unknown command, Can Only use Operations from 1-5")
		}
		if choice == 5 {
			fmt.Println("GoodBye!")
			break
		}
	}
}
