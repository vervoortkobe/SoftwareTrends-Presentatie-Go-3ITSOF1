package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
)

func printNumbers(prefix string, count int, wg *sync.WaitGroup) {
	color := Green
	if wg != nil {
		color = Blue
		defer wg.Done() // Signal task completion
	}

	for i := 1; i <= count; i++ {
		fmt.Printf("%s%s: %d%s\n", color, prefix, i, Reset)
		time.Sleep(100 * time.Millisecond)
	}
}

func Goroutines() {
	fmt.Println("------------ Goroutines ------------")

	var wg sync.WaitGroup

	wg.Add(1) // Voeg task toe aan WaitGroup
	go printNumbers("Goroutine", 5, &wg)

	printNumbers("Main", 5, nil)

	wg.Wait() // Wacht op alle goroutines completed zijn
	fmt.Println("Done!")
}
