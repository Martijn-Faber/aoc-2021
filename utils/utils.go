package utils

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func Atoi(str string) int {
	num, err := strconv.Atoi(str)
	
	if err != nil {
		panic(err)
	}

	return num
}

func MaxInt(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func MinInt(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func ReadInput(fileName string) (result []string) {
	file, err := os.Open(fileName)

	if err != nil {
		panic(err)
	}

	defer func(file *os.File) {
		err := file.Close()

		if err != nil {
			panic(err)
		}
	}(file)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if line != "" {
			result = append(result, line)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return
}