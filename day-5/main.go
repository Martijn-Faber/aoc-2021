package main

import (
	"fmt"
	"strings"

	"adventofcode/utils"
)

func main() {
	Part1()
	Part2()
}

type Point struct {
	x int
	y int
}

type Line struct {
	start Point
	end Point
}

func Part1() {
	input := utils.ReadInput("input.txt")
	 var lines []Line
	 var maxX int
	 var maxY int

	for _, input := range input {
		split := strings.Split(input, " -> ")
		one := strings.Split(split[0], ",")
		two := strings.Split(split[1], ",")

		line := Line{Point{utils.Atoi(one[0]), utils.Atoi(one[1])}, Point{utils.Atoi(two[0]), utils.Atoi(two[1])}}
		lines = append(lines, line)

		maxX = utils.MaxInt(utils.MaxInt(line.start.x, maxX), line.end.x)
		maxY = utils.MaxInt(utils.MaxInt(line.start.y, maxY), line.end.y)
	}

	diagram := make([][]int, maxY + 1)
	for i := range diagram {
		diagram[i] = make([]int, maxX + 1)
	}

	for _, line := range lines {
		if line.start.x == line.end.x {
			for y := utils.MinInt(line.start.y, line.end.y); y <= utils.MaxInt(line.start.y, line.end.y); y++ {
				diagram[y][line.start.x]++
			}
		} else if line.start.y == line.end.y {
			for x := utils.MinInt(line.start.x, line.end.x); x <= utils.MaxInt(line.start.x, line.end.x); x++ {
				diagram[line.start.y][x]++
			}
		}
	}

	var count int
	for row := 0; row < len(diagram); row++ {
		for col := 0; col < len(diagram[row]); col++ {
			if diagram[row][col] >= 2 {
				count++
			}	
		}
	}

	fmt.Printf("answser part 1: %v\n", count)
}

func Part2() {
	input := utils.ReadInput("input.txt")
	var lines []Line
	var maxX int
	var maxY int
	
 for _, input := range input {
	 split := strings.Split(input, " -> ")
	 one := strings.Split(split[0], ",")
	 two := strings.Split(split[1], ",")

	 line := Line{Point{utils.Atoi(one[0]), utils.Atoi(one[1])}, Point{utils.Atoi(two[0]), utils.Atoi(two[1])}}
	 lines = append(lines, line)

	 maxX = utils.MaxInt(utils.MaxInt(line.start.x, maxX), line.end.x)
	 maxY = utils.MaxInt(utils.MaxInt(line.start.y, maxY), line.end.y)
 }

 diagram := make([][]int, maxY + 1)
 for i := range diagram {
	 diagram[i] = make([]int, maxX + 1)
 }

 for _, line := range lines {
	 if line.start.x == line.end.x {
		 for y := utils.MinInt(line.start.y, line.end.y); y <= utils.MaxInt(line.start.y, line.end.y); y++ {
			 diagram[y][line.start.x]++
		 }
	 } else if line.start.y == line.end.y {
		 for x := utils.MinInt(line.start.x, line.end.x); x <= utils.MaxInt(line.start.x, line.end.x); x++ {
			 diagram[line.start.y][x]++
		 }
	 } else {
		llen := utils.MaxInt(line.start.x, line.end.x) - utils.MinInt(line.start.x, line.end.x)
		
		if line.start.x < line.end.x && line.start.y < line.end.y {
			for i := 0; i <= llen; i++ {
				diagram[line.start.y + i][line.start.x + i]++
			}
		}

		if line.start.x < line.end.x && line.start.y > line.end.y {
			for i := 0; i <= llen; i++ {
				diagram[line.start.y - i][line.start.x + i]++
			}
		}

		if line.start.x > line.end.x && line.start.y < line.end.y {
			for i := 0; i <= llen; i++ {
				diagram[line.start.y + i][line.start.x - i]++
			}
		}

		if line.start.x > line.end.x && line.start.y > line.end.y {
			for i := 0; i <= llen; i++ {
				diagram[line.start.y - i][line.start.x - i]++
			}
		}
	 }
 }


 var count int
 for row := 0; row < len(diagram); row++ {
	 for col := 0; col < len(diagram[row]); col++ {
		 if diagram[row][col] >= 2 {
			 count++
		 }	
	 }
 }

 fmt.Printf("answser part 2: %v", count)
}