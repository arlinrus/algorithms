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
	Bubblesort(emptyArr)
	duration := time.Since(start)
	fmt.Println("Время выполнения O(n^2): ", duration)

}

func createArr() (arr []int) {
	sizeArr := 15
	arr := make([]int, sizeArr)

	for i := 0; i < sizeArr; i++ {
		arr[i] = rand.Intn(sizeArr)

	}
	return
}
