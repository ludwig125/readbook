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
	fmt.Println("test")

	err := page.All("").SendKeys("\uE00D")
	if err != nil {
		log.Println(err)
	}
	err = page.All("img").SendKeys("\uE00D")
	if err != nil {
		log.Println(err)
	}
	err = page.All("div").SendKeys("\uE00D")
	if err != nil {
		log.Println(err)
	}
	err = page.All("id").SendKeys("\uE00D")
	if err != nil {
		log.Println(err)
	}
	err = page.All("css").SendKeys("\uE00D")
	if err != nil {
		log.Println(err)
	}
	err = page.All("CSS").SendKeys("\uE00D")
	if err != nil {
		log.Println(err)
	}
	err = page.All("html").SendKeys("\uE00D")
	if err != nil {
		log.Println(err)
	}
	err = page.All("hr").SendKeys("\uE00D")
	if err != nil {
		log.Println(err)
	}
	err = page.AllByName("").SendKeys("\uE00D")
	if err != nil {
		log.Println(err)
	}
	err = page.AllByName("img").SendKeys("\uE00D")
	if err != nil {
		log.Println(err)
	}
	err = page.AllByName("div").SendKeys("\uE00D")
	if err != nil {
		log.Println(err)
	}
	err = page.AllByName("id").SendKeys("\uE00D")
	if err != nil {
		log.Println(err)
	}
	err = page.AllByName("css").SendKeys("\uE00D")
	if err != nil {
		log.Println(err)
	}
	err = page.AllByName("CSS").SendKeys("\uE00D")
	if err != nil {
		log.Println(err)
	}
	err = page.AllByName("html").SendKeys("\uE00D")
	if err != nil {
		log.Println(err)
	}
	err = page.AllByName("hr").SendKeys("\uE00D")
	if err != nil {
		log.Println(err)
	}

	//page.MoveMouseBy(10, 100)
	//page.Find("").Click()
	//page.FindByName("").Click()
	//page.First("").Click()
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
