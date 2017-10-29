package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

const CELLS int = 30000

func main() {
	// create and initialise our cells
	var cells [CELLS] uint
	for i := 0; i < CELLS; i++ {
		cells[i] = 0
	}
	var curr int = 0

	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	program := string(data)

	for i := 0; i < len(program); i++ {
		switch program[i] {
		case '>':
			curr = (curr + 1) % CELLS
		case '<':
			curr = (curr + CELLS - 1) % CELLS
		case '+':
			cells[curr]++
		case '-':
			cells[curr]--
		case '.':
			fmt.Println(string(cells[curr]))
		case ',':
			_, err := fmt.Scanf("%d", &cells[curr])
			if err != nil {
				panic(err)
			}
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
