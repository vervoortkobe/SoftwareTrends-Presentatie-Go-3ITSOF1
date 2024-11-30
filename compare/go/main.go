package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func counter(amount int) int {
	sum := 0
	for i := 0; i < amount; i++ {
		sum++
	}
	return sum
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	prev, curr := 0, 1
	for i := 2; i <= n; i++ {
		prev, curr = curr, prev+curr
	}
	return curr
}

func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[len(arr)-1]
	left := make([]int, 0)
	right := make([]int, 0)

	for i := 0; i < len(arr)-1; i++ {
		if arr[i] < pivot {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}

	left = quickSort(left)
	right = quickSort(right)

	return append(append(left, pivot), right...)
}

func bubbleSort(arr []int) []int {
	n := len(arr)
	result := make([]int, len(arr))
	copy(result, arr)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if result[j] > result[j+1] {
				result[j], result[j+1] = result[j+1], result[j]
			}
		}
	}
	return result
}

func generateRandomArray(size int) []int {
	arr := make([]int, size)
	for i := range arr {
		arr[i] = rand.Intn(10000)
	}
	return arr
}

func printArraySample(arr []int, name string) string {
	if len(arr) <= 10 {
		return fmt.Sprintf("%s: %v", name, arr)
	}
	return fmt.Sprintf("%s: [%v, ..., %v]", name, arr[:5], arr[len(arr)-5:])
}

func padRight(str string, length int) string {
	if len(str) >= length {
		return str
	}
	return str + strings.Repeat(" ", length-len(str))
}

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Go Benchmark Results:")
	fmt.Println("-------------------")

	// Counter benchmark
	start := time.Now()
	sum := counter(100000)
	duration := time.Since(start).Nanoseconds()
	fmt.Printf("%s %10.6f ms\n", padRight(fmt.Sprintf("1. Counter: %d", sum), 70), float64(duration)/1000000)

	// Fibonacci benchmark
	start = time.Now()
	fibResult := fibonacci(100)
	duration = time.Since(start).Nanoseconds()
	fmt.Printf("%s %10.6f ms\n", padRight(fmt.Sprintf("2. Fibonacci: %d", fibResult), 70), float64(duration)/1000000)

	// QuickSort benchmark
	quickSortArr := generateRandomArray(1000)
	inputStr := printArraySample(quickSortArr, "   - Input")
	fmt.Println("3. Quicksort:")
	fmt.Println(inputStr)

	start = time.Now()
	sortedQuick := quickSort(quickSortArr)
	duration = time.Since(start).Nanoseconds()

	outputStr := printArraySample(sortedQuick, "   - Output")
	fmt.Printf("%s %10.6f ms\n", padRight(outputStr, 70), float64(duration)/1000000)

	// BubbleSort benchmark
	bubbleSortArr := generateRandomArray(1000)
	inputStr = printArraySample(bubbleSortArr, "   - Input")
	fmt.Println("4. BubbleSort:")
	fmt.Println(inputStr)

	start = time.Now()
	sortedBubble := bubbleSort(bubbleSortArr)
	duration = time.Since(start).Nanoseconds()

	outputStr = printArraySample(sortedBubble, "   - Output")
	fmt.Printf("%s %10.6f ms\n", padRight(outputStr, 70), float64(duration)/1000000)
}
