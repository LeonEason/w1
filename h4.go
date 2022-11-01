package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var slice []int
	var i int
	slice = make([]int, 0)
	rand.Seed(time.Now().UnixNano())
	for i = 1; i <= 100; i++ {
		slice = append(slice, rand.Intn(100)+1)
	}
	slice = sliceSort(slice)
	fmt.Println(slice)

}
func sliceSort(slice []int) []int { //倒序
	l := len(slice)
	for i := 0; i < l; i++ {
		for j := 0; j < l-i-1; j++ {
			if slice[j] < slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
	return slice
}
