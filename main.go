package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	iterations := 100000
	ticks := 10
	initial_slope := 2.00
	slope_delta := 2.00

	var wg sync.WaitGroup
	resultsChan := make(chan struct {
		index int
		data  []float64
	}, iterations)

	for i := 1; i <= iterations; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			fmt.Println("Running simulation number:", i)

			slope := initial_slope + slope_delta*float64(i)
			var value_container []float64

			for j := 0; j < ticks; j++ {
				y := math.Round(float64(j) * slope)
				value_container = append(value_container, y)
			}

			resultsChan <- struct {
				index int
				data  []float64
			}{index: i, data: value_container}
		}(i)
	}

	wg.Wait()
	close(resultsChan)

	// Collect results
	finalResults := make(map[int][]float64)
	for result := range resultsChan {
		finalResults[result.index] = result.data
	}

	fmt.Println(finalResults)
	duration := time.Since(start) // 🧾 Elapsed time
	fmt.Printf("Completed in %s\n", duration)

}
