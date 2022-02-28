package main

import "fmt"

func remove(slice []interface{}, idx int) []interface{} {
	return append(slice[:idx], slice[idx+1:]...)
}

func solve(arr []interface{}, answer []interface{}, sf int, num int) []interface{} {
	if len(arr) == 1 {
		answer = append(answer, arr[0])
		return answer
	} else {
		if sf+num < len(arr) {
			sf += num
			answer = append(answer, arr[sf])
			arr = remove(arr, sf)
			fmt.Println(answer)
		} else {
			sf = (sf + num) % len(arr)
			answer = append(answer, arr[sf])
			arr = remove(arr, sf)
		}
	}
	return solve(arr, answer, sf, num)
}

func Josephus(items []interface{}, k int) []interface{} {

	result := make([]interface{}, 0, len(items))

	if len(items) < 1 {
		return items
	} else {
		return solve(items, result, 0, k-1)

	}
}
