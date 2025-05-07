/* --- Day 1: Inverse Captcha --- */

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Node struct {
	data int
	next *Node
}

type CircularLinkedList struct {
	head *Node
	tail *Node
}

func NewCircularLinkedList() *CircularLinkedList {
	return &CircularLinkedList{head: nil, tail: nil}
}

func (cll *CircularLinkedList) isEmpty() bool {
	return cll.head == nil
}

func (cll *CircularLinkedList) AddNode(data int) {
	newNode := &Node{data: data, next: nil}

	if cll.isEmpty() {
		cll.head = newNode
		cll.tail = newNode
		newNode.next = cll.head
	} else {
		cll.tail.next = newNode
		cll.tail = newNode
		newNode.next = cll.head
	}
}

func (cll *CircularLinkedList) Length() (count int) {
	if cll.isEmpty() {
		fmt.Println("Circular Linked List is empty")
	} else {
		current := cll.head
		for {
			count++
			current = current.next
			if current == cll.head {
				break
			}
		}
	}
	return count
}
func main() {
	inputList := NewCircularLinkedList()
	infile, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// close file at end of program
	defer infile.Close()

	// read file word by word
	scanner := bufio.NewScanner(infile)
	scanner.Split(bufio.ScanWords)

	// store values into arr
	for scanner.Scan() {
		temp := scanner.Text()
		stringTemp := []rune(temp)

		for _, char := range stringTemp {
			inputList.AddNode(int(char - '0'))
		}
	}

	fmt.Println("Part 1:", findSumThatMatchesNext(inputList))
	fmt.Println("Part 2:", findSumThatMatchesHalfway(inputList, inputList.Length()))
}

func findSumThatMatchesNext(cll *CircularLinkedList) (sum int) {
	if cll.isEmpty() {
		fmt.Println("input list is empty...")
	} else {
		current := cll.head
		for {
			future := current.next.data
			if current.data == future {
				sum += current.data
			}
			current = current.next
			if current == cll.head {
				break
			}
		}
	}
	return sum
}

func findSumThatMatchesHalfway(cll *CircularLinkedList, length int) (sum int) {
	if cll.isEmpty() {
		fmt.Println("input list is empty...")
	} else {
		current := cll.head
		for {
			future := getHalfWayNumber(current, length)
			if current.data == future {
				sum += current.data
			}
			current = current.next
			if current == cll.head {
				break
			}
		}
	}
	return sum
}

func getHalfWayNumber(current *Node, length int) int {
	for i := range length / 2 {
		current = current.next
		if i == length {
			break
		}
	}

	return current.data
}
