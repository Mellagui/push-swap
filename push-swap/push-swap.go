package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) > 2 {
		fmt.Println("")
		return
	}

	arg := strings.Split(os.Args[1], " ")
	a := make([]int, len(arg))
	for index, value := range arg {
		intValue, err := strconv.Atoi(value)
		if err != nil {
			fmt.Println("Error")
			return
		}
		a[index] = intValue
	}

	// Check if duplacated exists or if it's already sorted
	if !duplicates(a) {
		fmt.Println("Error")
		return
	} else if ifSorted(a) {
		fmt.Println("")
		return
	}
	pushSwap(a)
}

func pushSwap(a []int) {
	result := []string{}
	b := []int{}
	for !ifSorted(a) && len(a) > 1 {
		mid := len(a) / 2
		index := findMinNbrIndex(a)

		if index == 0 {
			a, b = pb(a, b)
			result = append(result, "pb")
		} else if index < mid && a[1] == a[index] && a[1] < a[0] {
			a = sa(a)
			result = append(result, "sa")
		} else if index < mid {
			a = rra(a)
			result = append(result, "rra")
		} else if a[1] == a[index] && a[1] < a[0] {
			a = sa(a)
			result = append(result, "sa")
		} else {
			a = ra(a)
			result = append(result, "ra")
		}
	}

	for len(b) > 0 {
		a, b = pa(a, b)
		result = append(result, "pa")
	}

	for _, str := range result {
		fmt.Println(str)
	}
	fmt.Println(len(result))
}

func duplicates(a []int) bool {
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] == a[j] {
				return false
			}
		}
	}
	return true
}

func findMinNbrIndex(a []int) int {
	index := 0
	for i := 1; i < len(a); i++ {
		if a[i] < a[index] {
			index = i
		}
	}
	return index
}

func ifSorted(a []int) bool {
	for i := 0; i < len(a)-1; i++ {
		if a[i] > a[i+1] {
			return false
		}
	}
	return true
}

// push top element in stack b to stack a
func pa(a []int, b []int) ([]int, []int) {
	if len(b) > 1 {
		a, b = append([]int{b[0]}, a...), b[1:]
	} else {
		a, b = append([]int{b[0]}, a...), []int{}
	}
	return a, b
}

// push top element in stack a to stack b
func pb(a []int, b []int) ([]int, []int) {
	if len(a) > 1 {
		b, a = append([]int{a[0]}, b...), a[1:]
	} else {
		b, a = append([]int{a[0]}, b...), []int{}
	}
	return a, b
}

// swap 2 top elements of stack a
func sa(a []int) []int {
	if len(a) > 1 {
		a[0], a[1] = a[1], a[0]
	}
	return a
}

// swap 2 top elements of stack b
func sb(b []int) []int {
	if len(b) > 1 {
		b[0], b[1] = b[1], b[0]

	}
	return b
}

// swap 2 top elements in stack a & stack b
func ss(a []int, b []int) ([]int, []int) {
	sa(a)
	sb(b)
	return a, b
}

// reverse last element to first element in stack a
func ra(a []int) ([]int) {
	if len(a) > 1 {
		a = append([]int{a[len(a)-1]}, a[:len(a)-1]...)
	}
	return a
}

// reverse last element to first element in stack b
func rb(b []int) ([]int) {
	if len(b) > 1 {
		b = append([]int{b[len(b)-1]}, b[:len(b)-1]...)
	}
	return b
}

// reverse last element to first element in stack a & stack b
func rr(a []int, b []int) ([]int, []int) {
	a = ra(a)
	b = rb(b)
	return a, b
}

// reverse first element to last element in stack a
func rra(a []int) []int {
	if len(a) > 1 {
		a = append(a[1:], a[0])
	}
	return a
}

// reverse first element to last element in stack b
func rrb(b []int) []int {
	if len(b) > 1 {
		b = append(b[1:], b[0])
	}
	return b
}

// reverse first element to last element in stack a & stack b
func rrr(a []int, b []int) ([]int, []int) {
	a = rra(a)
	b = rrb(b)
	return a, b
}
