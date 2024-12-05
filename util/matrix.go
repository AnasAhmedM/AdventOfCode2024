package util

/**
 * Description:
 *
 * @author <anas.ahmed@zixel.cn>
 * @date 2024/12/04 9:20 pm
 */

// RecursiveSearch searches for a word in a matrix
// @param matrix: a 2D array of strings
// @param i: the row index
// @param j: the column index
// @param word: the word to search for, Has to be unique letters like ABC not AABC
// @param setDirection: the direction to search in, if empty it will search in all directions
// @return int: the number of times the word was found
func RecursiveSearch(matrix []string, i, j int, word string, setDirection string) (total int) {

	if len(matrix) == 0 || len(word) == 0 || i >= len(matrix) || j >= len(matrix[i]) {
		return -1
	}

	nextLetter := string(word[0])
	if string(matrix[i][j]) == nextLetter {
		total += RecursiveSearch(matrix, i, j, word[1:], setDirection)
		return
	}

	locs := GetAllDirections(matrix, i, j, nextLetter, setDirection)

	if len(word) == 1 {
		return len(locs)
	}

	for _, loc := range locs {
		switch loc {
		case "N":
			total += RecursiveSearch(matrix, i-1, j, word[1:], "N")
		case "S":
			total += RecursiveSearch(matrix, i+1, j, word[1:], "S")
		case "E":
			total += RecursiveSearch(matrix, i, j+1, word[1:], "E")
		case "W":
			total += RecursiveSearch(matrix, i, j-1, word[1:], "W")
		case "NE":
			total += RecursiveSearch(matrix, i-1, j+1, word[1:], "NE")
		case "NW":
			total += RecursiveSearch(matrix, i-1, j-1, word[1:], "NW")
		case "SE":
			total += RecursiveSearch(matrix, i+1, j+1, word[1:], "SE")
		case "SW":
			total += RecursiveSearch(matrix, i+1, j-1, word[1:], "SW")
		}
	}

	return total
}

// GetAllDirections returns all the directions that the next letter can be found in
// @param matrix: a 2D array of strings
// @param i: the row index
// @param j: the column index
// @param nextLetter: the letter to search for
// @param setDirection: the direction to search in, if empty it will search in all directions
// @return []string: the directions the next letter can be found in
func GetAllDirections(matrix []string, i, j int, nextLetter string, setDirection string) (locs []string) {
	if (setDirection == "" || setDirection == "N") && CheckNorth(matrix, i, j, nextLetter) {
		locs = append(locs, "N")
	}
	if (setDirection == "" || setDirection == "S") && CheckSouth(matrix, i, j, nextLetter) {
		locs = append(locs, "S")
	}
	if (setDirection == "" || setDirection == "E") && CheckEast(matrix, i, j, nextLetter) {
		locs = append(locs, "E")
	}
	if (setDirection == "" || setDirection == "W") && CheckWest(matrix, i, j, nextLetter) {
		locs = append(locs, "W")
	}
	if (setDirection == "" || setDirection == "NE") && CheckNorthEast(matrix, i, j, nextLetter) {
		locs = append(locs, "NE")
	}
	if (setDirection == "" || setDirection == "NW") && CheckNorthWest(matrix, i, j, nextLetter) {
		locs = append(locs, "NW")
	}
	if (setDirection == "" || setDirection == "SE") && CheckSouthEast(matrix, i, j, nextLetter) {
		locs = append(locs, "SE")
	}
	if (setDirection == "" || setDirection == "SW") && checkSouthWest(matrix, i, j, nextLetter) {
		locs = append(locs, "SW")
	}
	return
}

// CheckEast checks if the next letter is to the east of the current letter
// @param matrix: a 2D array of strings
// @param i: the row index
// @param j: the column index
// @param nextLetter: the letter to search for
// @return bool: true if the next letter is to the east of the current letter
func CheckEast(matrix []string, i, j int, nextLetter string) bool {
	if j+1 < len(matrix[i]) && string(matrix[i][j+1]) == nextLetter {
		return true
	}
	return false
}

// CheckWest checks if the next letter is to the west of the current letter
// @param matrix: a 2D array of strings
// @param i: the row index
// @param j: the column index
// @param nextLetter: the letter to search for
// @return bool: true if the next letter is to the west of the current letter
func CheckWest(matrix []string, i, j int, nextLetter string) bool {
	if j-1 >= 0 && string(matrix[i][j-1]) == nextLetter {
		return true
	}
	return false
}

// CheckNorth checks if the next letter is to the north of the current letter
// @param matrix: a 2D array of strings
// @param i: the row index
// @param j: the column index
// @param nextLetter: the letter to search for
// @return bool: true if the next letter is to the north of the current letter
func CheckNorth(matrix []string, i, j int, nextLetter string) bool {
	if i-1 >= 0 && string(matrix[i-1][j]) == nextLetter {
		return true
	}
	return false
}

// CheckSouth checks if the next letter is to the south of the current letter
// @param matrix: a 2D array of strings
// @param i: the row index
// @param j: the column index
// @param nextLetter: the letter to search for
// @return bool: true if the next letter is to the south of the current letter
func CheckSouth(matrix []string, i, j int, nextLetter string) bool {
	if i+1 < len(matrix) && string(matrix[i+1][j]) == nextLetter {
		return true
	}
	return false
}

// CheckNorthEast checks if the next letter is to the north east of the current letter
// @param matrix: a 2D array of strings
// @param i: the row index
// @param j: the column index
// @param nextLetter: the letter to search for
// @return bool: true if the next letter is to the north east of the current letter
func CheckNorthEast(matrix []string, i, j int, nextLetter string) bool {
	if i-1 >= 0 && j+1 < len(matrix[i]) && string(matrix[i-1][j+1]) == nextLetter {
		return true
	}
	return false
}

// CheckNorthWest checks if the next letter is to the north west of the current letter
// @param matrix: a 2D array of strings
// @param i: the row index
// @param j: the column index
// @param nextLetter: the letter to search for
// @return bool: true if the next letter is to the north west of the current letter
func CheckNorthWest(matrix []string, i, j int, nextLetter string) bool {
	if i-1 >= 0 && j-1 >= 0 && string(matrix[i-1][j-1]) == nextLetter {
		return true
	}
	return false
}

// CheckSouthEast checks if the next letter is to the south east of the current letter
// @param matrix: a 2D array of strings
// @param i: the row index
// @param j: the column index
// @param nextLetter: the letter to search for
// @return bool: true if the next letter is to the south east of the current letter
func CheckSouthEast(matrix []string, i, j int, nextLetter string) bool {
	if i+1 < len(matrix) && j+1 < len(matrix[i]) && string(matrix[i+1][j+1]) == nextLetter {
		return true
	}
	return false
}

// checkSouthWest checks if the next letter is to the south west of the current letter
// @param matrix: a 2D array of strings
// @param i: the row index
// @param j: the column index
// @param nextLetter: the letter to search for
// @return bool: true if the next letter is to the south west of the current letter
func checkSouthWest(matrix []string, i, j int, nextLetter string) bool {
	if i+1 < len(matrix) && j-1 >= 0 && string(matrix[i+1][j-1]) == nextLetter {
		return true
	}
	return false
}
