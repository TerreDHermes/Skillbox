/*
Напишите функцию, которая производит слияние двух отсортированных массивов длиной четыре и пять в один массив длиной девять.
*/
package main

import (
	"fmt"
)

func mergeArray(array1 [4]int, array2 [5]int) (array [9]int) {
	for i := 0; i < 9; i++ {
		if i < 4 {
			array[i] = array1[i]
		} else {
			array[i] = array2[i-4]
		}
	}
	return array
}

func main() {
	array1 := [4]int{1, 2, 3, 4}
	array2 := [5]int{5, 6, 7, 8, 9}
	array := mergeArray(array1, array2)
	fmt.Println(array)
}
