package main

import "fmt"

func CreatePhoneNumber(va [10]uint) string {
	var n []string
	for _, v := range va {
		n = append(n, fmt.Sprint(v))
	}
	return "(" + n[0] + n[1] + n[2] + ") " + n[3] + n[4] + n[5] + "-" + n[6] + n[7] + n[8] + n[9]
}
