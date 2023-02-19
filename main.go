package main

import "fmt"

func main() {
	inSlc := []int64{4, 1, 6, 5, 2, 9, 7, 10, 8, 3}
	// inSlc := []int64{1, 2, 3, 4, 5, 6, 9, 7, 10, 8}
	fmt.Printf("input: %v\n", inSlc)
	// outSlc := bubbleSort(inSlc)
	// outSlc := selectionSort(inSlc)
	outSlc := mergeSort(inSlc)
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

func mergeSort(inSlc []int64) []int64 {
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
	left := mergeSort(inSlc[:c])
	right := mergeSort(inSlc[c:])
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
