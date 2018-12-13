package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func parseInput() []int {
	var numbers []int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		value, _ := strconv.Atoi(scanner.Text())
		numbers = append(numbers, value)
	}
	return numbers
}

func main() {
	var netValue int
	var repeatedValue int
	var freq = make(map[int]uint)

	var numbers = parseInput()

outer:
	for iter := 0; ; iter++ {

		for _, value := range numbers {
			netValue += value

			freq[netValue]++
			if freq[netValue] > 1 {
				repeatedValue = netValue
				break outer
			}
		}

		if iter == 0 {
			fmt.Println("Result value:", netValue)
		}

	}

	fmt.Println("Repeated value:", repeatedValue)

}
