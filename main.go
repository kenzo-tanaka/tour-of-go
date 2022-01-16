package main

import "fmt"

func main() {
	i := 5
	pointer := &i
	fmt.Println(*pointer)

	// ポインタを操作することで
	// ポインタが指す変数の値を直接変更することができる
	*pointer = 21
	fmt.Println(i)
}
