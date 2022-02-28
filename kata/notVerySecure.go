package main

import "fmt"

func alphanumeric(str string) bool {
	for i := range str {
		if !((str[i] <= 90 && str[i] >= 65) || (str[i] >= 97 && str[i] <= 122) || (str[i] >= 48 && str[i] <= 57)) {
			fmt.Println(str[i], str[i])
			return false
		}
	}
	if str == "" {
		return false
	}
	return true
}
