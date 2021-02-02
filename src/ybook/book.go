package main

import (
	"flag"
	"fmt"
	"github.com/sclevine/agouti"
	"log"
	"os"
	"reflect"
	"strings"
	"syscall"
	"time"

	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	// ユーザデータ読み取り
	flag.Parse()
	if flag.Arg(0) == "" {
		log.Fatal("Failed to get username")
		os.Exit(0)
	}
	userName := flag.Arg(0)

	fmt.Print("Password: ")
	userPassword, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		log.Fatal("Failed to read password", err)
		os.Exit(0)
	}

	// ブラウザはChromeを指定して起動
	//driver := agouti.ChromeDriver(agouti.Browser("chrome"))
	driver := agouti.ChromeDriver(
		agouti.ChromeOptions("args", []string{
			//		"--headless",             // headlessモードの指定
			"--window-size=1280,800", // ウィンドウサイズの指定
		}),
		agouti.Debug,
	)
	if err := driver.Start(); err != nil {
		log.Fatalf("Failed to start driver: %v", err)
		os.Exit(0)
	}
	defer driver.Stop()

	page, err := driver.NewPage()
	if err != nil {
		log.Fatalf("Failed to open page: %v", err)
		os.Exit(0)
	}

	// password入力がうまく行かず失敗することがあるのでretryする
	MAX_RETRY := 3
	for i := 1; i <= MAX_RETRY; i++ {
		//res := login(page, userName, userPassword)
		//fmt.Printf("login res: %v\n", res)
		//time.Sleep(4 * time.Second)
		if login(page, userName, userPassword) {
			break
		}
		log.Printf("Failed to input password %d times", i)
		if i == MAX_RETRY {
			log.Fatalf("Failed to input password %d times", i)
			os.Exit(0)
		}
		// 再接続の際に時間を置く
		time.Sleep(2 * time.Second)
	}

	time.Sleep(1 * time.Second)
	if err := page.FindByID("btnSubmit").Submit(); err != nil {
		log.Fatalf("Failed to confirm password: %v", err)
	}

	time.Sleep(2 * time.Second)

	sessionBus := reflect.ValueOf(page.Session().Bus)
	sessionURL := reflect.Indirect(sessionBus)
	sessionField := sessionURL.FieldByName(`SessionURL`)
	sessionString := sessionField.String()
	sessionSplit := strings.SplitN(sessionString, "/", 7)
	sessionID := sessionSplit[len(sessionSplit)-1]
	fmt.Println(sessionID)

	// 目的のページに遷移
	//bookurl := fmt.Sprintf("https://ebookjapan.yahoo.co.jp/viewer?sessionid=%s&bookid=B00060176851&reflow=false&titleid=193760&type=purchased", sessionID)
	bookurl := "https://viewer-bookstore.yahoo.co.jp/?cid=218273&u0=1&rurl=http%3A%2F%2Fbookstore.yahoo.co.jp%2Fshelf.html%23ncrumb%3DblxIOqBtT9P"
	//bookurl := "https://ebookjapan.yahoo.co.jp/bookshelf"
	if err := page.Navigate(bookurl); err != nil {
		log.Fatalf("Failed to navigate bookstore page: %v", err)
	}

	time.Sleep(3 * time.Second)
	//log.Println("Attempt to screenshot")
	//if err := page.Screenshot("page.png"); err != nil {
	//	log.Fatalf("Failed to screenshot page: %v", err)
	//}
	//log.Println("Secceded to screenshot")
	log.Println("try to aaa")
	if err := page.Find("body").SendKeys("\uE00D"); err != nil {
		log.Fatalf("faled to aaa")
	}
	time.Sleep(2 * time.Second)
	//page.Find("body").SendKeys("\uE012")
	//time.Sleep(2 * time.Second)
	//page.Find("html").SendKeys("\uE012")
	//time.Sleep(2 * time.Second)
	//page.Find("body").SendKeys("\uE00D") // SPACE
	//time.Sleep(2 * time.Second)
	//page.Find("html").SendKeys("\uE00D") // SPACE
	//time.Sleep(2 * time.Second)
	//page.Find("img").SendKeys("\uE012")
	//time.Sleep(2 * time.Second)
	//page.Find("img").SendKeys("\uE00D") // SPACE
	//time.Sleep(2 * time.Second)
	//page.Find("div").SendKeys("\uE012")
	//time.Sleep(2 * time.Second)
	//page.Find("div").SendKeys("\uE00D") // SPACE
	//time.Sleep(2 * time.Second)
	//page.FindByID("contents").SendKeys("\uE00D") // SPACE
	//time.Sleep(2 * time.Second)
	//page.FindByClass("pages").SendKeys("\uE00D") // SPACE
	//time.Sleep(2 * time.Second)

	// これで何故か拡大
	//	page.Size(1000, 800)
	//	for i := 0; i < 1000; i += 100 {
	//		page.MoveMouseBy(i, 100)          // SPACE
	//		page.FindByID("contents").Click() // SPACE
	//		time.Sleep(2 * time.Second)
	//		log.Println(i)
	//	}
	//	time.Sleep(2 * time.Second)
	//	page.FindByID("contents").Click() // SPACE

	page.Size(1000, 800)
	time.Sleep(2 * time.Second)
	page.FindByID("contents").SendKeys("\uE00D") // SPACE
	log.Println("space")
	page.FindByID("contents").ScrollFinger(100, 0)
	time.Sleep(2 * time.Second)
	page.FindByClass("pages").Click()
	page.FindByClass("pt-img").Click() // SPACE
	page.FindByClass("pages").ScrollFinger(400, 0)
	page.FindByClass("pt-img").ScrollFinger(400, 0) // SPACE
	for i := 0; i < 950; i += 100 {
		page.MoveMouseBy(i, 10) // SPACE
		page.FindByID("contents").ScrollFinger(i, 0)
		page.FindByID("contents").Click() // SPACE
		page.FindByID("content_base").Click()
		page.FindByID("content").Click() // SPACE
		page.FindByID("content_base").ScrollFinger(i, 0)
		page.FindByID("content").ScrollFinger(i, 0) // SPACE
		time.Sleep(3 * time.Second)
		log.Println(i)
	}
	time.Sleep(2 * time.Second)
	page.FindByID("contents").Click() // SPACE

	//time.Sleep(2 * time.Second)
	//page.FindByClass("pages").Click() // SPACE
	//time.Sleep(2 * time.Second)
	//page.Find("contents").Click() // SPACE
	//time.Sleep(2 * time.Second)
	//page.Find("pages").Click() // SPACE
	log.Println("success to aaa")

	//for i := 1; i < 2; i++ {
	//	pnum := fmt.Sprintf("~/Pictures/worldtrigger1/%d.jpg", i)
	//	page.Screenshot(pnum)

	//}
	// 処理完了後、3秒間ブラウザを表示しておく
	time.Sleep(6 * time.Second)
}

func login(page *agouti.Page, userName string, userPassword []byte) bool {
	// ログインページに遷移
	if err := page.Navigate("https://login.yahoo.co.jp/config/login?.src=www&.done=https://www.yahoo.co.jp"); err != nil {
		log.Fatalf("Failed to navigate: %v", err)
		os.Exit(0)
	}
	// IDの要素を取得し、値を設定
	identity := page.FindByID("username")
	identity.Fill(userName)
	// 速すぎると機械だと疑われて文字認証画面に飛ばされるので待つ
	time.Sleep(2 * time.Second)
	// formをサブミット
	if err := page.FindByID("btnNext").Click(); err != nil {
		log.Fatalf("Failed to confirm username: %v", err)
		os.Exit(0)
	}

	// ページ更新後のFindが追いつかない可能性があるので少し待つ
	time.Sleep(1 * time.Second)
	// Passの要素を取得し、値を設定
	//password := page.FindByID("passwd")
	password := page.FindByName("passwd")
	//fmt.Printf("pass %v\n", password)
	password.Fill(string(userPassword))

	// Findがうまくできていないとpasswordが入力されないことがある
	// Visibleでpasswordが入力されているのか確認する
	// 入力されているとtrue, 空のままだとfalse

	// https://godoc.org/github.com/sclevine/agouti#Selection.Visible
	// Visible returns true if all of the elements that the selection refers to are visible.
	// passwordがきちんと入力されているとbtnClearShowがCSSに現れる
	// btnClearShowがないとVisibleの判定でfalseになるっぽい
	// 正直判定基準がわからない
	password.FindByClass("btnClearShow")
	inputRes, err := password.Visible()
	if err != nil {
		log.Fatalf("Failed to get password visible: %v", err)
	}
	//fmt.Printf("input result %v\n", inputRes)
	return inputRes
}
