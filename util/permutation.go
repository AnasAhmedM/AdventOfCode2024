package util

/**
 * Description:
 *
 * @author <anas.ahmed@zixel.cn>
 * @date 2024/12/07 10:30 am
 */

// GetCombos Recursive function to generate all possible combinations of operators
// Possible values of n are passed in nList, and r is the number of values to combine
// Returns a list of all possible combinations
func GetCombos(nList []string, r int) [][]string {
	if r == 1 {
		// Return each n as its own combination (base case)
		var combos [][]string
		for _, n := range nList {
			combos = append(combos, []string{n})
		}
		return combos
	}

	// Recursive step: Get combinations for smaller count
	subCombos := GetCombos(nList, r-1)
	var combos [][]string

	// Append each n to all sub-combinations
	for _, n := range nList {
		for _, subCombo := range subCombos {
			newCombo := append([]string{n}, subCombo...)
			combos = append(combos, newCombo)
		}
	}

	return combos
}
