package main

import (
	"fmt"
	"github.com/mpaarating/aoc2019/day2/part1/intcodemath"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

const dataFile = "../part1/data/data.txt"

func main() {
	input := processInput(dataFile)

	Loop:
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			result := calculate(input, noun, verb)
			if result == 19690720 {
				fmt.Printf("%d%d\n", noun, verb)
				break Loop
			}
		}
	}
}

func processInput(path string) []int {
	inputBytes, err := ioutil.ReadFile(dataFile)
	if err != nil {
		log.Fatal(err)
	}
	text := strings.TrimSpace(string(inputBytes))
	var inputs []int
	for _, value := range strings.Split(text, ",") {
		inputNum, err := strconv.Atoi(value)
		if err != nil {
			log.Fatal(err)
		}
		inputs = append(inputs, inputNum)
	}
	return inputs
}

func calculate(input []int, noun int, verb int) int {
	memory := make([]int, len(input))
	copy(memory, input)

	memory[1] = noun
	memory[2] = verb
	instructionAddress := 0
	for {
		command := memory[instructionAddress]
		switch command {
		case 1:
			intcodemath.Add(memory, instructionAddress)
			instructionAddress += 4
		case 2:
			intcodemath.Multiply(memory, instructionAddress)
			instructionAddress += 4
		case 99:
			return memory[0]
		default:
			return 0
		}
	}
}
