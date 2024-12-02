package util

import "sort"

/**
 * Description:
 *
 * @author <anas.ahmed@zixel.cn>
 * @date 2024/12/01 2:16 pm
 */

func SortList(a *[]int, dir string) {
	list := *a
	if dir == "asc" {
		sort.Slice(list, func(i, j int) bool {
			return list[i] < list[j]
		})
	} else {
		sort.Slice(*a, func(i, j int) bool {
			return list[i] > list[j]
		})
	}

	*a = list
}

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
