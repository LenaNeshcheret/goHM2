package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	numbersChannel := make(chan []int)
	minMaxChannel := make(chan [2]int)

	go generateNumbers(5, 300, 10, minMaxChannel, numbersChannel)
	go findMixMax(numbersChannel, minMaxChannel)
	time.Sleep(10 * time.Second)
}

func generateNumbers(a int, b int, amount int, input chan [2]int, output chan []int) {
	fmt.Printf("Generate %d numbers between [ %d,  %d ] \n", amount, a, b)
	numbers := make([]int, amount)

	for i := 0; i < amount; i++ {
		numbers[i] = rand.Intn(b-a+1) + a
		fmt.Printf("%d, ", numbers[i])
	}
	output <- numbers
	defer close(output)
	for result := range input {
		fmt.Printf("\nMAX = %d, \nMIN = %d \n", result[1], result[0])
	}

}
func findMixMax(input chan []int, output chan [2]int) {
	numbers := <-input
	minVal, maxVal := numbers[0], numbers[1]

	for _, num := range numbers {
		if num < minVal {
			minVal = num
		}
		if num > maxVal {
			maxVal = num
		}
	}
	output <- [2]int{minVal, maxVal}
	close(output)
}
