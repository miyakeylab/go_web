package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// メイン関数
// 期待するコマンドライン引数：「./kadai_3 http://google.co.jp google」の様な(urlと検索ワード)形式
func main() {

	// 引数チェック
	if len(os.Args) > 2 {

		rtn, _ := http.Get(os.Args[1])
		defer rtn.Body.Close()

		byteArray, _ := ioutil.ReadAll(rtn.Body)
		body := string(byteArray)

		if strings.Index(body, os.Args[2]) != -1 {
			//含まれる場合の処理
			fmt.Println("含む")
			GetHttpData(os.Args[1])
		} else {
			//含まれない場合の処理
			fmt.Println("含まない")
		}
	} else {
		fmt.Println("エラー:引数が不足しています")
	}
}

func GetHttpData(url string) {
	doc, _ := goquery.NewDocument(url)
	n := 0
	doc.Find("a").Each(func(_ int, s *goquery.Selection) {
		url, _ := s.Attr("href")
		if strings.HasPrefix(url, os.Args[2]) == true {
			fmt.Println(string(n) + url)
			n++
		} else {
			fmt.Println(string(n))
		}

	})
}
