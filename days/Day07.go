package days

import (
	"AdventOfCode2024/util"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

/**
 * Description:
 *
 * @author <anas.ahmed@zixel.cn>
 * @date 2024/12/07 10:00 am
 */

type NumAndOperands struct {
	Num      int
	Operands []int
	True     bool
}

func Day07() {
	fmt.Println("Day 07:")
	day7part1()
	day7part2()
}

func day7part1() (rules []string, invalidOrders [][]int) {
	numList := day7common()

	operators := []string{"+", "*"}
	sum := 0

	// Try each combination of operators on operands, if the result equals the number, add it to the sum
	wg := sync.WaitGroup{}
	for c := range numList {
		wg.Add(1)
		go func(numEntry *NumAndOperands) {
			defer wg.Done()
			combos := util.GetCombos(operators, len(numEntry.Operands)-1)
			for _, combo := range combos {
				eval := numEntry.Operands[0]
				for i, operator := range combo {
					num := numEntry.Operands[i+1]

					switch operator {
					case "+":
						eval += num
					case "*":
						eval *= num
					}
				}

				if eval == numEntry.Num {
					numEntry.True = true
					break
				}
			}
		}(&numList[c])
	}
	wg.Wait()

	// Cal sum
	for _, numEntry := range numList {
		if numEntry.True {
			sum += numEntry.Num
		}
	}

	// Print the sum
	fmt.Println("Part 1:")
	fmt.Println(sum)

	return
}

func day7part2() {
	numList := day7common()

	operators := []string{"+", "*", "|"}
	sum := 0

	wg := sync.WaitGroup{}
	// Try each combination of operators on operands, if the result equals the number, add it to the sum
	for c := range numList {
		wg.Add(1)
		go func(numEntry *NumAndOperands) {
			defer wg.Done()
			combos := util.GetCombos(operators, len(numEntry.Operands)-1)
			for _, combo := range combos {

				eval := numEntry.Operands[0]
				for i, operator := range combo {
					num := numEntry.Operands[i+1]

					switch operator {
					case "+":
						eval += num
					case "*":
						eval *= num
					case "|":
						evalString := fmt.Sprintf("%d%d", eval, num)
						eval, _ = strconv.Atoi(evalString)
					}
				}

				if eval == numEntry.Num {
					numEntry.True = true
					break
				}
			}
		}(&numList[c])
	}
	wg.Wait()

	// Cal sum
	for _, numEntry := range numList {
		if numEntry.True {
			sum += numEntry.Num
		}
	}

	// Print the sum
	fmt.Println("Part 2:")
	fmt.Println(sum)
}

func day7common() (numList []NumAndOperands) {
	// Read txt file from ../inputs/day01.txt
	file, err := os.Open("inputs/Day07.EXAMPLE")
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
		colonIndex := strings.Index(line, ":")
		numStr := line[:colonIndex]
		num, _ := strconv.Atoi(numStr)
		numEntry := NumAndOperands{Num: num}

		operands := strings.Split(line[colonIndex+2:], " ")
		for _, operand := range operands {
			num, _ = strconv.Atoi(operand)
			numEntry.Operands = append(numEntry.Operands, num)
		}

		numList = append(numList, numEntry)
	}

	return
}
