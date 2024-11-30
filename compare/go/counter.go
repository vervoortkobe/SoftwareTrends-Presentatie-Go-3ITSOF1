package main

func Counter(amount int) int {
	sum := 0
	for i := 0; i < amount; i++ {
		sum++
	}
	return sum
}
