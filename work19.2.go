/*
Отсортируйте массив длиной шесть пузырьком.
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

const size = 6

func generateArray() (array [size]int) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		array[i] = rand.Intn(20)
	}
	return
}

func bubble(array [size]int) [size]int {
	for i := 0; i < size; i++ {
		for j := 1; j < size-i; j++ {
			if array[j-1] > array[j] {
				array[j], array[j-1] = array[j-1], array[j]
			}
		}
	}
	return array
}

func main() {
	array := generateArray()
	fmt.Println("Сгенерированный массив: \n", array)
	array = bubble(array)
	fmt.Println("После сортировки пузырьком:\n", array)

}
