package main

import (
	"fmt"
	"sort"
	"os"
	"bytes"
)

func main() {
	list := os.Args[1:]

	resultCh := make(chan string)

	go permuteArgs(resultCh, list)

    for p := range resultCh {
        fmt.Println(p)
    }
}

func permuteArgs(ch chan string, words []string) {
	defer close(ch)

	ptrs := make([]*string, len(words))

	for i := range words {
		ptrs[i] = &words[i]
	}

	sort.Slice(ptrs, func(i, j int) bool {
		return *ptrs[i] < *ptrs[j]
	})

	var buffer bytes.Buffer
	for {
		buffer.Reset()

		for _, ptr := range ptrs {
			buffer.WriteString(*ptr)
		}

		ch <- buffer.String()

		if !nextPermutation(ptrs) {
			break
		}
	}

	return
}

func nextPermutation(arr []*string) bool {
	n := len(arr)

	k := -1
	for i := n - 2; i >= 0; i-- {
		if *arr[i] < *arr[i+1] {
			k = i
			break
		}
	}

	if k == -1 {
		return false
	}

	l := -1
	for i := n - 1; i > k; i-- {
		if *arr[k] < *arr[i] {
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
