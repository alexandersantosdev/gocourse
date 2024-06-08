package main

import "fmt"

func main() {

	n := 45

	tasks := make(chan int, n)
	results := make(chan int, n)

	go worker(tasks, results)
	go worker(tasks, results)
	go worker(tasks, results)
	go worker(tasks, results)
	go worker(tasks, results)
	go worker(tasks, results)
	go worker(tasks, results)
	go worker(tasks, results)

	for i := 0; i <= n; i++ {
		tasks <- i
	}
	close(tasks)

	for i := 0; i <= n; i++ {
		res := <-results
		fmt.Println(res)
	}
}

// task chan receive data, result chan send data
func worker(tasks <-chan int, results chan<- int) {
	for t := range tasks {
		results <- fibonacci(t)
	}
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-2) + fibonacci(n-1)
}
