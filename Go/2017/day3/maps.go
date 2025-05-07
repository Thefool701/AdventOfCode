package main

import "fmt"

func main() {
	// Create the outer map
	outerMap := make(map[int]map[int]int)

	// Initialize and populate the first inner map
	innerMap1 := make(map[int]int)
	innerMap1[1] = 10
	innerMap1[2] = 20

	// Assign the first inner map to the outer map
	outerMap[1] = innerMap1

	// Initialize and populate the second inner map
	innerMap2 := make(map[int]int)
	innerMap2[3] = 30
	innerMap2[4] = 40

	// Assign the second inner map to the outer map
	outerMap[2] = innerMap2

	// Accessing values
	fmt.Println("Outer Map:", outerMap)
	fmt.Println("Value at outerMap[1][1]:", outerMap[1][1]) // Output: 10
	fmt.Println("Value at outerMap[2][3]:", outerMap[2][3]) // Output: 30

	// Iterating over the outer map
	for key, innerMap := range outerMap {
		fmt.Println("Key:", key)
		for innerKey, value := range innerMap {
			fmt.Printf("  Inner Key: %d, Value: %d\n", innerKey, value)
		}
	}
}
