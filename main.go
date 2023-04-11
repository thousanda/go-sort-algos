package main

import (
	"fmt"
	"github.com/thousanda/go-sort-algos/sort"
	"github.com/thousanda/go-sort-algos/sortimpl"
)

func main() {
	inSlc := []int64{4, 1, 6, 5, 2, 9, 7, 10, 8, 3}
	// inSlc := []int64{1, 2, 3, 4, 5, 6, 9, 7, 10, 8}
	fmt.Printf("input: %v\n", inSlc)
	s, err := newSorter("insertion")
	if err != nil {
		fmt.Println(err)
		return
	}
	outSlc := s.Sort(inSlc)
	fmt.Printf("output: %v\n", outSlc)
}

func newSorter(mode string) (sort.Sorter, error) {
	switch mode {
	case "bubble":
		return sortimpl.NewBubbleSorter(), nil
	case "selection":
		return sortimpl.NewSelectionSorter(), nil
	case "merge":
		return sortimpl.NewMergeSorter(), nil
	case "insertion":
		return sortimpl.NewInsertionSorter(), nil
	}
	return nil, fmt.Errorf("not implemented mode: %v", mode)
}
