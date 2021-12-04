package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var inputs []string
const boardSize = 5

type Board = [boardSize][boardSize]int

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
}

func Part1() {
	var boards []Board
	var lastDrawnNumber int

	numbers := strings.Split(inputs[0], ",")

	// loop through the input data and append it to the 2d boards slice 
	// start at index 2 because the first line are the numbers and the secund line is a space
	// increase i by the boardsize and add the 1 for the space
	for i := 2; i < len(inputs); i+=boardSize+1 {
		var board = Board{}

		for row := 0; row < boardSize; row++ {
			line := strings.Fields(inputs[i + row])

			for col, num := range line {
				num, _ := strconv.ParseInt(num, 0, 64)
				board[row][col] = int(num)
			}
		}

		boards = append(boards, board)
	}

	for num := 0; num < len(numbers); num++ {
		for board := 0; board < len(boards); board++ {
			for row := 0; row < boardSize; row++ {
				for col := 0; col < boardSize; col++ {
					num, _ := strconv.ParseInt(numbers[num], 0, 64)

						if boards[board][row][col] == int(num) {
							// set board index to negative 1
							// these are marked
							boards[board][row][col] = -1
							// set last drawn number
							lastDrawnNumber = int(num)
						}
				}
			}
		}

		// set bingo's default value to negative 1 so we can set the int to the board index
		var bingo = -1
		for board := 0; board < len(boards); board++ {
			// check for bingo in rows
			for row := 0; row < boardSize; row++ {
				var count int
				
				for col := 0; col < boardSize; col++ {
					if boards[board][row][col] == -1 {
						count++
					}

					if count == boardSize {
						// set bingo to the board index
						bingo = board
					}
				}
			}

			// check for bingo in cols
			for col := 0; col < boardSize; col++ {
				var count int
				for row := 0; row < boardSize; row++ {
					if boards[board][row][col] == -1 {
						count++
					}
				}

				if count == boardSize {
					// set bingo to the board index
					bingo = board
				}
			}
		}
		
		

		// check if bingo is greater than 0 because we set bingo to the board index which can't be negative
		if bingo >= 0 {
			var sum int

			// calculate the sum of the unmarked numbers
			for col := 0; col < boardSize; col++ {
				for row := 0; row < boardSize; row++ {
					if boards[bingo][row][col] >= 0	{
						sum += boards[bingo][row][col]
					}	
				}
			}

			fmt.Printf("anwser part 1: %v", sum * lastDrawnNumber)
			break
		}
	}
}