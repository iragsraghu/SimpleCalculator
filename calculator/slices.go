package calculator

import "fmt"

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
