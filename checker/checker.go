package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//os.Args = append(os.Args, )
	// Init variables
	oparations := []string{}

	// Create a new scanner that reads from standard input
	if len(os.Args[1:]) == 0 {
		return
	}
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter text (press Ctrl+D to end):")

	// Loop through each line of input
	for scanner.Scan() {
		// Print the line that was read
		oparations = append(oparations, scanner.Text())
	}

	if len(oparations) == 1 {
		fmt.Println(1)
		oparations = strings.Split(oparations[0], "\\n")
	}
	fmt.Println(oparations)

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
	}

}
