package main

import "fmt"

func main() {

	fmt.Println(fatorial(5))

	fmt.Println(fibonacci(10))

	posicao := 10

	for i := 1; i < posicao; i++ {
		fmt.Println(fibonacci(i))
	}

}

func fatorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * fatorial(n-1)
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
