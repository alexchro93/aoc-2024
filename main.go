package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/alexchro93/aoc-2024/day16"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: solutions <day>")
		return
	}

	day, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Invalid day:", os.Args[1])
		return
	}

	switch day {
	case 16:
		day16.Run()
	default:
		fmt.Printf("Day %d not implemented yet.\n", day)
	}
}
