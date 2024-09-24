package main

import (
	"fmt"
	"tp"
)

func main() {
	coins := []int{1, 2, 5}
	amount := 11
	result := tp.Ft_coin(coins, amount)
	fmt.Println(result) // Affichera 3 (5 + 5 + 1 = 11)

	fmt.Println(tp.Ft_max_substring("abcabcbb")) // Résultat attendu : 3 ("abc")
	fmt.Println(tp.Ft_max_substring("bbbbb"))    // Résultat attendu : 1 ("b")
	fmt.Println(tp.Ft_max_substring("pwwkew"))   // Résultat attendu : 3 ("wke")

	fmt.Println(tp.Ft_min_window("ADOBECODEBANC", "ABC")) // Résultat attendu : "BANC"
	fmt.Println(tp.Ft_min_window("a", "aa"))              // Résultat attendu : ""

	fmt.Println(tp.FindMissingNumber([]int{3, 1, 2}))                   // Résultat attendu : 0
	fmt.Println(tp.FindMissingNumber([]int{0, 1}))                      // Résultat attendu : 2
	fmt.Println(tp.FindMissingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1})) // Résultat attendu : 8

	fmt.Println(tp.Ft_non_overlap([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}})) // Résultat attendu : 1
	fmt.Println(tp.Ft_non_overlap([][]int{{1, 2}, {2, 3}}))                 // Résultat attendu : 0
	fmt.Println(tp.Ft_non_overlap([][]int{{1, 2}, {1, 2}, {1, 2}}))         // Résultat attendu : 2

	fmt.Println(tp.Ft_profit([]int{7, 1, 5, 3, 6, 4})) // Résultat attendu : 5
	fmt.Println(tp.Ft_profit([]int{7, 6, 4, 3, 1}))    // Résultat attendu : 0
}
