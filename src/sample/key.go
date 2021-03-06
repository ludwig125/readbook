package main

import (
	"github.com/sclevine/agouti"
	//"log"
	"time"
)

func main() {
	// Chromeを利用することを宣言
	//agoutiDriver := agouti.ChromeDriver()
	//agoutiDriver := agouti.ChromeDriver(agouti.ChromeOptions("args", []string{"--headless", "--disable-gpu", "--no-sandbox"}))
	agoutiDriver := agouti.ChromeDriver()
	agoutiDriver.Start()
	defer agoutiDriver.Stop()
	page, _ := agoutiDriver.NewPage()

	// 自動操作
	page.Navigate("https://www.yahoo.co.jp/")

	// keyboard操作
	// https://github.com/sclevine/agouti/issues/61
	time.Sleep(2 * time.Second)
	page.Find("body").SendKeys("\uE00D") // SPACE
	time.Sleep(2 * time.Second)
	//page.Find("html").SendKeys("\uE00D")
	time.Sleep(2 * time.Second)
	page.Find("body").SendKeys("\uE00F") // PAGE_DOWN
	time.Sleep(2 * time.Second)
	//page.Find("html").SendKeys("\uE00F")
	time.Sleep(2 * time.Second)
	page.Find("body").SendKeys("\uE015") // DOWN
	time.Sleep(2 * time.Second)
	//page.Find("html").SendKeys("\uE015")

	time.Sleep(6 * time.Second)
}
