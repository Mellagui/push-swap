package main

import (
	"fmt"
	"os"
	"strings"

	fn "push-swap/myFunctions"
)

func main() {
	if len(os.Args) > 2 || len(os.Args) == 1 {
		fmt.Print("")
		return
	}

	arg := strings.Split(os.Args[1], " ")
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
	pushSwap(a)
}

func pushSwap(a []int) {
	b := []int{}
	lenMax := len(a)
	result := []string{}

	for !fn.IfSorted(a) && len(a) > 1 {
		mid := len(a) / 2
		index := fn.FindMinNbrIndex(a)

		if index == 0 {
			a, b = fn.Pb(a, b)
			result = append(result, "pb")
		} else if index < mid && a[1] == a[index] && a[1] < a[0] {
			a = fn.Sa(a)
			result = append(result, "sa")
		} else if index < mid {
			if lenMax > 4 {
				a = fn.Ra(a)
				result = append(result, "ra")
			} else {
				a = fn.Rra(a)
				result = append(result, "rra")
			}
		} else if a[1] == a[index] && a[1] < a[0] {
			a = fn.Sa(a)
			result = append(result, "sa")
		} else {
			if lenMax > 4 {
				a = fn.Rra(a)
				result = append(result, "rra")
			} else {
				a = fn.Ra(a)
				result = append(result, "ra")
			}
		}
	}

	for len(a) != lenMax && len(b) > 0 {
		a, b = fn.Pa(a, b)
		result = append(result, "pa")
	}

	for _, str := range result {
		fmt.Println(str)
	}
}