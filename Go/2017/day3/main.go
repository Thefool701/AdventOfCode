/* --- Day 3: Spiral Memory --- */
// INCOMPLETE: Problem with Maps
package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	var input int
	infile, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer infile.Close()

	scanner := bufio.NewScanner(infile)

	for scanner.Scan() {
		temp := scanner.Text()
		input, err = strconv.Atoi(temp)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Part 1:", findShortestPath(input))
	fmt.Println("Part 2:", findValueLargerThanPuzzleInput(input))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func findShortestPath(input int) (result int) {
	ringNr := math.Floor(math.Ceil(math.Sqrt(float64(input))) / 2)

	ringFirstNr := math.Pow(ringNr*2-1, 2) + 1
	ringLastNr := math.Pow(ringNr*2+1, 2)
	nrOfItemsInRing := ringLastNr - ringFirstNr + 1
	positionInRingFromBottomRight := input - int(ringFirstNr) + 1
	nrOfItemsPerLoop := nrOfItemsInRing / 4
	positionInLoop := positionInRingFromBottomRight % int(nrOfItemsPerLoop)
	bonusExpense := math.Abs(ringNr - float64(positionInLoop))
	result = int(math.Abs(ringNr + float64(bonusExpense)))
	return result
}

func findValueLargerThanPuzzleInput(input int) int {
	curLoc := complex(0, 0)
	curDir := complex(0, 1)
	numSteps, sz, left, curSteps := 1, 1, 1, 0
	// FIX: Entry into nil map error
	vals := make(map[int]map[int]int)

	innerVals1 := make(map[int]int)
	innerVals1[0] = 1

	vals[0] = innerVals1
	innerVals2 := make(map[int]int)
	innerVals2[0] = 1

	vals[1] = innerVals2
	for vals[int(real(curLoc))][int(imag(curDir))] <= input {
		numSteps++
		curSteps++
		curLoc += curDir
		if curSteps == sz {
			if left == 1 {
				left--
				curDir *= complex(0, 1)
			} else {
				left = 1
				curDir *= complex(0, 1)
				sz++
			}
			curSteps = 0
		}
		vals[int(real(curLoc))][int(imag(curDir))] +=
			vals[int(real(curLoc))+1][int(imag(curDir))]
		vals[int(real(curLoc))][int(imag(curDir))] +=
			vals[int(real(curLoc))-1][int(imag(curDir))]
		vals[int(real(curLoc))][int(imag(curDir))] +=
			vals[int(real(curLoc))][int(imag(curDir))+1]
		vals[int(real(curLoc))][int(imag(curDir))] +=
			vals[int(real(curLoc))][int(imag(curDir))-1]
		vals[int(real(curLoc))][int(imag(curDir))] +=
			vals[int(real(curLoc))+1][int(imag(curDir))+1]
		vals[int(real(curLoc))][int(imag(curDir))] +=
			vals[int(real(curLoc))-1][int(imag(curDir))-1]
		vals[int(real(curLoc))][int(imag(curDir))] +=
			vals[int(real(curLoc))-1][int(imag(curDir))+1]
		vals[int(real(curLoc))][int(imag(curDir))] +=
			vals[int(real(curLoc))+1][int(imag(curDir))-1]

	}
	result := vals[int(real(curLoc))][int(imag(curDir))]

	return result
}
