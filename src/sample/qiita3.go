package main

import (
	"github.com/sclevine/agouti"
	"log"
)

func main() {
	// Chromeを利用することを宣言
	//agoutiDriver := agouti.ChromeDriver()
	agoutiDriver := agouti.ChromeDriver(agouti.ChromeOptions("args", []string{"--headless", "--disable-gpu", "--no-sandbox"}))
	agoutiDriver.Start()
	defer agoutiDriver.Stop()
	page, _ := agoutiDriver.NewPage()

	// 自動操作
	page.Navigate("https://qiita.com/")
	log.Println(page.Title())
	page.Screenshot("Screenshot01.png")

	page.FindByLink("もっと詳しく").Click()
	log.Println(page.Title())
	page.Screenshot("Screenshot02.png")
}
