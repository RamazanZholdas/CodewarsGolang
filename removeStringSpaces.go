package main

func NoSpace(word string) string {
	res := ""
	for _, v := range word {
		if string(v) == " " {
			continue
		}
		res += string(v)
	}
	return res
}
