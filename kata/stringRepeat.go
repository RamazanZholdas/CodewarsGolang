package main

func RepeatStr(repetitions int, value string) string {
	str := value
	for repetitions > 1 {
		value += str
		repetitions--
	}
	return value
}
