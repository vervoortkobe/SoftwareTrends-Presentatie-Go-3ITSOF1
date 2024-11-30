package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func GenerateRandomArray(size int) []int {
	arr := make([]int, size)
	for i := range arr {
		arr[i] = rand.Intn(100_000_000)
	}
	return arr
}

func PrintArraySample(arr []int, name string) string {
	if len(arr) <= 10 {
		return fmt.Sprintf("%s: %v", name, arr)
	}
	return fmt.Sprintf("%s: [%v, ..., %v]", name, arr[:5], arr[len(arr)-5:])
}

func PadRight(str string, length int) string {
	if len(str) >= length {
		return str
	}
	return str + strings.Repeat(" ", length-len(str))
}
