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
	var isValid bool

	if x1+2 == x2 && (y1+1 == y2 || y1-1 == y2) {
		isValid = true
	}
	if x1-2 == x2 && (y1+1 == y2 || y1-1 == y2) {
		isValid = true
	}
	if y1+2 == y2 && (x1+1 == x2 || x1-1 == x2) {
		isValid = true
	}
	if y1-2 == y2 && (x1+1 == x2 || x1-1 == x2) {
		isValid = true
	}
	if isValid {
		fmt.Println("ДА")
	} else {
		fmt.Println("НЕТ")
	}
}
