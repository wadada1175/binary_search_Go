package main

import (
	"fmt"
	"time"
)

// array[0..n-1]からtargetを探す
// targetの位置をpとして
// targetがないとき、p=-1
// targetがあるとき、0<=p<=n-1 && target=array[p]
func main() {
	data := []int{
		101, 103, 107, 109, 113, 127, 131, 137, 139, 149,
		151, 157, 163, 167, 173, 179, 181, 191, 193, 197,
		199, 211, 223, 227, 229, 233, 239, 241, 251, 257,
		263, 269, 271, 277, 281, 283, 293, 307, 311, 313,
		317, 331, 337, 347, 349, 353, 359, 367, 373, 379,
		383, 389, 397, 401, 409, 419, 421, 431, 433, 439,
	}
	target := 433

	start := time.Now() //時間計測開始

	p := solve(data, target)

	fmt.Printf("値 %d は配列の %d 番目に見つかりました。", target, p)
	t := time.Since(start)
	fmt.Println("\n処理時間は", t)
}

func solve(data []int, target int) int {

	arrayLen := len(data)
	low := 0
	high := arrayLen - 1
	var mid int
	for {
		if high < low {
			return -1
		}
		mid = (low + high) / 2

		if data[mid] == target {
			return mid
		}

		if data[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

}
