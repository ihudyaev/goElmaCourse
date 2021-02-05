// Get distinct numbers and distinct count from Array
package main

import "fmt"

func getUnique(nums []int) (unique []int, unqCount int) {
	//unsing map to check unique elements
	m := map[int]bool{}

	for _, n := range nums {
		if !m[n] {
			m[n] = true
			unique = append(unique, n)
		}
	}
	unqCount = len(unique)
	return
}

func main() {
	nums := []int{1, 1, 4, 4, 1, 9, 7, 7, 7, 0}
	var (
		unq       []int
		uniqueCnt int
	)

	unq, uniqueCnt = getUnique(nums)

	fmt.Printf("Source %v\nUnique  %v\nUnique elements count = %d\n", nums, unq, uniqueCnt)
}
