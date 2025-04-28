package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("=== Welcome to CLI Calculator ===")
	fmt.Println("Available operations: +, -, *, /, sqrt, pow")
	fmt.Println("----------------------------------")

	for {
		fmt.Print("Enter first number (or type 'exit' to quit): ")
		input1, _ := reader.ReadString('\n')
		input1 = strings.TrimSpace(input1)
		if input1 == "exit" {
			fmt.Println("Goodbye!")
			break
		}
		num1, err := strconv.ParseFloat(input1, 64)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}

		fmt.Print("Enter operation (+, -, *, /, sqrt, pow): ")
		op, _ := reader.ReadString('\n')
		op = strings.TrimSpace(op)

		var num2 float64
		if op != "sqrt" {
			fmt.Print("Enter second number: ")
			input2, _ := reader.ReadString('\n')
			input2 = strings.TrimSpace(input2)
			num2, err = strconv.ParseFloat(input2, 64)
			if err != nil {
				fmt.Println("Invalid input. Please enter a number.")
				continue
			}
		}

		switch op {
		case "+":
			fmt.Printf("Result: %.2f\n", num1+num2)
		case "-":
			fmt.Printf("Result: %.2f\n", num1-num2)
		case "*":
			fmt.Printf("Result: %.2f\n", num1*num2)
		case "/":
			if num2 == 0 {
				fmt.Println("Error: Division by zero!")
			} else {
				fmt.Printf("Result: %.2f\n", num1/num2)
			}
		case "sqrt":
			if num1 < 0 {
				fmt.Println("Error: Cannot take square root of a negative number!")
			} else {
				fmt.Printf("Result: %.2f\n", math.Sqrt(num1))
			}
		case "pow":
			fmt.Printf("Result: %.2f\n", math.Pow(num1, num2))
		default:
			fmt.Println("Unknown operation.")
		}

		fmt.Println("----------------------------------")
	}
}
