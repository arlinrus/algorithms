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
	insertionArr(emptyArr)
	duration := time.Since(start)
	fmt.Println("Время выполнения O(n^2): ", duration)

}

func createArr() (arr []int) {
	sizeArr := 1000
	arr = make([]int, sizeArr)
	for i := 0; i < sizeArr; i++ {
		arr[i] = rand.Intn(sizeArr)

	}
	return
}
func insertionArr(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--

		}
		arr[j+1] = key
	}
	return arr
}
