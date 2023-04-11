package sortimpl

import (
	"fmt"
	"github.com/thousanda/go-sort-algos/sort"
)

type quickSorter struct{}

func NewQuickSorter() sort.Sorter {
	return new(quickSorter)
}

func (s *quickSorter) Sort(inSlc []int64) []int64 {
	n := len(inSlc)

	// 0個か1個ならそのまま返す
	if n == 0 || n == 1 {
		fmt.Println("1個です", inSlc)
		return inSlc
	}

	// 2個なら必要に応じて入れ替えて返す
	if n == 2 {
		fmt.Println("2個です", inSlc)
		if inSlc[0] < inSlc[1] {
			return []int64{inSlc[0], inSlc[1]}
		}
		return []int64{inSlc[1], inSlc[0]}
	}

	// 3個以上ならpivotを決めて分割していく

	// 今回は適当に一番前の値をpivotとする
	// NOTE: 高速にできるだけ中央値に近い値を選ぶとパフォーマンスが上がる
	pivot := inSlc[0]
	fmt.Println("pivot: ", pivot)

	// pivotより小さいやつ (left) と 大きいやつ (right) を入れておくスライス
	left := make([]int64, n)
	right := make([]int64, n)

	// leftとrightにそれぞれ何個入っているか
	var leftNum, rightNum int

	// leftとrightに振り分けていく
	for i := 1; i < n; i++ {
		if inSlc[i] < pivot { // pivotより小さいものをleftに
			left[leftNum] = inSlc[i]
			leftNum++
		} else { // pivot以上のものをrightに
			right[rightNum] = inSlc[i]
			rightNum++
		}
	}
	fmt.Println("振り分けた", left, leftNum, pivot, right, rightNum)

	// pivotはrightの一部に入れる これいる？？？？？
	// right[rightNum] = pivot

	// leftとrightそれぞれに再度クイックソートを実行する
	var sortedLeft, sortedRight []int64
	if leftNum > 0 {
		sortedLeft = s.Sort(left[:leftNum])
	}
	if rightNum > 0 {
		sortedRight = s.Sort(right[:rightNum])
	}

	// leftとpivotとrightを結合
	outSlc := make([]int64, n)
	for i := 0; i < leftNum; i++ {
		outSlc[i] = sortedLeft[i]
	}
	outSlc[leftNum] = pivot
	for i := 0; i < rightNum; i++ {
		outSlc[i+leftNum+1] = sortedRight[i]
	}

	return outSlc
}
