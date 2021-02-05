// rotate array
// right or left direction for N times
package main

import "fmt"

//direction:
// true - right
// false - left
func rotate(numbers []int, n int, direction bool) []int {
	//Если длина массива меньше количества вращений, то корректируем n (сокращаем количество полных оборотов массива), чтобы не тратить ресурсы
	var nr int
	nr = n % len(numbers)
	//если массив пустой или из одного числа или количество вращений ноль, то возвращаем исходный массив
	if nr < 1 || len(numbers) <= 1 {
		return numbers
	}

	//вращение вправо
	if direction == true {
		r := len(numbers) - nr
		numbers = append(numbers[r:], numbers[:r]...)
	}

	//вращение влево
	if direction == false {
		numbers = append(numbers[nr:], numbers[:nr]...)
	}

	return numbers
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	fmt.Printf("Before rotate %v\n", nums)

	var (
		leftRightType bool
		rotateNum     int
		rotateType    string
	)

	rotateNum = 123
	leftRightType = true
	if leftRightType == true {
		rotateType = "right"
	} else {
		rotateType = "left"
	}

	nums = rotate(nums, rotateNum, leftRightType)

	fmt.Printf("After %s rotate  %d times  %v\n", rotateType, rotateNum, nums)
}
