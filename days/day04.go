package days

import (
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
	day4part2solution1()
	day4part2solution2()
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

			total += recursiveSearch(matrix, i, j, word[1:], "")
		}
	}

	fmt.Println("Part 1: ")
	fmt.Println(total)
}

func day4part2solution1() {
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

			mCount := 0
			sCount := 0
			if checkNorthWest(matrix, i, j, "M") {
				mCount++
			}

			if checkNorthEast(matrix, i, j, "M") {
				mCount++
			}

			if checkSouthWest(matrix, i, j, "M") {
				mCount++
			}

			if checkSouthEast(matrix, i, j, "M") {
				mCount++
			}

			if checkNorthWest(matrix, i, j, "S") {
				sCount++
			}

			if checkNorthEast(matrix, i, j, "S") {
				sCount++
			}

			if checkSouthWest(matrix, i, j, "S") {
				sCount++
			}

			if checkSouthEast(matrix, i, j, "S") {
				sCount++
			}

			if mCount == 2 && sCount == 2 {
				if !(checkNorthWest(matrix, i, j, "S") && checkSouthEast(matrix, i, j, "S")) && !(checkNorthWest(matrix, i, j, "M") && checkSouthEast(matrix, i, j, "M")) {
					total++
				}

			}
		}
	}

	fmt.Println("Part 2, Solution 1: ")
	fmt.Println(total)
}

func day4part2solution2() {
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

	fmt.Println("Part 2, Solution 2: ")
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

func recursiveSearch(matrix []string, i, j int, word string, setDirection string) (total int) {

	nextLetter := string(word[0])

	locs := checkAllDirections(matrix, i, j, nextLetter, setDirection)

	if len(word) == 1 {
		return len(locs)
	}

	for _, loc := range locs {
		switch loc {
		case "N":
			total += recursiveSearch(matrix, i-1, j, word[1:], "N")
		case "S":
			total += recursiveSearch(matrix, i+1, j, word[1:], "S")
		case "E":
			total += recursiveSearch(matrix, i, j+1, word[1:], "E")
		case "W":
			total += recursiveSearch(matrix, i, j-1, word[1:], "W")
		case "NE":
			total += recursiveSearch(matrix, i-1, j+1, word[1:], "NE")
		case "NW":
			total += recursiveSearch(matrix, i-1, j-1, word[1:], "NW")
		case "SE":
			total += recursiveSearch(matrix, i+1, j+1, word[1:], "SE")
		case "SW":
			total += recursiveSearch(matrix, i+1, j-1, word[1:], "SW")
		}
	}

	return total
}

func checkAllDirections(matrix []string, i, j int, nextLetter string, setDirection string) (locs []string) {
	if (setDirection == "" || setDirection == "N") && checkNorth(matrix, i, j, nextLetter) {
		locs = append(locs, "N")
	}
	if (setDirection == "" || setDirection == "S") && checkSouth(matrix, i, j, nextLetter) {
		locs = append(locs, "S")
	}
	if (setDirection == "" || setDirection == "E") && checkEast(matrix, i, j, nextLetter) {
		locs = append(locs, "E")
	}
	if (setDirection == "" || setDirection == "W") && checkWest(matrix, i, j, nextLetter) {
		locs = append(locs, "W")
	}
	if (setDirection == "" || setDirection == "NE") && checkNorthEast(matrix, i, j, nextLetter) {
		locs = append(locs, "NE")
	}
	if (setDirection == "" || setDirection == "NW") && checkNorthWest(matrix, i, j, nextLetter) {
		locs = append(locs, "NW")
	}
	if (setDirection == "" || setDirection == "SE") && checkSouthEast(matrix, i, j, nextLetter) {
		locs = append(locs, "SE")
	}
	if (setDirection == "" || setDirection == "SW") && checkSouthWest(matrix, i, j, nextLetter) {
		locs = append(locs, "SW")
	}
	return
}

func checkEast(matrix []string, i, j int, nextLetter string) bool {
	if j+1 < len(matrix[i]) && string(matrix[i][j+1]) == nextLetter {
		return true
	}
	return false
}

func checkWest(matrix []string, i, j int, nextLetter string) bool {
	if j-1 >= 0 && string(matrix[i][j-1]) == nextLetter {
		return true
	}
	return false
}

func checkNorth(matrix []string, i, j int, nextLetter string) bool {
	if i-1 >= 0 && string(matrix[i-1][j]) == nextLetter {
		return true
	}
	return false
}

func checkSouth(matrix []string, i, j int, nextLetter string) bool {
	if i+1 < len(matrix) && string(matrix[i+1][j]) == nextLetter {
		return true
	}
	return false
}

func checkNorthEast(matrix []string, i, j int, nextLetter string) bool {
	if i-1 >= 0 && j+1 < len(matrix[i]) && string(matrix[i-1][j+1]) == nextLetter {
		return true
	}
	return false
}

func checkNorthWest(matrix []string, i, j int, nextLetter string) bool {
	if i-1 >= 0 && j-1 >= 0 && string(matrix[i-1][j-1]) == nextLetter {
		return true
	}
	return false
}

func checkSouthEast(matrix []string, i, j int, nextLetter string) bool {
	if i+1 < len(matrix) && j+1 < len(matrix[i]) && string(matrix[i+1][j+1]) == nextLetter {
		return true
	}
	return false
}

func checkSouthWest(matrix []string, i, j int, nextLetter string) bool {
	if i+1 < len(matrix) && j-1 >= 0 && string(matrix[i+1][j-1]) == nextLetter {
		return true
	}
	return false
}
