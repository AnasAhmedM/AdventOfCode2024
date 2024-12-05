package days

import (
	"AdventOfCode2024/util"
	"bufio"
	"fmt"
	"os"
)

/**
 * Description:
 *
 * @author <anas.ahmed@zixel.cn>
 * @date 2024/12/04 10:00 am
 */

func Day04() {
	fmt.Println("Day 04:")
	day4part1()
	day4part2()
}

func day4part1() {
	matrix := day4common()
	word := "XMAS"
	total := 0
	for i, row := range matrix {
		for j, curr := range row {
			if curr != 'X' {
				continue
			}

			total += util.RecursiveSearch(matrix, i, j, word, "")
		}
	}

	fmt.Println("Part 1: ")
	fmt.Println(total)
}

func day4part2() {
	matrix := day4common()

	total := 0
	for i, row := range matrix {
		for j, curr := range row {
			if i == 0 || j == 0 || i == len(matrix)-1 || j == len(row)-1 {
				continue
			}

			if curr != 'A' {
				continue
			}

			w1 := string(matrix[i-1][j-1]) + "A" + string(matrix[i+1][j+1])
			w2 := string(matrix[i+1][j-1]) + "A" + string(matrix[i-1][j+1])

			if (w1 == "MAS" || w1 == "SAM") && (w2 == "MAS" || w2 == "SAM") {
				total++
			}
		}
	}

	fmt.Println("Part 2: ")
	fmt.Println(total)
}

func day4common() (matrix []string) {
	// Read txt file
	file, err := os.Open("inputs/Day04.txt")
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
		matrix = append(matrix, line)
	}

	return
}
