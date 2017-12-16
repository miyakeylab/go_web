package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// メイン関数
// 期待するコマンドライン引数：「./kadai_1 28」の様な形式
func main() {

	var i int
	// 引数チェック
	if len(os.Args) > 1 {
		// 引数を数値化
		i, _ = strconv.Atoi(os.Args[1])

		fmt.Println("引数:", i)
		// 約数の足し算結果
		ans, err := ProcFactorNumber(i)

		if ans > 0 {
			fmt.Println("素因数分解の回答:", ans)
		} else {
			fmt.Println("エラー:", err)
		}
	} else {
		fmt.Println("エラー:引数が不足しています")
	}
}

// 素因数分解処理
// x:正の自然数
// result:xの約数の足し算結果（エラー時:0）
// errObj:エラー型(エラーメッセージ用)
func ProcFactorNumber(x int) (result int, errObj error) {
	result = 0

	// 引数チェック
	if x > 0 {
       i := 2;
       tmp := x;

        for (i * i) <= x{
            if tmp % i == 0{
                tmp /= i
                fmt.Println("回答:", i)
                //fmt.Println("tmp:", tmp)
                ProcFactorNumber(tmp)
                break;
            }else{
                i++
            }
        }
        
        fmt.Println("終わり回答:", tmp)
   
        
	} else {
		errObj = errors.New("正の自然数ではありません")
	}
	return
}
