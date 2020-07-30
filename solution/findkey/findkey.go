package main

import (
	"fmt"
	"time"
)

var maze = [6][8]string{
	{"#", "#", "#", "#", "#", "#", "#", "#"},
	{"#", ".", ".", ".", ".", ".", ".", "#"},
	{"#", ".", "#", "#", "#", ".", ".", "#"},
	{"#", ".", ".", ".", "#", ".", "#", "#"},
	{"#", ".", "#", ".", ".", ".", ".", "#"},
	{"#", "#", "#", "#", "#", "#", "#", "#"},
}

const enableDebug = false

var (
	scanPosRow = -1
	scanPosCol = -1

	northStep = 0
	eastStep  = 0
	southStep = 0

	enableNorth bool
	enableEast  bool
	enableSouth bool

	isStop = false
	solved = false

	solutions    [][6][8]string
	tempSolution [6][8]string
)

func scan(posRow, posCol int) bool {
	if posRow == 0 && posCol == len(maze[0])-1 {
		isStop = true
	}
	if posRow < 0 && posCol < 0 {
		posRow = len(maze) - 1
		posCol = 0
	}

	if enableDebug {
		time.Sleep(2 * time.Second)
		displayMaze := maze
		displayMaze[posRow][posCol] = "X"
		fmt.Println()
		for _, m := range displayMaze {
			fmt.Println(m)
		}
	}

	if maze[posRow][posCol] == "." {
		enableNorth = true
		enableEast = false
		enableSouth = false
		northStep = 0
		eastStep = 0
		southStep = 0
		solved = false
		scanPosRow = posRow
		scanPosCol = posCol
		return true
	}

	if maze[posRow][posCol] == "#" {
		if len(maze[posRow]) > posCol+1 {
			scan(posRow, posCol+1)
		} else {
			if posRow > 0 {
				posRow--
				posCol = 0
				scan(posRow, posCol)
			}
		}
	}
	return false
}

func move(posRow, posCol int) bool {
	if solved {
		return false
	}

	tempSolution[posRow][posCol] = "X"
	if enableDebug {
		time.Sleep(2 * time.Second)
		displayMaze := maze
		displayMaze[posRow][posCol] = "X"
		fmt.Println()
		for _, m := range displayMaze {
			fmt.Println(m)
		}
	}

	// got solution
	if enableSouth && southStep > 0 && posRow+1 < len(maze) && maze[posRow+1][posCol] == "#" {
		return true
	}
	// check south after east step
	if enableEast && eastStep > 0 && posRow+1 < len(maze) && maze[posRow+1][posCol] == "." {
		enableSouth = true
		enableEast = false
		enableNorth = false
	}
	// check east
	if enableNorth && northStep > 0 && posCol-1 >= 0 && maze[posRow][posCol-1] == "." {
		enableEast = true
		enableNorth = false
		enableSouth = false
	}

	if enableNorth {
		if maze[posRow-1][posCol] == "." {
			northStep++
			move(posRow-1, posCol)
		} else {
			return false
		}
	}
	if enableEast {
		if maze[posRow][posCol-1] == "." {
			eastStep++
			move(posRow, posCol-1)
		}
	}
	if enableSouth {
		if maze[posRow+1][posCol] == "." {
			southStep++
			if move(posRow+1, posCol) {
				solved = true
				solutions = append(solutions, tempSolution)
				return false
			}
		}
	}
	return false
}

func process() {
	scan(scanPosRow, scanPosCol+1)
	if !isStop {
		tempSolution = maze
		move(scanPosRow, scanPosCol)
		process()
	}
}

func displaySolution() {
	fmt.Println("\nSolution")
	mergedSolution := maze
	for index, solution := range solutions {
		fmt.Println("Solution", index+1)
		for i, m := range solution {
			fmt.Println(m)
			for j, c := range m {
				if c == "X" {
					mergedSolution[i][j] = c
				}
			}
		}
		fmt.Println("")
	}
	totalPoint := 0
	fmt.Println("\nMerged Solution")
	for _, merged := range mergedSolution {
		fmt.Println(merged)
		for _, char := range merged {
			if char == "X" {
				totalPoint++
			}
		}
	}
	fmt.Printf("\nTotal for probability of key = %d point", totalPoint)
}

func main() {
	scan(scanPosRow, scanPosCol)
	move(scanPosRow, scanPosCol)
	process()
	displaySolution()
}
