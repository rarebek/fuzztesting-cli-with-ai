package main

import (
	"fmt"
	"sync"
)

func processData(data int, wg *sync.WaitGroup, results chan<- int) {
	defer wg.Done()

	result := data * 2
	results <- result
}

func main() {
	var wg sync.WaitGroup

	dataSets := []int{1, 2, 3, 4, 5}
	results := make(chan int, len(dataSets))

	for _, data := range dataSets {
		wg.Add(1)
		go processData(data, &wg, results)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Println(result)
	}
}
