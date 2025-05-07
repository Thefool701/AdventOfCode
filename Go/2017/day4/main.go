/* --- Day 4: High-Entropy Passphrases --- */

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	infile, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(infile)
	var validPassphraseCounter int
	var nonAnagramPassPhraseCounter int
	for scanner.Scan() {
		var frans []string
		frans = strings.Split(scanner.Text(), " ")
		validPassphraseCounter += passphraseAuthenticator(frans)
		nonAnagramPassPhraseCounter += anagramPassphraseAuthenticator(frans)
	}
	fmt.Println("Part 1:", validPassphraseCounter)
	fmt.Println("Part 2:", nonAnagramPassPhraseCounter)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func passphraseAuthenticator(frans []string) int {
	passMap := make(map[string]bool)

	for i := range frans {
		_, ok := passMap[frans[i]]
		if !ok {
			passMap[frans[i]] = true
		} else {
			return 0
		}
	}

	return 1
}

func anagramPassphraseAuthenticator(frans []string) int {
	passMap := make(map[string]bool)

	for i := range frans {
		res := []rune(frans[i])
		sort.Slice(res, func(i, j int) bool {
			return res[i] < res[j]
		})
		kindFrans := runesToString(res)
		_, ok := passMap[kindFrans]
		if !ok {
			passMap[kindFrans] = true
		} else {
			return 0
		}
	}

	return 1
}
func runesToString(runes []rune) (outString string) {
	// don't need index so
	for _, v := range runes {
		outString += string(v)
	}
	return
}
