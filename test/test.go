package main

import (
	"fmt"
	"tp"
)

func main() {
	coins := []int{1, 2, 5}
	amount := 11
	result := tp.Ft_coin(coins, amount)
	fmt.Println(result)

	fmt.Println(tp.Ft_max_substring("abcabcbb"))
	fmt.Println(tp.Ft_max_substring("bbbbb"))
	fmt.Println(tp.Ft_max_substring("pwwkew"))

	fmt.Println(tp.Ft_min_window("ADOBECODEBANC", "ABC"))
	fmt.Println(tp.Ft_min_window("a", "aa"))

	fmt.Println(tp.FindMissingNumber([]int{3, 1, 2}))
	fmt.Println(tp.FindMissingNumber([]int{0, 1}))
	fmt.Println(tp.FindMissingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}))

	fmt.Println(tp.Ft_non_overlap([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}))
	fmt.Println(tp.Ft_non_overlap([][]int{{1, 2}, {2, 3}}))
	fmt.Println(tp.Ft_non_overlap([][]int{{1, 2}, {1, 2}, {1, 2}}))

	fmt.Println(tp.Ft_profit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(tp.Ft_profit([]int{7, 6, 4, 3, 1}))
}
