package main

import "fmt"

type Vertext struct {
	X int
	Y int
}

func main() {
	i := 5
	pointer := &i
	fmt.Println(*pointer)

	// ポインタを操作することで
	// ポインタが指す変数の値を直接変更することができる
	*pointer = 21
	fmt.Println(i)

	// https://go-tour-jp.appspot.com/moretypes/2
	fmt.Println(Vertext{1, 2})

	// https://go-tour-jp.appspot.com/moretypes/4
	v := Vertext{X: 1}
	p := &v
	p.X = 1e9
	fmt.Println(v)

	// https://go-tour-jp.appspot.com/moretypes/6
	// 配列は定義時に長さを設定する
	var array [2]string
	array[0] = "Kenzo"
	array[1] = "Tanaka"
	fmt.Println(array)
}
