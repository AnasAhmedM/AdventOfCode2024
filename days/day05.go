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
 * @date 2024/12/05 10:10 AM
 */

func Day05() {
	fmt.Println("Day 05:")
	day5part2()
}

func day5part1() (rules []string, invalidOrders [][]int) {
	rules, orders := day5common()

	total := 0
	for _, order := range orders {
		validOrder := true
		for i := 0; i < len(order)-1; i++ {
			for j := i + 1; j < len(order); j++ {
				if !util.InSliceString(rules, fmt.Sprintf("%d|%d", order[i], order[j])) {
					validOrder = false
					break
				}
			}
		}
		if validOrder {
			midIndex := len(order) / 2
			total += order[midIndex]
		} else {
			invalidOrders = append(invalidOrders, order)
		}
	}

	// Print the sum
	fmt.Println("Part 1:")
	fmt.Println(total)

	return
}

func day5part2() {
	rules, invalidOrders := day5part1()

	for _, order := range invalidOrders {
		for i := 0; i < len(order)-1; i++ {
			for j := i + 1; j < len(order); j++ {
				if !util.InSliceString(rules, fmt.Sprintf("%d|%d", order[i], order[j])) {
					// Swap the two values
					order[i], order[j] = order[j], order[i]
				}
			}
		}
	}

	total := 0
	for _, order := range invalidOrders {
		midIndex := len(order) / 2
		total += order[midIndex]
	}

	// Print the sum
	fmt.Println("Part 2:")
	fmt.Println(total)
}

func day5common() (rules []string, orders [][]int) {
	// Read txt file from ../inputs/day01.txt
	file, err := os.Open("inputs/Day05.EXAMPLE")
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

	ordersStart := false
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			ordersStart = true
			continue
		}

		if !ordersStart {
			rules = append(rules, line)
		} else {
			var order []int
			ordersStr := strings.Split(line, ",")
			for _, i := range ordersStr {
				o, _ := strconv.Atoi(i)
				order = append(order, o)
			}
			orders = append(orders, order)
		}
	}

	return
}
