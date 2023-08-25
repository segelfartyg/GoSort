// Module for sorting arrays with established sorting algorithms
package main

import "fmt"

var unsortedList = []int{4, 6, 2, 21, 13, 17, 3, 7, 1, 22, 55}

func main() {
	fmt.Println(unsortedList)
	fmt.Println(Merge(unsortedList))
}

func Bubble(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if arr[i] < arr[j] {
				temp := arr[j]
				arr[j] = arr[i]
				arr[i] = temp
			}
		}
	}
	return arr
}

func Insertion(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		key := i
		for key > 0 {
			if arr[key] < arr[key-1] {
				temp := arr[key]
				arr[key] = arr[key-1]
				arr[key-1] = temp
			}
			key--
		}
	}
	return arr
}

func Selection(arr []int) []int {

	for i := 0; i < len(arr) - 1; i++ {
		minIndex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
			minIndex = j
			}
		}

		if(minIndex != i) {
			temp := arr[i]
			arr[i] = arr[minIndex]
			arr[minIndex] = temp
		}
	
	}
	return arr
}

func Merge(arr []int) []int {

	if len(arr) <= 1 {
		return arr
	}
	
	middle := len(arr) / 2

	FirstHalf := Merge(arr[0:middle])
	SecondHalf := Merge(arr[middle:])

	return merge(FirstHalf, SecondHalf)
}

func merge(left []int, right []int) []int {

	result := make([]int, 0, len(left)+len(right))
	l, r := 0, 0

	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}

	result = append(result, left[l:]...)
	result = append(result, right[r:]...)

	return result

}

func split(splitArr []int) []int {
	return splitArr
}

func Quick(arr []int) []int {
	return arr
}

func Heap(arr [] int) []int {
	return arr
}
