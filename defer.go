package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("hello world")

	// プログラムの中ではとある処理の後に後片付けが必要な場合がある。
	// Pythonのwith句のようにさまざまな言語が「後処理を実行」のメカニズムを用意している。
	// Golangも準備と同じ場所で予約するのが1番確実だ。という原則は変わらない。距離が離れると、整合が取れていない場合の見落としが増える。

	// ファイルを開く
	f, err := os.Create("sample.txt")
	if err != nil {
		fmt.Println("err", err)
		return
	}
	// この関数のスコープを抜けたら自動でファイルをクローズ
	defer f.Close()
	io.WriteString(f, "hello world")
}
