package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Go Benchmark Results:")
	fmt.Println("-------------------")

	// Counter benchmark
	start := time.Now()
	sum := Counter(1_000_000)
	duration := time.Since(start).Nanoseconds()
	fmt.Printf("%s %10.6f ms\n", PadRight(fmt.Sprintf("1. Counter: %d", sum), 70), float64(duration)/1000000)

	// Fibonacci benchmark
	start = time.Now()
	fibResult := Fibonacci(1_000)
	duration = time.Since(start).Nanoseconds()
	fmt.Printf("%s %10.6f ms\n", PadRight(fmt.Sprintf("2. Fibonacci: 1000 - %d", fibResult), 70), float64(duration)/1000000)

	// QuickSort benchmark
	quickSortArr := GenerateRandomArray(10_000)
	inputStr := PrintArraySample(quickSortArr, "   - Input")
	fmt.Println("3. Quicksort:")
	fmt.Println(inputStr)

	start = time.Now()
	sortedQuick := QuickSort(quickSortArr)
	duration = time.Since(start).Nanoseconds()

	outputStr := PrintArraySample(sortedQuick, "   - Output")
	fmt.Printf("%s %10.6f ms\n", PadRight(outputStr, 70), float64(duration)/1000000)

	// BubbleSort benchmark
	bubbleSortArr := GenerateRandomArray(10_000)
	inputStr = PrintArraySample(bubbleSortArr, "   - Input")
	fmt.Println("4. BubbleSort:")
	fmt.Println(inputStr)

	start = time.Now()
	sortedBubble := BubbleSort(bubbleSortArr)
	duration = time.Since(start).Nanoseconds()

	outputStr = PrintArraySample(sortedBubble, "   - Output")
	fmt.Printf("%s %10.6f ms\n", PadRight(outputStr, 70), float64(duration)/1000000)
}
