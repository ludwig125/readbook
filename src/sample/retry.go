package main

import (
	"fmt"
)

func main() {
	var count int
	for i := 1; i < 4; i++ {
		count += 1
		if fail2times(count) {
			fmt.Println("res: true")
			break
		}
		fmt.Println("res: false")
	}
}

func fail2times(count int) bool {
	fmt.Println(count)
	if count < 3 {
		return false
	}
	return true
}
