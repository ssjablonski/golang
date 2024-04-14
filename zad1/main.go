package main

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
	"strings"
)

func factorialContainsBytes(nickBytes []int, n int) int {
	factorial := new(big.Int).MulRange(1, int64(n))
	factorialString := factorial.String()
	for _, code := range nickBytes {
		if !strings.Contains(factorialString, strconv.Itoa(code)) {
			return factorialContainsBytes(nickBytes, n+1)
		}
	}
	return n
}

var callCounts = make(map[int]int)

func fibonacci(n int) int {
	callCounts[n]++

	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func closestToStrongNumber(callCounts map[int]int, strongNumber int) int {
	closestKey := 0
	smallestDiff := math.MaxInt32
	for key, value := range callCounts {
		diff := abs(strongNumber - value)
		if diff < smallestDiff {
			smallestDiff = diff
			closestKey = key
		}
	}
	return closestKey
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func main() {
	var name string
	var surname string
	var nick string

	fmt.Println("Enter your name: ")
	fmt.Scan(&name)

	fmt.Println("Enter your surname: ")
	fmt.Scan(&surname)

	nick = name[0:3] + surname[0:3]
	lowerNick := strings.ToLower(nick)
	fmt.Println("Your nick in lowercase is:", lowerNick)
	nickASCII := []int{int(lowerNick[0]), int(lowerNick[1]), int(lowerNick[2]), int(lowerNick[3]), int(lowerNick[4]), int(lowerNick[5])}

	strongNumber := factorialContainsBytes(nickASCII, 1)
	fmt.Println("Your strong number is:", strongNumber)

	fibonacci(30)
	weakNumber := closestToStrongNumber(callCounts, strongNumber)
	fmt.Println("Your weak number is:", weakNumber)

}
