package main

import (
	"adventofcode/utils"
	"fmt"
	"strings"
)

const boardSize = 5

type Board = [boardSize][boardSize]int

func main() {
	// Part1()
	Part2()
}

func Part1() {
	input := utils.ReadInput("input.txt")
	var boards []Board
	var lastDrawnNumber int

	numbers := strings.Split(input[0], ",")

	// loop through the input data and append it to the 2d boards slice 
	// start at index 1 because the first line are the numbers
	for i := 1; i < len(input); i+=boardSize {
		var board = Board{}

		for row := 0; row < boardSize; row++ {
			line := strings.Fields(input[i + row])

			for col, num := range line {
				board[row][col] = utils.Atoi(num)
			}
		}

		boards = append(boards, board)
	}

	for num := 0; num < len(numbers); num++ {
		for board := 0; board < len(boards); board++ {
			for row := 0; row < boardSize; row++ {
				for col := 0; col < boardSize; col++ {
						if boards[board][row][col] == utils.Atoi(numbers[num]) {
							// set board index to negative 1
							// these are marked
							boards[board][row][col] = -1
							// set last drawn number
							lastDrawnNumber = utils.Atoi(numbers[num])
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

func Part2() {
	input := utils.ReadInput("input.txt")
	var boards []Board
	var lastDrawnNumber int

	numbers := strings.Split(input[0], ",")

	// loop through the input data and append it to the 2d boards slice 
	// start at index 1 because the first line are the numbers
	for i := 1; i < len(input); i+=boardSize {
		var board = Board{}

		for row := 0; row < boardSize; row++ {
			line := strings.Fields(input[i + row])

			for col, num := range line {
				board[row][col] = utils.Atoi(num)
			}
		}

		boards = append(boards, board)
	}

	for num := 0; num < len(numbers); num++ {
		for board := 0; board < len(boards); board++ {
			for row := 0; row < boardSize; row++ {
				for col := 0; col < boardSize; col++ {
						if boards[board][row][col] == utils.Atoi(numbers[num]) {
							// set board index to negative 1
							// these are marked
							boards[board][row][col] = -1
							// set last drawn number
							lastDrawnNumber = utils.Atoi(numbers[num])
						}
				}
			}
		}

		// set bingo's default value to negative 1 so we can set the int to the board index
		var bingo = -1
		var remainingBoards []Board

		for board := 0; board < len(boards); board++ {
			bingo = -1
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

			if bingo < 0 {
				remainingBoards = append(remainingBoards, boards[board])
			}
		}

		// check if bingo is greater than 0 because we set bingo to the board index which can't be negative
		if bingo >= 0 && len(boards) == 1 {
			var sum int

			// calculate the sum of the unmarked numbers
			for col := 0; col < boardSize; col++ {
				for row := 0; row < boardSize; row++ {
					if boards[bingo][row][col] >= 0	{
						sum += boards[bingo][row][col]
					}
				}
			}

			fmt.Printf("anwser part 2: %v", sum * lastDrawnNumber)
			break
		} else {
			fmt.Println(len(remainingBoards))
			boards = remainingBoards
		}
	}
}