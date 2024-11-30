package main

import "math/big"

func Fibonacci(n int) *big.Int {
	if n <= 1 {
		return big.NewInt(int64(n))
	}
	prev, curr := big.NewInt(0), big.NewInt(1)
	for i := 2; i <= n; i++ {
		prev, curr = curr, new(big.Int).Set(prev).Add(prev, curr)
	}
	return curr
}
