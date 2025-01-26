package main

import (
	"fmt"
)

func main() {
	var (
		x1 int
		y1 int
		x2 int
		y2 int
	)
	fmt.Scan(&x1, &y1, &x2, &y2)

	if x1+2 == x2 && (y1+1 == y2 || y1-1 == y2) {
		fmt.Println("ДА")
		return
	}
	if x1-2 == x2 && (y1+1 == y2 || y1-1 == y2) {
		fmt.Println("ДА")
		return
	}
	if y1+2 == y2 && (x1+1 == x2 || x1-1 == x2) {
		fmt.Println("ДА")
		return
	}
	if y1-2 == y2 && (x1+1 == x2 || x1-1 == x2) {
		fmt.Println("ДА")
		return
	}
	fmt.Println("НЕТ")
}
