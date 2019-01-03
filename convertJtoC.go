package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// Information はJSONファイルの構造体です。
type Information struct {
	ID      int    `json:"id"`
	Subject string `json:"subject"`
	StartAt string `json:"start_at"`
}

func main() {
	// JSONファイル読み込み
	bytes, err := ioutil.ReadFile("./input.json")
	if err != nil {
		log.Fatal(err)
	}

	// JSONデコード
	var Informations []Information
	if err := json.Unmarshal(bytes, &Informations); err != nil {
		log.Fatal(err)
	}

	//CSVファイル作成
	file, err := os.OpenFile("./output.csv", os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	// ヘッダーテキストを書き込む
	_, err = file.WriteString("id, subject, start_at\n")
	if err != nil {
		log.Println(err)
	}

	// デコードしたデータを表示
	for _, info := range Informations {
		fmt.Fprintln(file, info.ID, ",", info.Subject, ",", info.StartAt)
	}
}
