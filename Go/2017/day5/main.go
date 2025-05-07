package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	infile, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer infile.Close()
	var length int
	arr := []int{}
	scanner := bufio.NewScanner(infile)
	for scanner.Scan() {
		temp := scanner.Text()
		angela, _ := strconv.Atoi(temp)
		arr = append(arr, angela)
		length++
	}

	fmt.Println("Part 1:", findTheExit(arr, length))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

func findTheExit(arr []int, length int) int {
	var currPos int // Position in the Array
	var status bool
	var offset int // Instructions of Jumps
	var steps int
	for !status {
		// Get oFfset and Jump
		offset = arr[currPos]
		// Increment Position
		arr[currPos]++
		currPos += offset
		if currPos > length {
			return steps
		}
		steps++
	}
	return steps
}
