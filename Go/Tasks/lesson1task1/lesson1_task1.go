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
	//ограничиваем максимальное значение, чтобы квадрат числа не вышел за пределы выбранных типов
	// можно также использовать приведение типов и uint чтобы увеличить максимально допустимые значения
	// также максимальное значение зависит от разрядности системы
	if strconv.IntSize == 64 {
		x = rand.Intn(46340)
	}
	if strconv.IntSize == 32 {
		x = rand.Intn(303700499)
	}
	return square(x)
}

func getRandomSquareSequence(seqLen int) {
	// функция для получения последовательности

	//настраиваем радномизатор, чтобы не получить псевдо случайную последовательность
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < seqLen; i++ {
		fmt.Println(getRandomSquare())
	}

}

func main() {
	//Инфо о системе на которой компилируем программу - для наглядности и контроля
	const PtrSize = 32 << uintptr(^uintptr(0)>>63)
	fmt.Println("System Info")
	fmt.Println(runtime.GOOS, runtime.GOARCH)
	fmt.Println(strconv.IntSize, PtrSize)
	//Выполнение функции - получение последовательности квадратов случайных чисел
	fmt.Println("\nRandom Sequence")
	getRandomSquareSequence(10)

}
