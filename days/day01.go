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
 * @date 2024/12/01 2:03 pm
 */

func Day01() {
	fmt.Println("Day 01:")
	day1part1()
	day1part2()
}

func day1part1() {
	leftList, rightList := day1common()

	util.SortList(&leftList, "asc")
	util.SortList(&rightList, "asc")

	// Find difference between the two lists and sum
	sum := 0
	for i := 0; i < len(leftList); i++ {
		sum += util.Abs(leftList[i] - rightList[i])
	}

	// Print the sum
	fmt.Println("Part 1:")
	fmt.Println(sum)
}

func day1part2() {
	leftList, rightList := day1common()

	sum := 0
	score := 0
	for i := 0; i < len(leftList); i++ {
		for j := 0; j < len(rightList); j++ {
			if leftList[i] == rightList[j] {
				score += 1
			}
		}
		sum += score * leftList[i]
		score = 0
	}

	// Print the sum
	fmt.Println("Part 2:")
	fmt.Println(sum)
}

func day1common() (leftList []int, rightList []int) {
	// Read txt file from ../inputs/day01.txt
	file, err := os.Open("inputs/Day01.EXAMPLE")
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
		lineMid := strings.Index(line, "-")
		if lineMid != -1 {
			leftNumber, _ := strconv.Atoi(line[:lineMid])
			rightNumber, _ := strconv.Atoi(line[lineMid+1:])

			if leftNumber <= 0 || rightNumber <= 0 {
				fmt.Printf("Invalid input on line:%s, left:%d, right:%d", line, leftNumber, rightNumber)
				continue
			}

			leftList = append(leftList, leftNumber)
			rightList = append(rightList, rightNumber)
		} else {
			fmt.Println("Invalid input found!")
		}
	}

	return leftList, rightList
}
