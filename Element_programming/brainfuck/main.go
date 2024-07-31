package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		// fmt.Println("Usage: go run . \"BrainfuckProgram\"")
		return
	}
	program := os.Args[1]

	memory := make([]byte, 2048)
	fmt.Println(memory)
	ptr := 0
	ip := 0 // Instruction Pointer

	for ip < len(program) {
		fmt.Println(program[ip])
		switch program[ip] {
		case '>':
			ptr++
		case '<':
			ptr--
		case '+':
			memory[ptr]++
		case '-':
			memory[ptr]--
		case '.':
			fmt.Print(string(memory[ptr]))
		case '[':
			if memory[ptr] == 0 {
				// Find matching ]
				nested := 1
				for ip += 1; nested > 0 && ip < len(program); ip++ {
					if program[ip] == '[' {
						nested++
					} else if program[ip] == ']' {
						nested--
					}
				}
			}
		case ']':
			if memory[ptr] != 0 {
				// Find matching [
				nested := 1
				f  - {
					if program[ip] == ']' {
						nested++
					} else if program[ip] == '[' {
						nested--
					}
				}
			}
		}
		ip++
	}
}
