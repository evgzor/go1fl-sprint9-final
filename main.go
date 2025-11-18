package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	if size < 1 {
		return []int{}
	}

	elements := make([]int, size)
	var randSource = rand.NewSource(time.Now().UnixNano())
	rng := rand.New(randSource)
	for i := range size {
		elements[i] = rng.Int()
	}

	return elements
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}

	result := data[0]
	for _, element := range data[1:] {
		if element > result {
			result = element
		}
	}
	return result
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}
	if len(data) == 1 {
		return data[0]
	}
	if len(data)%CHUNKS != 0 {
		return 0
	}
	maxElements := make([]int, CHUNKS)

	var wg sync.WaitGroup
	for i := range CHUNKS {
		wg.Add(1) // инкрементируем счётчик перед запуском горутины

		size := len(data)
		startIdx := i * size / CHUNKS
		endIdx := i*size/CHUNKS + size/CHUNKS
		go func(idx int, elements []int) {
			// уменьшаем счётчик, когда горутина завершает работу
			defer wg.Done()

			result := elements[0]
			for _, element := range elements[1:] {
				if element > result {
					result = element
				}
			}
			// захватили индекс ... и не нужна нам каналы и мютексы
			maxElements[idx] = result

		}(i, data[startIdx:endIdx])
	}
	wg.Wait()

	result := maxElements[0]
	for _, element := range maxElements[1:] {
		if element > result {
			result = element
		}
	}

	return result
}

func main() {
	fmt.Printf("Генерируем %d целых чисел", SIZE)
	elements := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	timeStart := time.Now()
	max := maximum(elements)
	timeEnd := time.Now()

	elapsed := timeEnd.Sub(timeStart).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков", CHUNKS)
	timeStart = time.Now()
	max = maxChunks(elements)
	timeEnd = time.Now()
	elapsed = timeEnd.Sub(timeStart).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
