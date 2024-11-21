package myFunctions

import (
	"fmt"
	"strconv"
)

func StringToInt(a []int, arg []string) ([]int, error) {
	for index, str := range arg {
		nbr, err := strconv.Atoi(str)
		if err != nil {
			fmt.Println("Error")
			return a, err
		}
		a[index] = nbr
	}
	return a, nil
}

func Duplicates(a []int) bool {
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] == a[j] {
				return false
			}
		}
	}
	return true
}

func FindMinNbrIndex(a []int) int {
	index := 0
	for i := 1; i < len(a); i++ {
		if a[i] < a[index] {
			index = i
		}
	}
	return index
}

func IfSorted(a []int) bool {
	for i := 0; i < len(a)-1; i++ {
		if a[i] > a[i+1] {
			return false
		}
	}
	return true
}

// push top element in stack b to stack a
func Pa(a []int, b []int) ([]int, []int) {
	if len(b) > 1 {
		a, b = append([]int{b[0]}, a...), b[1:]
	} else if len(b) == 1 {
		a, b = append([]int{b[0]}, a...), []int{}
	}
	return a, b
}

// push top element in stack a to stack b
func Pb(a []int, b []int) ([]int, []int) {
	if len(a) >= 1 {
		b, a = append([]int{a[0]}, b...), a[1:]
	}
	return a, b
}

// swap 2 top elements of stack a
func Sa(a []int) []int {
	if len(a) > 1 {
		a[0], a[1] = a[1], a[0]
	}
	return a
}

// swap 2 top elements of stack b
func Sb(b []int) []int {
	if len(b) > 1 {
		b[0], b[1] = b[1], b[0]
	}
	return b
}

// swap 2 top elements in stack a & stack b
func Ss(a []int, b []int) ([]int, []int) {
	Sa(a)
	Sb(b)
	return a, b
}

// reverse first element to last element in stack a
func Ra(a []int) []int {
	if len(a) > 1 {
		a = append(a[1:], a[0])
	}
	return a
}

// reverse first element to last element in stack b
func Rb(b []int) []int {
	if len(b) > 1 {
		b = append(b[1:], b[0])
	}
	return b
}

// reverse first element to last element in stack a & stack b
func Rr(a []int, b []int) ([]int, []int) {
	a = Ra(a)
	b = Rb(b)
	return a, b
}

// reverse last element to first element in stack a
func Rra(a []int) []int {
	if len(a) > 1 {
		a = append([]int{a[len(a)-1]}, a[:len(a)-1]...)
	}
	return a
}

// reverse last element to first element in stack b
func Rrb(b []int) []int {
	if len(b) > 1 {
		b = append([]int{b[len(b)-1]}, b[:len(b)-1]...)
	}
	return b
}

// reverse last element to first element in stack a & stack b
func Rrr(a []int, b []int) ([]int, []int) {
	a = Rra(a)
	b = Rrb(b)
	return a, b
}