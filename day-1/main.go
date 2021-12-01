package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var list []int
var c1 = 0
var c2 = 0

func main() {
	f, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
			t := scanner.Text()
			num, err := strconv.ParseInt(t, 0, 64)

			if err != nil {
				fmt.Println(err)
			}

			list = append(list, int(num))
	}

	// part 1
	for i := 1; i < len(list); i++ {
		if list[i] > list[i - 1] {
			c1++
		}
	}

	// part 2
	var prev =  list[0] + list[1] + list[2]
	for i := 3; i < len(list); i++ {
		var new = list[i - 2] + list[i - 1] + list[i - 0]
		if new > prev {
			c2++;
		}
		prev = new
	}

	fmt.Printf("part 1: %v\npart 2: %v", c1, c2)

	if err := scanner.Err(); err != nil {
			log.Fatal(err)
	}
}