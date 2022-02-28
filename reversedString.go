package main

func Solution(word string) string {
	str := ""
	for i := len(word) - 1; i >= 0; i-- {
		str += string(word[i])
	}
	return str
}
