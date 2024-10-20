package test

import (
	"fmt"
	"strings"
	"testing"
)

const (
	strNumber = "0123456789"
	strAz     = "abcdefghijklmnopqrstuvwxyz"
	strAZ     = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func TestPass(t *testing.T) {

	results := PermutationsStr(strings.Split(strNumber, ""), 4)

	for s := range results {
		fmt.Printf("%#v\n", strings.Join(s, ""))
	}
}

func gen(words []string, minLen, maxLen int, check func(pwd string) bool) {
	//sourceLen := len(words)
	//for passLen := minLen; passLen <= maxLen; minLen++ {
	//	go gen2(words, passLen, check)
	//}
}

func gen2(words []string, passLen int, check func(pwd string) bool) {
	//sourceLen := len(words)
	//for sI := 0; sI < sourceLen; sI++ {
	//	source := sourceList[0]
	//
	//}
}

func gen3(word string, passLen int, check func(pwd string) bool) {
	//wordLen := len(word)
	//for i := 0; i < wordLen; i++ {
	//	char := word[i]
	//}
}

func gen4(word string, passLen int, check func(pwd string) bool) {
	//for {
	//
	//}
	//sourceLen := len(words)
	//for i := 0; i < passLen; i++ {
	//	source := sourceList[0]
	//
	//}
}

// GenCombinations generates, from two natural numbers n > r,
// all the possible combinations of r indexes taken from 0 to n-1.
// For example if n=3 and r=2, the result will be:
// [0,1], [0,2] and [1,2]
func GenCombinations(n, r int) <-chan []int {

	if r > n {
		panic("Invalid arguments")
	}

	ch := make(chan []int)

	go func() {
		result := make([]int, r)
		for i := range result {
			result[i] = i
		}

		temp := make([]int, r)
		copy(temp, result) // avoid overwriting of result
		ch <- temp

		for {
			for i := r - 1; i >= 0; i-- {
				if result[i] < i+n-r {
					result[i]++
					for j := 1; j < r-i; j++ {
						result[i+j] = result[i] + j
					}
					temp := make([]int, r)
					copy(temp, result) // avoid overwriting of result
					ch <- temp
					break
				}
			}
			if result[0] >= n-r {
				break
			}
		}
		close(ch)

	}()
	return ch
}

// CombinationsStr generates all the combinations of r elements
// extracted from an slice of strings
func CombinationsStr(iterable []string, r int) chan []string {

	ch := make(chan []string)

	go func() {

		length := len(iterable)

		for comb := range GenCombinations(length, r) {
			result := make([]string, r)
			for i, val := range comb {
				result[i] = iterable[val]
			}
			ch <- result
		}

		close(ch)
	}()
	return ch
}

// GenPermutations generates, given a number n,
// all the n factorial permutations of the integers
// from 0 to n-1
func GenPermutations(n int) <-chan []int {
	if n < 0 {
		panic("Invalid argument")
	}

	ch := make(chan []int)

	go func() {
		var finished bool

		result := make([]int, n)

		for i := range result {
			result[i] = i
		}

		temp := make([]int, n)
		copy(temp, result) // avoid overwriting of result
		ch <- temp

		for {
			finished = true

			for i := n - 1; i > 0; i-- {

				if result[i] > result[i-1] {
					finished = false

					minGreaterIndex := i
					for j := i + 1; j < n; j++ {
						if result[j] < result[minGreaterIndex] && result[j] > result[i-1] {
							minGreaterIndex = j
						}

					}

					result[i-1], result[minGreaterIndex] = result[minGreaterIndex], result[i-1]

					//sort from i to n-1
					for j := i; j < n; j++ {
						for k := j + 1; k < n; k++ {
							if result[j] > result[k] {
								result[j], result[k] = result[k], result[j]
							}

						}
					}
					break
				}
			}

			if finished {
				close(ch)
				break
			}
			temp := make([]int, n)
			copy(temp, result) // avoid overwriting of result
			ch <- temp

		}

	}()
	return ch
}

// PermutationsStr generates all the permutations of r elements
// extracted from an slice of strings
func PermutationsStr(iterable []string, r int) chan []string {

	ch := make(chan []string)

	go func() {

		length := len(iterable)

		for comb := range GenCombinations(length, r) {
			for perm := range GenPermutations(r) {
				result := make([]string, r)
				for i := 0; i < r; i++ {
					result[i] = iterable[comb[perm[i]]]
				}
				ch <- result
			}
		}

		close(ch)
	}()
	return ch
}

// PermutationsList generates all the permutations of r elements
// extracted from a List (an arbitrary list of elements).
// A List can be created, for instance, as follows:
// myList := List{"a", "b", 13, 3.523}
func PermutationsList(iterable []any, r int) chan []any {

	ch := make(chan []any)

	go func() {

		length := len(iterable)

		for comb := range GenCombinations(length, r) {
			for perm := range GenPermutations(r) {
				result := make([]any, r)
				for i := 0; i < r; i++ {
					result[i] = iterable[comb[perm[i]]]
				}
				ch <- result
			}
		}

		close(ch)
	}()
	return ch
}
