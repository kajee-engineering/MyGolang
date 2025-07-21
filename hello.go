package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

// ジェネリクス
func Sum[T int | float32 | float64](nums []T) T {

	return 0
}

func main() {
	fmt.Println("Hello World")

	//// 構造体
	//// この構造体は型であり。メモリに存在するものではない。メモリ上に「インスタンス」を作らなければデータを保存したりできない。
	//type Book struct {
	//	// 大文字で始まる名前にしなければ、外部のパッケージから利用できない。
	//	// 構造体は中にメソッドを定義できない。外にメソッドの定義はできる。
	//	Title      string // イミュータブルなので、変更したい場合は新しくstringを作成する必要がある。
	//	Author     string
	//	Publisher  string
	//	ReleasedAt time.Time // 構造体。
	//}
	//
	//// インスタンス作成(フィールドは全てゼロ値に初期化)
	//var b Book
	//// フィールドを初期化しながらインスタンスを作成
	//b2 := Book{
	//	Title: "Twisted Network Programming Essentials",
	//}
	//
	//// フィールドを初期化しながらインスタンスを作成
	//// 変数にはポインターを格納
	//b3 := &Book{
	//	Title: "The Go Programming Language",
	//}

	type Book struct {
		Title      string    `json:"title"`
		Author     string    `json:"author"`
		Publisher  string    `json:"publisher"`
		ReleasedAt time.Time `json:"released_at"`
		ISBN       string    `json:"isbn"`
	}

	// book.jsonファイルの中身を構造体にロードするコードが次のコードです。
	f, err := os.Open("book.json")
	if err != nil {
		log.Fatal("file open error: ", err)
	}
	d := json.NewDecoder(f)
	var b Book
	d.Decode(&b)
	fmt.Println(b)
}
