package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"strconv"
	"time"
)

func square(x int) int {
	return x * x
}

func getRandomSquare() int {
	var x int

	if strconv.IntSize == 64 {
		x = rand.Intn(46340)
	}
	if strconv.IntSize == 32 {
		x = rand.Intn(303700499)
	}
	return square(x)
}

func getRandomSquareSequence(seqLen int) {

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < seqLen; i++ {
		fmt.Println(getRandomSquare())
	}

}

func main() {
	const PtrSize = 32 << uintptr(^uintptr(0)>>63)
	fmt.Println("System Info")
	fmt.Println(runtime.GOOS, runtime.GOARCH)
	fmt.Println(strconv.IntSize, PtrSize)
	fmt.Println("\nRandom Sequence")
	getRandomSquareSequence(10)

}
