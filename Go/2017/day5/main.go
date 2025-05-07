package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	infile, err := os.Open("testcase.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer infile.Close()
	var length int
	arr := []int{}
	scanner := bufio.NewScanner(infile)
	for scanner.Scan() {
		temp := scanner.Text()
		fileInput, _ := strconv.Atoi(temp)
		arr = append(arr, fileInput)
		length++
	}

	fmt.Println("Part 1:", findTheExit(arr, length))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

func findTheExit(arr []int, length int) int {
	var offset int
	var position int
	var steps int
	for i := range arr {
		offset = arr[position]
		arr[position]++
		position += offset
		steps++
		if position > length {
			fmt.Println(i)
			fmt.Println("escaped the maze...")
			break
		}
	}
	fmt.Println(arr)
	return steps
}
