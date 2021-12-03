package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var inputs []string
var bits = 12

func main() {
	f, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		t := scanner.Text()
		inputs = append(inputs, t)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	Part1()
	Part2()
}

func Part1() {
	// power consumption
	var pc = 0
	var resultG string
	var resultE string

	for j := 0; j < bits; j++ {
		var zeros = CountBits(inputs, j, "0")
		var ones = CountBits(inputs, j, "1")

		if ones > zeros {
			resultG += "1"
			resultE += "0"
		} else {
			resultG += "0"
			resultE += "1"
		}
	}

	pc = BinToDec(resultG) * BinToDec(resultE)
	fmt.Printf("result part 1: %v\n\n", pc)
}

func Part2() {
	var inputsOX = inputs
	var inputsCO = inputs
	// oxygen generator rating
	var OXgr = 0
	// CO2 scrubber rating
	var COsr = 0

	for k := 0; k < bits; k++ {
		var zeros = CountBits(inputsOX, k, "0")
		var ones = CountBits(inputsOX, k, "1")

		var newList = inputsOX
		inputsOX = make([]string, 0)

		if ones >= zeros {
			for _, input := range newList {
				if input[k:k+1] == "1" {
					inputsOX = append(inputsOX, input)
				}
			}
		} else {
			for _, input := range newList {
				if input[k:k+1] == "0" {
					inputsOX = append(inputsOX, input)
				}
			}
		}

		if len(inputsOX) == 1 {
			OXgr = BinToDec(inputsOX[0])
			// fmt.Printf("part 2 ox:\nbinary: %v\ndecimal: %v\n\n", inputsOX[0], OXgr)
			break
		}
	}

	for k := 0; k < bits; k++ {
		var zeros = CountBits(inputsCO, k, "0")
		var ones = CountBits(inputsCO, k, "1")

		var newList = inputsCO
		inputsCO = make([]string, 0)

		if zeros <= ones {
			for _, input := range newList {
				if input[k:k+1] == "0" {
					inputsCO = append(inputsCO, input)
				}
			}
		} else {
			for _, input := range newList {
				if input[k:k+1] == "1" {
					inputsCO = append(inputsCO, input)
				}
			}
		}

		if len(inputsCO) == 1 {
			COsr = BinToDec(inputsCO[0])
			// fmt.Printf("part 2 co:\nbinary: %v\ndecimal: %v\n\n", inputsCO[0], COsr)
			break
		}
	}

	fmt.Printf("result part 2: %v", COsr * OXgr)
}

func BinToDec(binarystr string) int {
	dec, err := strconv.ParseInt(binarystr, 2, 64)

	if err != nil {
		fmt.Println(err)
	}

	return int(dec)
}

func CountBits(arr []string, col int, bitval string) (count int) {
	for _, input := range arr {
		if input[col:col+1] == bitval {
			count++
		}
	}
	return
}
