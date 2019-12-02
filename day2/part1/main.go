package main

import (
	"fmt"
	"github.com/mpaarating/aoc2019/day2/part1/csvprocessor"
	"github.com/mpaarating/aoc2019/day2/part1/intcodemath"
	"log"
)

const dataFile = "data/data.txt"

func main() {
	intCode, err := csvprocessor.ProcessCsv(dataFile)
	if err != nil {
		log.Fatal(err)
	}

	// Per AoC instructions:
	// Reset to "1202 program alarm" state
	// by replacing the two values below
	intCode[1] = 12
	intCode[2] = 2

	currentPosition := 0
	for currentPosition < len(intCode) {
		command := intCode[currentPosition]
		switch command {
			case 1:
				intcodemath.Add(intCode, currentPosition)
			case 2:
				intcodemath.Multiply(intCode, currentPosition)
			case 99:
				break
			default:
				fmt.Errorf("FATAL ERROR: OPCODE NOT FOUND")
		}
		currentPosition += 4
	}

	fmt.Printf("Value at position 0 is: %d \n", intCode[0])
}
