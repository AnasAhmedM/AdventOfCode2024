package days

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

/**
 * Description:
 *
 * @author <anas.ahmed@zixel.cn>
 * @date 2024/12/03 10:05 am
 */

func Day03() {
	fmt.Println("Day 03:")
	day3part1()
	day3part2()
}

func day3part1() {
	txt := day3common("01")

	results := regexp.MustCompile(`(mul\([0-9]{1,3},[0-9]{1,3}\))`).FindAllString(txt, -1)
	var inputs []string
	for _, input := range results {
		inputs = append(inputs, input[4:len(input)-1])
	}

	sum := 0
	for _, input := range inputs {
		nums := strings.Split(input, ",")
		num1, num2 := 0, 0
		num1, _ = strconv.Atoi(nums[0])
		num2, _ = strconv.Atoi(nums[1])
		sum += num1 * num2
	}

	fmt.Println("Part 1:")
	fmt.Println(sum)
}

func day3part2() {
	txt := day3common("02")

	results := regexp.MustCompile(`(mul\([0-9]{1,3},[0-9]{1,3}\))`).FindAllString(txt, -1)

	var inputs []string
	for _, input := range results {
		inputIndex := strings.Index(txt, input)
		dontIndex := strings.LastIndex(txt[:inputIndex], "don't()")
		doIndex := strings.LastIndex(txt[:inputIndex], "do()")

		if doIndex > dontIndex || dontIndex == -1 {
			results = regexp.MustCompile(`([0-9]{1,3},[0-9]{1,3})`).FindAllString(input, 1)
			inputs = append(inputs, results[0])
			if doIndex != -1 {
				txt = txt[doIndex:]
			}
		}
	}

	sum := 0
	for _, input := range inputs {
		nums := strings.Split(input, ",")
		num1, num2 := 0, 0
		num1, _ = strconv.Atoi(nums[0])
		num2, _ = strconv.Atoi(nums[1])
		sum += num1 * num2
	}

	fmt.Println("Part 2:")
	fmt.Println(sum)
}

func day3common(version string) (txt string) {
	//	// Read txt file
	file, err := os.Open("inputs/Day03.EXAMPLE" + version)
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
		txt += scanner.Text()
	}

	return
}
