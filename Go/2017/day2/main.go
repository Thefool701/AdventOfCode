/* --- Day 2: Corruption Checksum --- */

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	infile, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer infile.Close()

	scanner := bufio.NewScanner(infile)
	var checkSum int
	var checkSum2 int
	for scanner.Scan() {
		checkSum += findRowDifference(scanner.Text())
		checkSum2 += findRowEvenDivisibility(scanner.Text())
	}

	fmt.Println("Part 1:", checkSum)
	fmt.Println("Part 2:", checkSum2)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

func findRowDifference(input string) int {
	var temp []string
	var row []int
	input = strings.Replace(input, "\t", " ", -1)
	temp = strings.Split(input, " ")
	for i := range temp {
		foo, err := strconv.Atoi(temp[i])
		row = append(row, foo)
		if err != nil {
			log.Fatal(err)
		}

	}
	rowMin := slices.Min(row)
	rowMax := slices.Max(row)

	return rowMax - rowMin
}

func findRowEvenDivisibility(input string) (frans int) {
	var temp []string
	var row []int

	input = strings.Replace(input, "\t", " ", -1)
	temp = strings.Split(input, " ")
	for i := range temp {
		foo, err := strconv.Atoi(temp[i])
		row = append(row, foo)
		if err != nil {
			log.Fatal(err)
		}
	}

	for i := range row {
		for j := range row {
			if row[i] == row[j] {
				continue
			}
			if res := row[i] % row[j]; res == 0 {
				frans = row[i] / row[j]
			}
		}
	}

	return frans
}
