package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	valuesChan := make(chan int)
	powValuesChan := make(chan int)

	go generateIntValues(valuesChan)
	go powIntValues(valuesChan, powValuesChan)

	for powValue := range powValuesChan {
		fmt.Println(powValue)
	}
}

// generateIntValues - Генерирует случайные значения и отправляет в канал resultChan
func generateIntValues(resultChan chan<- int) {
	defer close(resultChan)
	for i := 0; i < 10; i++ {
		value := rand.Intn(100)
		resultChan <- value
	}
}

// powIntValues - Возводит в квадрат значения полученные через канал valuesChan и отправляет результат в resultChan
func powIntValues(valuesChan <-chan int, resultChan chan<- int) {
	defer close(resultChan)
	for value := range valuesChan {
		resultChan <- int(math.Pow(float64(value), 2))
	}
}
