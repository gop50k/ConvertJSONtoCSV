package main

import (
	"encoding/csv"
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
	bytes, err := ioutil.ReadFile("/Users/nakagawago/Desktop/convertJtoC/input.json")
	if err != nil {
		log.Fatal(err)
	}

	// JSONデコード
	var Informations []Information
	if err := json.Unmarshal(bytes, &Informations); err != nil {
		log.Fatal(err)
	}

	fmt.Println("id, subject, start_at")
	// デコードしたデータを表示
	for _, i := range Informations {
		fmt.Printf("%d, %s, %s\n", i.ID, i.Subject, i.StartAt)
	}

	//書き込みファイル作成
	file, err := os.Create("/Users/nakagawago/Desktop/convertJtoC/output.csv")
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	writer.Write(Informations)
	writer.Flush()

}
