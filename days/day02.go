package days

import (
	"AdventOfCode2024/util"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/**
 * Description:
 *
 * @author <anas.ahmed@zixel.cn>
 * @date 2024/12/02 10:14 am
 */

func Day02() {
	fmt.Println("Day 02:")
	day2part1()
	day2part2()
}

func day2part1() {
	matrix := day2common()

	safeTotal := 0
	for _, row := range matrix {
		safe := isSafe(row)

		if safe {
			safeTotal++
		}
	}

	fmt.Println("Part 1:")
	fmt.Println(safeTotal)
}

func day2part2() {
	matrix := day2common()

	safeTotal := 0
	for _, row := range matrix {
		safe := isSafe(row)
		if !safe {
			for i := 0; i < len(row); i++ {
				// Copy row but remove element at index i
				copyRow := make([]int, len(row))
				copy(copyRow, row)
				copyRow = append(copyRow[:i], copyRow[i+1:]...)

				safe = isSafe(copyRow)

				if safe {
					break
				}
			}
		}

		if safe {
			safeTotal++
		}

	}

	fmt.Println("Part 2:")
	fmt.Println(safeTotal)
}

func day2common() (matrix [][]int) {
	// Read txt file
	file, err := os.Open("inputs/Day02.EXAMPLE")
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
	index := 0
	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Split(line, " ")

		matrix = append(matrix, make([]int, len(nums)))

		for i, num := range nums {
			matrix[index][i], _ = strconv.Atoi(num)
		}

		index++
	}

	return matrix
}

func isSafe(row []int) (safe bool) {
	safe = true
	dir := 0 // 1 = increase, -1 = decrease

	if row[0] > row[1] {
		dir = -1
	} else {
		dir = 1
	}

	for i := 0; i < len(row)-1; i++ {
		if dir == 1 && row[i] > row[i+1] {
			safe = false
			break
		} else if dir == -1 && row[i] < row[i+1] {
			safe = false
			break
		}

		rowDiff := util.Abs(row[i] - row[i+1])
		if rowDiff < 1 || rowDiff > 3 {
			safe = false
			break
		}
	}

	return safe
}
