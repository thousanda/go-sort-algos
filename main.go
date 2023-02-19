package main

import "fmt"

func main() {
	inSlc := []int64{4, 1, 6, 5, 2, 9, 7, 10, 8, 3}
	// inSlc := []int64{1, 2, 3, 4, 5, 6, 9, 7, 10, 8}
	fmt.Printf("input: %v\n", inSlc)
	// outSlc := bubbleSort(inSlc)
	outSlc := selectionSort(inSlc)
	fmt.Printf("output: %v\n", outSlc)
}

func bubbleSort(inSlc []int64) []int64 {
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

func selectionSort(inSlc []int64) []int64 {
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
