package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// create and initialise our cells
	var cells [30000]uint
	for i := 0; i < 30000; i++ {
		cells[i] = 0
	}
	var curr int = 0

	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println("Could not read stdout")
	}
	program := string(data)

	for i := 0; i < len(program); i++ {
		char := program[i]
		switch char {
		case '>':
			curr = (curr + 1) % 30000
		case '<':
			curr = (curr + 30000 - 1) % 30000
		case '+':
			cells[curr]++
		case '-':
			cells[curr]--
		case '.':
			fmt.Println(string(cells[curr]))
		case ',':
			fmt.Scanf("%d", &cells[curr])
		case '[':
			if cells[curr] == 0 {
				for program[i] != ']' {
					i++
				}
			}
		case ']':
			if cells[curr] != 0 {
				for program[i] != '[' {
					i--
				}
			}
		default:
		}
	}
}
