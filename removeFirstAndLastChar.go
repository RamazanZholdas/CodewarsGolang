package main

func RemoveChar(word string) string {
	str := ""
	for i := 1; i <= len(word)-2; i++ {
		str += string(word[i])
	}
	return str
}
