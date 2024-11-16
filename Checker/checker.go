package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	fn "push-swap/myFunctions"
)

func main() {
	operations := []string{}
	arg := []string{}

	if len(os.Args[1:]) == 0 || len(os.Args) > 2 {
		return
	}
	scanner := bufio.NewScanner(os.Stdin)

	arg = strings.Split(os.Args[1], " ")
	a := make([]int, len(arg))

	//convert string to int
	a, err := fn.StringToInt(a, arg)
	if err != nil {
		return
	}

	// Check if duplacated exists or if it's already sorted
	if !fn.Duplicates(a) {
		fmt.Println("Error")
		return
	} else if fn.IfSorted(a) {
		return
	}

	// Loop through each line of input
	for scanner.Scan() {
		// Print the line that was read
		operations = append(operations, scanner.Text())
	}

	if len(operations) == 0 {
		fmt.Println("Error")
		return
	} else if len(operations) == 1 {
		operations = strings.Split(operations[0], "\\n")
	}

	if operations[len(operations)-1] == "" {
		operations = operations[:len(operations)-1]
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
	}

	b := []int{}
	for _, op := range operations {
		switch op {
		case "pa":
			a, b = fn.Pa(a, b)
		case "pb":
			a, b = fn.Pb(a, b)
		case "sa":
			a = fn.Sa(a)
		case "sb":
			b = fn.Sb(b)
		case "ss":
			a, b = fn.Ss(a, b)
		case "ra":
			a = fn.Ra(a)
		case "rb":
			b = fn.Rb(b)
		case "rr":
			a, b = fn.Rr(a, b)
		case "rra":
			a = fn.Rra(a)
		case "rrb":
			b = fn.Rrb(b)
		case "rrr":
			a, b = fn.Rrr(a, b)
		default:
			fmt.Println("Error")
			return
		}
	}

	if len(a) == len(arg) && fn.IfSorted(a) {
		fmt.Println("OK")
		return
	} else {
		fmt.Println("KO")
		return
	}
}
