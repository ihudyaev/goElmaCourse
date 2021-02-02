package main

import (
	"fmt"
)

const (
	zerobit = iota
	onebit
)

func bitCount(x int) int {
	var (
		i, prevBit, curBit, maxLen, curLen int
		str                                string
	)
	prevBit = -1
	maxLen = 0
	curLen = 1
	curBit = -1
	i = 0
	for x > 0 {
		curBit = x % 2

		if curBit == prevBit && prevBit >= 0 {
			curLen++
		}
		if curBit != prevBit && prevBit >= 0 {
			curLen = 1
		}
		if curLen >= maxLen {
			maxLen = curLen
		}

		if curBit == onebit {
			str = str + "1"

		}
		if curBit == zerobit {
			str = str + "0"
		}
		x = x / 2

		i++
		prevBit = curBit
	}
	fmt.Println(str)

	return maxLen

}

func main() {
	var l int
	// Вызов функции - принимает число в десятичном виде
	// возвращает максимальное количество последовательных двичных символов
	// печатает в консоль число в двоичном инверированном виде - для контроля
	l = bitCount(1253505)
	fmt.Println(l)

}
