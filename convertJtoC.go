package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strings"
)

// Information はjsonファイルの構造体
type Information struct {
	ID      int    `json:"id"`
	Subject string `json:"subject"`
	StartAt string `json:"start_at"`
}

func main() {
	// Structのjsonタグを取得
	s := Information{}
	t := reflect.TypeOf(s)

	// 取得したjsonタグを格納するスライスを宣言
	var lenInfo []string

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		j := field.Tag.Get("json")
		lenInfo = append(lenInfo, j)
	}

	// jsonファイル読み込み
	bytes, err := ioutil.ReadFile("./input.json")
	if err != nil {
		log.Fatal(err)
	}

	// jsonファイルデコード
	Informations := []Information{}
	if err := json.Unmarshal(bytes, &Informations); err != nil {
		log.Fatal(err)
	}

	//CSVファイル作成
	file, err := os.OpenFile("./output.csv", os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	// スライスをカンマ区切りの文字列に変換
	strInfo := strings.Join(lenInfo, ", ")

	// ヘッダー書き込み
	_, err = file.WriteString(strInfo)
	_, err = file.WriteString("\n")
	if err != nil {
		log.Println(err)
	}

	// デコードしたデータを表示
	for _, info := range Informations {
		fmt.Fprintln(file, info.ID, ",", info.Subject, ",", info.StartAt)
	}
}
