package main

import (
	"fmt"
	"github.com/sclevine/agouti"
	"log"
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
	page.Navigate("https://viewer-bookstore.yahoo.co.jp/;_ylt=A2RjEZSvgS9cCBUAYXcFjvt7?cid=218273&u0=1&rurl=https%3A%2F%2Fbookstore.yahoo.co.jp%2Fshoshi-218273%2F")

	// keyboard操作
	// https://github.com/sclevine/agouti/issues/61
	time.Sleep(2 * time.Second)
	log.Println("space")

	//fmt.Println(page.HTML())
	fmt.Println("test")

	page.Find("body").SendKeys("\uE00D")
	time.Sleep(1 * time.Second)
	page.Find("html").SendKeys("\uE00D")
	time.Sleep(1 * time.Second)
	page.Find("css").SendKeys("\uE00D")
	time.Sleep(1 * time.Second)
	page.All("body").SendKeys("\uE00D")
	time.Sleep(1 * time.Second)
	page.All("html").SendKeys("\uE00D")
	time.Sleep(1 * time.Second)
	page.All("css").SendKeys("\uE00D")
	time.Sleep(1 * time.Second)
	page.AllByClass("pc_device").SendKeys("\uE00D")
	time.Sleep(1 * time.Second)
	page.FindByID("container").SendKeys("\uE00D")
	time.Sleep(1 * time.Second)
	page.FindByID("tagjs").SendKeys("\uE00D")
	time.Sleep(1 * time.Second)
	page.Find("body").SendKeys("\uE012")
	time.Sleep(1 * time.Second)
	page.Find("html").SendKeys("\uE012")
	time.Sleep(1 * time.Second)
	page.Find("css").SendKeys("\uE012")
	time.Sleep(1 * time.Second)
	page.All("body").SendKeys("\uE012")
	time.Sleep(1 * time.Second)
	page.All("html").SendKeys("\uE012")
	time.Sleep(1 * time.Second)
	page.All("css").SendKeys("\uE012")
	time.Sleep(1 * time.Second)
	page.AllByClass("pc_device").SendKeys("\uE012")
	time.Sleep(1 * time.Second)
	page.FindByID("container").SendKeys("\uE012")
	time.Sleep(1 * time.Second)
	page.FindByID("tagjs").SendKeys("\uE012")
	page.Find("body").Click()
	time.Sleep(1 * time.Second)
	page.Find("html").Click()
	time.Sleep(1 * time.Second)
	page.Find("css").Click()
	time.Sleep(1 * time.Second)
	page.All("body").Click()
	time.Sleep(1 * time.Second)
	page.All("html").Click()
	time.Sleep(1 * time.Second)
	page.All("css").Click()
	time.Sleep(1 * time.Second)
	page.AllByClass("pc_device").Click()
	time.Sleep(1 * time.Second)
	page.FindByID("container").Click()
	time.Sleep(1 * time.Second)
	page.FindByID("tagjs").Click()

	//page.Find("html").SendKeys("\uE00D")
	//	time.Sleep(2 * time.Second)
	//	page.Find("body").SendKeys("\uE00F") // PAGE_DOWN
	//	time.Sleep(2 * time.Second)
	//	//page.Find("html").SendKeys("\uE00F")
	//	time.Sleep(2 * time.Second)
	//	page.Find("body").SendKeys("\uE015") // DOWN
	//	time.Sleep(2 * time.Second)
	//	//page.Find("html").SendKeys("\uE015")

	//page.Size(1080, 800)
	//time.Sleep(2 * time.Second)
	//for i := 1000; i < 2050; i += 100 {
	//	page.MoveMouseBy(i, 100) // SPACE
	//	//time.Sleep(4 * time.Second)
	//	//page.FindByID("contents").Click() // SPACE
	//	page.First("contents").Click() // SPACE
	//	log.Println("click")
	//	time.Sleep(3 * time.Second)
	//	log.Println(i)
	//}
	time.Sleep(6 * time.Second)
}
