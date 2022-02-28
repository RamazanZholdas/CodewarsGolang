package main

func CountSheeps(numbers []bool) int {
	num := 0
	for _, v := range numbers {
		if v {
			num++
		}
	}
	return num
}
