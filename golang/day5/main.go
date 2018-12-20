package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"

	lane "gopkg.in/oleiade/lane.v1"
)

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func parseInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	seq := scanner.Text()
	return seq
}

func splitSequence(str string, exclude_rune rune) *lane.Deque {
	var seq *lane.Deque = lane.NewDeque()
	for _, c := range str {
		if exclude_rune != 0 && unicode.ToLower(c) == unicode.ToLower(exclude_rune) {
			continue
		}
		seq.Append(c)
	}
	return seq
}

func isMatch(c1 rune, c2 rune) bool {
	if unicode.ToLower(c1) == unicode.ToLower(c2) {
		if unicode.IsLower(c1) && unicode.IsUpper(c2) ||
			unicode.IsUpper(c1) && unicode.IsLower(c2) {
			return true
		}
	}
	return false
}

func fold(seq *lane.Deque) *lane.Deque {

	output_seq := lane.NewDeque()

	for {
		c := seq.Shift()

		if c == nil {
			break
		}

		last_char := output_seq.Last()

		if last_char != nil && isMatch(c.(rune), last_char.(rune)) {
			// destroy the pair
			output_seq.Pop()
			continue
		}
		output_seq.Append(c)
	}

	return output_seq

}

func main() {
	input_str := parseInput()
	sequence := splitSequence(input_str, 0)

	min_len := fold(sequence).Size()

	fmt.Printf("Number of fragments: %d\n", min_len)

	for _, c := range "abcdefghijklmnopqrstuvwxyz" {
		sequence := splitSequence(input_str, c)
		min_len = min(fold(sequence).Size(), min_len)
	}

	fmt.Printf("Minimal possible sequence length: %d\n", min_len)

}
