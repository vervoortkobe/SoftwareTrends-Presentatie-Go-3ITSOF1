package main

import (
	"fmt"
)

// Returnt max value in array van type T
func Max[T any](items []T, compare func(a, b T) bool) T {
	if len(items) == 0 {
		var zeroValue T
		return zeroValue // Returnt zero value van type T (0)
	}

	max := items[0]
	for _, item := range items {
		if compare(item, max) {
			max = item
		}
	}
	return max
}

func Generics() {
	fmt.Println("------------- Generics -------------")

	// Integers
	integers := []int{1, 3, 5, 2, 4}
	maxInt := Max(integers, func(a, b int) bool { return a > b })
	fmt.Printf("Max int: %d\n", maxInt)

	// Strings
	strings := []string{"kat", "hond", "aap", "papegaai"}
	maxStr := Max(strings, func(a, b string) bool { return a > b })
	fmt.Printf("Max string: %s\n", maxStr)
}
