package main

import (
	"errors"
	"fmt"
)

// --- 複数の戻り値 ---
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("ゼロ除算できません")
	}
	return a / b, nil
}

// --- 名前付き戻り値 ---
func swap(a, b int) (x, y int) {
	x = b
	y = a
	return // x, yの現在値がそのまま返される(naked return)
}

// --- 可変長引数 ---
func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// --- 高階関数 ---
func apply(nums []int, f func(int) int) []int {
	result := make([]int, len(nums))
	for i, n := range nums {
		result[i] = f(n)
	}
	return result
}

func main() {
	fmt.Println("=== 複数の戻り値 ===")
	result, err := divide(10, 3)
	if err != nil {
		fmt.Println("エラー：", err)
	} else {
		fmt.Printf("10 / 3 = %.2f\n", result)
	}

	_, err = divide(10, 0)
	if err != nil {
		fmt.Println("エラー：", err)
	}

	fmt.Println("\n=== 名前付き戻り値 ===")
	a, b := swap(1, 2)
	fmt.Println(a, b)

	fmt.Println("\n=== 可変長引数 ===")
	fmt.Println(sum(1, 2, 3))
	fmt.Println(sum(10, 20, 30, 40, 50))
	numbers := []int{100, 200, 300}
	fmt.Println(sum(numbers...))

	fmt.Println("\n=== 高階関数 ===")
	nums := []int{1, 2, 3, 4, 5}
	doubled := apply(nums, func(n int) int { return n * 2 })
	fmt.Println("doubled:", doubled)
	squared := apply(nums, func(n int) int { return n * n })
	fmt.Println("squared:", squared)
}
