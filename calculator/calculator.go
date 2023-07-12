package calculator

import "fmt"

func Sum(a, b int64) int64 {
	return a + b
}

func Sub(a, b int64) int64 {
	return a - b
}

func Multipy(a, b int64) int64 {
	return a * b
}

func Division(a, b int64) int64 {
	return a / b
}

func DataStructureWithLoop() {
	tables := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	for k, v := range tables {
		fmt.Printf("Index %v \t Value %v", k, v)
	}
}

func DataStructureWithMap() {
	personInfo := map[string]int{
		"Raghavendra": 30,
		"PIN Code":    582101,
	}
	for k, v := range personInfo {
		fmt.Printf("Index %v \t Value %v", k, v)
	}
}
