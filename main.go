package main

import (
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	list := os.Args[1:]

	if len(list) == 0 {
		fmt.Printf("Usage: %s <word1> [word2] [word3] ...\n", os.Args[0])
		os.Exit(1)
	}

	permuteArgs(list)
}

func permuteArgs(words []string) {
	sort.Strings(words)

	var buffer bytes.Buffer
	for {
		buffer.Reset()

		for _, w := range words {
			buffer.WriteString(w)
		}

		fmt.Println(buffer.String())

		if !nextPermutation(words) {
			break
		}
	}
}

func nextPermutation(arr []string) bool {
	n := len(arr)

	k := -1
	for i := n - 2; i >= 0; i-- {
		if arr[i] < arr[i+1] {
			k = i
			break
		}
	}

	if k == -1 {
		return false
	}

	l := -1
	for i := n - 1; i > k; i-- {
		if arr[k] < arr[i] {
			l = i
			break
		}
	}

	arr[k], arr[l] = arr[l], arr[k]

	k++
	n--
	for k < n {
		arr[k], arr[n] = arr[n], arr[k]
		k++
		n--
	}

	return true
}
