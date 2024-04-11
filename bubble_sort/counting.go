package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	arr := createArr()
	fmt.Println(arr[:10])

	emptyArr := make([]int, len(arr))
	copy(emptyArr, arr)
	start := time.Now()
	countSort(emptyArr)
	duration := time.Since(start)
	fmt.Println("Время выполнения O(n^2): ", duration)

}

func createArr() (arr []int) {
	sizeArr := 11
	arr = make([]int, sizeArr)
	for i := 0; i < sizeArr; i++ {
		arr[i] = rand.Intn(sizeArr)

	}
	return
}

func countSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		minIndex := 1
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}

		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
	return arr

}
