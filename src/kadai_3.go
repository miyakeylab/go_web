package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

var g_url_list []string // 全体リンクリスト

// メイン関数
// 期待するコマンドライン引数：「./kadai_3 https://www.google.co.jp/ google」の様な(urlと検索ワード)形式
func main() {

	// 引数チェック
	if len(os.Args) > 2 {
		ans := ""
		// 引数のURLに対する検索
		if SearchHttpString(os.Args[1], os.Args[2]) == true {
			//含まれる場合の処理終了
			ans = os.Args[1]

		} else {
			// リンクを保持(重複チェック用)
			g_url_list = append(g_url_list, os.Args[1])

			main_list := GetHttpLink(os.Args[1])
			//含まれない場合の処理
			_, ans = MainLinkLoop(main_list, os.Args[2])

		}

		if ans != "" {
			fmt.Println("回答：" + ans)
		} else {
			fmt.Println("該当URLなし")
		}

	} else {
		fmt.Println("エラー:引数が不足しています")
	}
}

func MainLinkLoop(url_list []string, search string) (exist bool, ans string) {

	exist = false
	//含まれない場合の処理
	for _, value := range url_list {

		if SearchHttpString(value, search) == true {
			ans = value
			exist = true
			break
		}
	}
	if exist == false {
		for _, value := range url_list {
			sub_list := GetHttpLink(value)
			exist, ans = MainLinkLoop(sub_list, search)
		}
	}
	return
}

// 基底URLからリンクを抽出し、リストにして返す
// base_url:リンクを抽出するリンク
// url_list:リンクリスト
func GetHttpLink(base_url string) (url_list []string) {

	fmt.Println("【基底URL】:" + base_url)

	doc, _ := goquery.NewDocument(base_url)

	// リンクチェック
	doc.Find("a").Each(func(_ int, s *goquery.Selection) {
		html_url, _ := s.Attr("href")

		// 対象リンクになるか確認
		if strings.Index(html_url, base_url) == 0 {

			// 階層が下であることを確認
			if len(base_url) < len(html_url) {
				// 存在チェック
				if contains(url_list, html_url) == false && contains(g_url_list, html_url) == false {
					url_list = append(url_list, html_url)
					g_url_list = append(g_url_list, html_url)
				}
			}
		} else if strings.Index(html_url, "/") == 0 && strings.Index(html_url, "//") != 0 {
			html_url = strings.Replace(html_url, "/", "", 1)
			cnt_str := base_url + html_url
			tmp_str := os.Args[1] + html_url
			// 存在チェック
			if contains(url_list, cnt_str) == false && contains(g_url_list, cnt_str) == false && contains(g_url_list, tmp_str) == false && contains(g_url_list, html_url) == false {
				url_list = append(url_list, cnt_str)
				g_url_list = append(g_url_list, cnt_str)
				g_url_list = append(g_url_list, html_url)
			}
		}
	})
	return
}

// URLから取得したHttp文字列から特定の文字列が存在するか検索する
// access_url:検索用リンク
// search:検索キーワード
// exist:結果
func SearchHttpString(access_url string, search string) (exist bool) {

	fmt.Println("検索URL：" + access_url)
	exist = false

	// Httpアクセス
	rtn, _ := http.Get(access_url)

	// close処理
	defer rtn.Body.Close()

	byteArray, _ := ioutil.ReadAll(rtn.Body)
	body := string(byteArray)

	// 検索ワードが含まれるか
	if strings.Contains(body, search) == true {
		exist = true
	}

	return
}

// スライスの存在チェック処理
func contains(str []string, data string) bool {
	for _, value := range str {
		if value == data {
			return true
		}
	}
	return false
}
