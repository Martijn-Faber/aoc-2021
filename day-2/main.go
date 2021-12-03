package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var inputs []string

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
	var forward = 0
	var depth = 0

	for _, input := range inputs {
			split := strings.Split(input, " ")
			pos := split[0]
			num, err := strconv.ParseInt(split[1], 0, 0)

			if err != nil {
				fmt.Println(err)
			}

			switch pos {
			case "forward":
				forward += int(num)
			case "down":
				depth += int(num)
			case "up":
				depth -= int(num)
			}
	}

	fmt.Printf("part 1: %v\n", forward * depth)
}

func Part2() {
	var forward = 0
	var depth = 0
	var aim = 0

	for _, input := range inputs {
			split := strings.Split(input, " ")
			pos := split[0]
			num, err := strconv.ParseInt(split[1], 0, 0)

			if err != nil {
				fmt.Println(err)
			}

			switch pos {
			case "forward":
				forward += int(num)
				depth += aim * int(num)
			case "down":
				aim += int(num)
			case "up":
				aim -= int(num)
			}
	}

	fmt.Printf("part 2: %v\n", forward * depth)
}