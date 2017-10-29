package main

import (
	"fmt"
	"os"
	"io/ioutil"
)


func main() {
	// create and initialise our cells
	var cells [30000]uint;
	for i := 0; i < 30000; i++ {
		cells[i] = 0;
	}
	var curr int = 0;

	data, err := ioutil.ReadAll(os.Stdin);
	if err != nil {
		fmt.Println("Could not read stdout");
	}
	program := string(data);

	for i := 0; i < len(program); {
		char := program[i]
		switch char {
		case '>':
			curr = (curr + 1) % 30000;
			i++;
		case '<':
			curr = (curr + 30000 - 1) % 30000;
			i++;
		case '+':
			cells[curr]++;
			i++;
		case '-':
			cells[curr]--;
			i++;
		case '.':
			fmt.Println(string(cells[curr]));
			i++;
		case ',':
			fmt.Scanf("%d", &cells[curr])
			i++;
		case '[':
			if cells[curr] != 0 {
				i++;
			} else {
				for program[i] != ']' {
					i++;
				}
				i++;
			}
		case ']':
			if cells[curr] == 0 {
				i++;
			} else {
				for program[i] != '[' {
					i--;
				}
				i++;
			}
		default:
			i++;
		}
	}
}
