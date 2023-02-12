package main

import "fmt"

func main() {
	// inSls := []int64{4, 1, 6, 5, 2, 9, 7, 10, 8, 3}
	inSls := []int64{1, 2, 3, 4, 5, 6, 9, 7, 10, 8}
	fmt.Printf("input: %v\n", inSls)
	outSls := bubbleSort(inSls)
	fmt.Printf("output: %v\n", outSls)
}

func bubbleSort(inSls []int64) []int64 {
	n := len(inSls)
	outSls := make([]int64, n)

	// コピー
	for i := 0; i < n; i++ {
		outSls[i] = inSls[i]
	}

	// ソート開始
	for i := n - 1; i > 0; i-- { // n-1, n-2, ..., 1 (n-1 loops)
		fmt.Printf("Loop: %d\n", i)
		isNoChaned := true
		for j := 0; j < i; j++ { // 0, 1, 2, ..., i-1 (i loops)
			if outSls[j] > outSls[j+1] {
				// j番目とj+1番目を交換
				tmp := outSls[j]
				outSls[j] = outSls[j+1]
				outSls[j+1] = tmp
				isNoChaned = false
			}
		}
		fmt.Printf("%v\n", outSls)
		// 一度も交換してなかったらソート済みなので返す
		if isNoChaned {
			return outSls
		}
	}

	return outSls
}
