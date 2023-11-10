package main

import "fmt"

func subsets(nums []int) [][]int {
	result := [][]int{}
	var backtrack func(int, []int)
	backtrack = func(first int, curr []int) {
		temp := make([]int, len(curr))
		copy(temp, curr)
		result = append(result, temp)

		for i := first; i < len(nums); i++ {
			curr = append(curr, nums[i])
			backtrack(i+1, curr)
			curr = curr[:len(curr)-1]
		}
	}

	backtrack(0, []int{})
	return result
}

func main() {
	test := []int{1, 3, 2, 4}
	fmt.Println("Subconjuntos do conjunto {1, 3, 2, 4}:")
	subconjuntos := subsets(test)
	fmt.Println(subconjuntos)

	fmt.Println("Subconjuntos do conjunto vazio {}:")
	subconjuntosEmpty := subsets([]int{})

	fmt.Println(subconjuntosEmpty)

	fmt.Println("Subconjuntos do conjunto {1}:")
	subconjuntosOfConjunto := subsets([]int{1})
	fmt.Println(subconjuntosOfConjunto)
}
