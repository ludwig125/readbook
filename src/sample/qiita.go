package main

import (
	"log"
	//"strings"

	"github.com/sclevine/agouti"
)

//SOY CMSにログインできなかた時に適切にエラーを表示するか？テスト
func main() {
	driver := agouti.ChromeDriver()
	err := driver.Start()
	if err != nil {
		log.Fatal(err)
	}
	page, err := driver.NewPage(agouti.Browser("chrome"))
	err = page.Navigate("https://www.yahoo.co.jp/")
	if err != nil {
		log.Fatal(err)
	}
	//htmlソースコードをとってくる
	html, err := page.HTML()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(html)

	/*
		//ログインしてログインボタンを押す前にエラーメッセジーが出ていたらエラー
		n := strings.Index(html, "ログインに失敗しました。")
		if n > 0 {
			log.Print("エラーメッセージがあります")
		}

		btn := page.FindByButton("ログイン")
		err = btn.Submit()
		if err != nil {
			log.Fatal(err)
		}

		//htmlソースコードをとってくる
		html, err = page.HTML()
		if err != nil {
			log.Fatal(err)
		}

		//ボタンを押したときにエラーメッセージが出ていなかったら、エラー
		n = strings.Index(html, "ログインに失敗しました。")
		if n < 0 {
			log.Print("エラーメッセージがありませんでした")

		}
	*/
}
