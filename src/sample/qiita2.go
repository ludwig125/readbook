// main.go
package main

import (
	"github.com/sclevine/agouti"
	"log"
)

func main() {
	driver := agouti.ChromeDriver()
	if err := driver.Start(); err != nil {
		log.Fatalf("Failed to start driver:%v", err)
	}
	defer driver.Stop()

	page, err := driver.NewPage(agouti.Browser("chrome"))
	if err != nil {
		log.Fatalf("Failed to open page:%v", err)
	}

	if err := page.Navigate("http://qiita.com/"); err != nil {
		log.Fatalf("Failed to navigate:%v", err)
	}
	log.Println("Attempt to screenshot")
	page.Screenshot("./chrome_qiita.jpg")
}
