package sortimpl

import (
	"fmt"
	"github.com/thousanda/go-sort-algos/sort"
)

type insertionSorter struct{}

func NewInsertionSorter() sort.Sorter {
	return new(insertionSorter)
}

func (s *insertionSorter) Sort(inSlc []int64) []int64 {
	n := len(inSlc)
	outSlc := make([]int64, n)

	// ソート開始
	for i := 0; i < n; i++ {
		newNum := inSlc[i]
		fmt.Println("i: ", i)
		fmt.Println("new num: ", newNum)
		// 新しい値を入れる場所を決める
		insertPlace := 0
		// 左から順に新しい値と比較していく
		for j := 0; j < i; j++ {
			// 新しい値の方が小さくなったらそこ
			if newNum < outSlc[j] {
				insertPlace = j
				break
			}
			// 最後まで行っても上記の条件を満たさなかったときは一番右に入れる
			if j == i-1 {
				insertPlace = i
			}
		}
		fmt.Println("ここ！", insertPlace)

		// 入れたい場所とそれより右のやつを右にずらす
		for j := i; j > insertPlace; j-- {
			outSlc[j] = outSlc[j-1]
		}
		fmt.Println("ずらしたよ: ", outSlc)

		// 新しい値を入れる
		outSlc[insertPlace] = newNum
		fmt.Println("途中経過\n===\n", outSlc)
		fmt.Println("===\n")
	}

	return outSlc
}
