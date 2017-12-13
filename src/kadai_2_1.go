package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {

	var i int
	// 引数チェック
	if len(os.Args) > 1 {
		// 引数を数値化
		i, _ = strconv.Atoi(os.Args[1])

		fmt.Println("引数:", i)
		// 約数の足し算結果
		ans, err := ProcDivisorAdd(i)

		if ans > 0 {
			fmt.Println("約数の合計:", ans)
		} else {
			fmt.Println("エラー:", err)
		}
	} else {
		fmt.Println("エラー:引数が不足しています")
	}
}

// 約数の足し算処理
// x:正の自然数
// result:xの約数の足し算結果（エラー時:0）
// errObj:エラー型(エラーメッセージ用)
func ProcDivisorAdd(x int) (result int, errObj error) {
	result = 0

	// 引数チェック
	if x > 0 {
		n := 0
		for n < x {
			// 約数チェック
			if x%(n+1) == 0 {
				result += (n + 1)
			}
			n++
		}
	} else {
		errObj = errors.New("正の自然数ではありません")
	}
	return
}
