package main

import "fmt"

// インタフェースを定義
type Talker interface {
	Talk()
}

// 構造体を定義
type Greeter struct {
	name string
}

// 構造体はTalkerインタフェースで定義されているメソッドを持っている
// funcキーワードとメソッドシグネチャの間にレシーバ((g Greeter))をおくと，構造体にメソッドを定義したことになる
func (g Greeter) Talk() {
	fmt.Println("Hello, my name is %s\n", g.name)
}

func main() {
	// インタフェースの型を持つ変数を宣言
	var talker Talker
	// インタフェースを満たす構造体のポインタを代入できる
	talker = &Greeter{"stoshiya"}
	talker.Talk()
}