package sortimpl

import (
	"fmt"
	"github.com/thousanda/go-sort-algos/sort"
)

type selectionSorter struct{}

func NewSelectionSorter() sort.Sorter {
	return new(selectionSorter)
}

func (s *selectionSorter) Sort(inSlc []int64) []int64 {
	n := len(inSlc)
	outSlc := make([]int64, n)

	// コピー
	for i := 0; i < n; i++ {
		outSlc[i] = inSlc[i]
	}

	for i := 0; i < n-1; i++ {
		// 最小値を探す
		minIdx := i
		minVal := outSlc[i]
		for j := i + 1; j < n; j++ {
			if outSlc[j] < minVal {
				minIdx = j
				minVal = outSlc[j]
			}
		}
		// 最小値を未整列部分の先頭に移動する
		tmp := outSlc[i]
		outSlc[i] = outSlc[minIdx]
		outSlc[minIdx] = tmp
		fmt.Printf("%v\n---\n", outSlc)
	}

	return outSlc
}
