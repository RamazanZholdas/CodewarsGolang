package main

func SmallestIntegerFinder(numbers []int) int {
	var temp int
	for i, v := range numbers {
		if i == 0 {
			temp = v
			continue
		} else if v < temp {
			temp = v
		}
	}
	return temp
}
