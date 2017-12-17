package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"        
	"bufio"
)

// メイン関数
// 期待するコマンドライン引数：「./kadai_1 28」の様な形式
func main() {

	// 引数チェック
	if len(os.Args) > 1 {
		
        length := 100 // 上限
		capacity := 200 // 容量
		input_data := make([]string, length, capacity)
        
        // 入力ファイル
        fp, err := os.Open(os.Args[1])
        
        if err != nil {
	         // エラーで処理抜け
            panic(err)
        }
    	// close処理
		defer fp.Close()
    
     
	    scanner := bufio.NewScanner(fp)
	    cnt := 0
        for scanner.Scan() {
            input_data[cnt] = scanner.Text()
            cnt++
        }
        // エラー確認
        if err := scanner.Err(); err != nil {
	         // エラーで処理抜け
            panic(err)
        }
	    
	    var csv_data string
	    
	    for index, value := range input_data {

			if index < cnt {
				// 引数を数値化
				i, _ := strconv.Atoi(value)
		
				fmt.Println("引数:", i)
				
				// 素因数分解結果
				ans, err := MainFactorNumber(i)
		
				if ans !="" {
					fmt.Println("素因数分解の回答:", ans)
					csv_data += ans + "\n"
				} else {
					fmt.Println("エラー:", err)
				}
			}else{
				break;
			}
		}
		
		// 結果出力
		fp_out, err := os.OpenFile("result.csv", os.O_WRONLY|os.O_CREATE, 0666)
	    if err != nil {
	         // エラーで処理抜け
	         panic(err)
	    }
	    defer fp_out.Close()
	    fp_out.Write(([]byte)(csv_data))
	    
	} else {
		fmt.Println("エラー:引数が不足しています")
	}
}

// 素因数分解メイン処理
// x:正の自然数
// result:xの素因数分解結果文字列（エラー時空文字）
// errObj:エラー型(エラーメッセージ用)
func MainFactorNumber(x int) (result string, errObj error) {

	result += strconv.Itoa(x) + ","

	// 引数チェック
	if x > 0 {
		result = ProcFactorNumber(x,2,result)
	} else {
		errObj = errors.New("正の自然数ではありません")
	}
	return
}



// 素因数分解処理
// x:正の自然数
// count:素数
// data:素因数分解の結果文字列
// result:xの素因数分解の結果（xが112の時"2,2,2,2,7"）
func ProcFactorNumber(x int,count int,data string) (result string) {

	result = data
	i := count;
	tmp := x;

    for (i * i) <= x{
        if tmp % i == 0{
            tmp /= i
            result += strconv.Itoa(i) + ","
            result = ProcFactorNumber(tmp,i,result)
            return;
        }else{
            i++
        }
    }
    
	result += strconv.Itoa(tmp)
	return
}
