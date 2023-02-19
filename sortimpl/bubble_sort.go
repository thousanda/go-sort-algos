package sortimpl

import (
	"fmt"
	"github.com/thousanda/go-sort-algos/sort"
)

type bubbleSorter struct{}

func NewBubbleSorter() sort.Sorter {
	return new(bubbleSorter)
}

func (s *bubbleSorter) Sort(inSlc []int64) []int64 {
	n := len(inSlc)
	outSlc := make([]int64, n)

	// コピー
	for i := 0; i < n; i++ {
		outSlc[i] = inSlc[i]
	}

	// ソート開始
	for i := n - 1; i > 0; i-- { // n-1, n-2, ..., 1 (n-1 loops)
		fmt.Printf("Loop: %d\n", i)
		isNoChaned := true
		for j := 0; j < i; j++ { // 0, 1, 2, ..., i-1 (i loops)
			if outSlc[j] > outSlc[j+1] {
				// j番目とj+1番目を交換
				tmp := outSlc[j]
				outSlc[j] = outSlc[j+1]
				outSlc[j+1] = tmp
				isNoChaned = false
			}
		}
		fmt.Printf("%v\n", outSlc)
		// 一度も交換してなかったらソート済みなので返す
		if isNoChaned {
			return outSlc
		}
	}

	return outSlc
}
