package sortimpl

import (
	"fmt"
	"github.com/thousanda/go-sort-algos/sort"
)

type mergeSorter struct{}

func NewMergeSorter() sort.Sorter {
	return new(mergeSorter)
}

func (s *mergeSorter) Sort(inSlc []int64) []int64 {
	fmt.Printf("input: %v\n", inSlc)
	n := len(inSlc)
	if n == 1 { // 1要素だったらそのまま返す
		return inSlc
	} else if n == 2 { // 2要素だったら昇順にして返す
		if inSlc[0] < inSlc[1] {
			return inSlc
		} else {
			return []int64{inSlc[1], inSlc[0]}
		}
	}

	// 3要素以上だったら分割する
	c := n / 2 // 真ん中らへん
	left := s.Sort(inSlc[:c])
	right := s.Sort(inSlc[c:])
	fmt.Printf("sorted! left: %v, right: %v\n", left, right)

	// 分割してソートしたスライスをマージする
	outSlc := make([]int64, n)
	il := 0
	ir := 0
	for i := 0; i < n; i++ {
		// どちらかを使い切っていたら残っているほうから取る
		if il == len(left) { // leftを使い切っていたら右から取る
			fmt.Printf("左が空なので右を取ります: %v\n", right[ir])
			outSlc[i] = right[ir]
			ir++
			continue
		}
		if ir == len(right) { // rightを使い切っていたら左から取る
			fmt.Printf("右が空なので左を取ります: %v\n", left[il])
			outSlc[i] = left[il]
			il++
			continue
		}

		// 大小比較して小さい方を取る
		if left[il] < right[ir] { // 左の方が小さいなら左を取る
			fmt.Printf("left[il] selected: %v\n", left[il])
			outSlc[i] = left[il]
			il++
		} else { // 右の方が小さいなら右を取る
			fmt.Printf("right[ir] selected: %v\n", right[ir])
			outSlc[i] = right[ir]
			ir++
		}
	}

	return outSlc
}
