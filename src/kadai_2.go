package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"        
	"bufio"
)

// メイン関数
// 期待するコマンドライン引数：「./kadai_1 test.txt」の様な形式
func main() {

	// 引数チェック
	if len(os.Args) > 1 {
		
		var input_data []string
        
        // 入力ファイル
        fp, err := os.Open(os.Args[1])
        
        if err != nil {
	         // エラーで処理抜け
            panic(err)
        }
    	// close処理
		defer fp.Close()
   
	    scanner := bufio.NewScanner(fp)
	    // cnt := 0
        for scanner.Scan() {
            input_data = append(input_data,scanner.Text())
        }
        // エラー確認
        if err := scanner.Err(); err != nil {
	         // エラーで処理抜け
            panic(err)
        }
	    
	    var csv_data string
	    
	    for _, value := range input_data {

			// 引数を数値化
			i, _ := strconv.Atoi(value)
	
			fmt.Println("引数:", i)
			
			// 素因数分解結果
			ans, err := MainFactorNumber(i)
	
			if err == nil {
				fmt.Println("素因数分解の回答:", ans)
				csv_data += ans + "\n"
			} else {
				// エラーを出力して処理は継続
				fmt.Println("エラー:", err)
			}
		}
		
		// 結果出力ファイル
		fp_out, err := os.OpenFile("result.csv", os.O_WRONLY|os.O_CREATE, 0666)
	    if err != nil {
	         // エラーで処理抜け
	         panic(err)
	    }
	    // close処理
	    defer fp_out.Close()
	    
	    // 書き込み
	    fp_out.Write(([]byte)(csv_data))
	    
	} else {
		fmt.Println("エラー:引数が不足しています")
	}
}

// 素因数分解メイン処理
// x:正の自然数
// result:xの素因数分解結果文字列
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
	
    for (count * count) <= x{
        if x % count == 0{
            x /= count
            result += strconv.Itoa(count) + ","
            result = ProcFactorNumber(x,count,result)
            return;
        }else{
            count++
        }
    }
    // 最後の素数
	result += strconv.Itoa(x)
	return
}
