package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	numChan := make(chan int)
	avgChan := make(chan float64)

	go generateRandomNumbers(numChan)
	go calculateAverage(numChan, avgChan)
	go printAverage(avgChan)

	time.Sleep(10 * time.Second)
}

func generateRandomNumbers(intCh chan int) {
	for {
		num := rand.Intn(100)
		fmt.Printf("%d, ", num)
		intCh <- num
		time.Sleep(2 * time.Second)
	}
}

func calculateAverage(input chan int, output chan float64) {
	var numbers []int
	for num := range input {
		numbers = append(numbers, num)
		fmt.Printf("Random numbers: %d,", numbers)
		total := 0
		for _, n := range numbers {
			total += n
		}
		average := float64(total) / float64(len(numbers))
		output <- average
	}
}

func printAverage(output chan float64) {
	for avg := range output {
		fmt.Printf(" Average: %.2f\n", avg)
	}
}
