package days

import (
	"AdventOfCode2024/util"
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

/**
 * Description:
 *
 * @author <anas.ahmed@zixel.cn>
 * @date 2024/12/06 10:20 am
 */

func Day06() {
	fmt.Println("Day 06:")
	day6part1()
	day6part2()
}

func day6part1() {
	puzzleMap, gX, gY := day6common()

	steps := 0
	gloc := "N"

	for {
		//fmt.Printf("Gx: %d, Gy: %d, Gloc: %s\n", gX, gY, gloc)
		cond := util.IsObstacleOrStop(puzzleMap, gX, gY, gloc, "#")
		//fmt.Println(cond)

		if cond == -1 {
			// Move forward
			puzzleMap[gX] = puzzleMap[gX][:gY] + "X" + puzzleMap[gX][gY+1:]
			switch gloc {
			case "N":
				gX = gX - 1
			case "S":
				gX = gX + 1
			case "E":
				gY = gY + 1
			case "W":
				gY = gY - 1
			}
		} else if cond == 1 {
			// Rotate 90 degrees to the right
			switch gloc {
			case "N":
				gloc = "E"
			case "S":
				gloc = "W"
			case "E":
				gloc = "S"
			case "W":
				gloc = "N"
			}
		} else {
			puzzleMap[gX] = puzzleMap[gX][:gY] + "X" + puzzleMap[gX][gY+1:]
			break
		}
	}

	//print the map
	for _, line := range puzzleMap {
		steps += strings.Count(line, "X")
	}

	fmt.Println("Part 1: ")
	fmt.Println(steps)
}

func day6part2() {
	puzzleMap, gX, gY := day6common()

	loops := 0
	for i, line := range puzzleMap {
		for j, _ := range line {
			if string(puzzleMap[i][j]) == "#" || (i == gX && j == gY) {
				continue
			}

			newPuzzleMap := util.CopySliceString(puzzleMap)
			newPuzzleMap[i] = newPuzzleMap[i][:j] + "#" + newPuzzleMap[i][j+1:]

			wg := sync.WaitGroup{}
			func() {
				wg.Add(1)
				stuck := StuckInLoop(newPuzzleMap, gX, gY)
				if stuck {
					loops++
				}
				wg.Done()
			}()

			wg.Wait()
		}
	}

	fmt.Println("Part 2: ")
	fmt.Println(loops)
}

func day6common() (puzzleMap []string, gX, gY int) {
	// Read txt file from ../inputs/day01.txt
	file, err := os.Open("inputs/Day06.EXAMPLE")
	if err != nil {
		return
	}

	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			fmt.Println("Error closing file")
		}
	}(file)

	// Convert file to string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		puzzleMap = append(puzzleMap, line)
	}

	for i, line := range puzzleMap {
		j := strings.Index(line, "^")
		if j > -1 {
			gX = i
			gY = j
			break
		}
	}

	return
}

func StuckInLoop(puzzleMap []string, gX, gY int) bool {
	gloc := "N"
	visitedCount := make(map[string]int)
	visitedCount[fmt.Sprintf("%d-%d", gX, gY)] = 1
	for {
		//fmt.Printf("Gx: %d, Gy: %d, Gloc: %s\n", gX, gY, gloc)
		cond := util.IsObstacleOrStop(puzzleMap, gX, gY, gloc, "#")
		//fmt.Println(cond)

		if cond == -1 {
			switch gloc {
			case "N":
				gX = gX - 1
			case "S":
				gX = gX + 1
			case "E":
				gY = gY + 1
			case "W":
				gY = gY - 1
			}
			visitedCount[fmt.Sprintf("%d-%d", gX, gY)]++
			if visitedCount[fmt.Sprintf("%d-%d", gX, gY)] > 10 {
				return true
			}
		} else if cond == 1 {
			// Rotate 90 degrees to the right
			switch gloc {
			case "N":
				gloc = "E"
			case "S":
				gloc = "W"
			case "E":
				gloc = "S"
			case "W":
				gloc = "N"
			}
		} else {
			break
		}
	}

	return false
}
