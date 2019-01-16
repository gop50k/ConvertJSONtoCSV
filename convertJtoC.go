package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"
)

// Information はJSONファイルの構造体です。
type Information struct {
	ID      int    `json:"id"`
	Subject string `json:"subject"`
	StartAt string `json:"start_at"`
}

func sub() {
	lenInfo := Information{}
	fmt.Println(lenInfo)

	leninfo := StructToMap(&lenInfo)
	fmt.Println(leninfo)
}

// StructToMap は
func StructToMap(data interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	elem := reflect.ValueOf(data).Elem()
	size := elem.NumField()

	for i := 0; i < size; i++ {
		field := elem.Type().Field(i).Name
		value := elem.Field(i).Interface()
		result[field] = value
	}

	return result
}

func main() {

	// 下で使うfor文のためにStructの要素数を取得
	// result := make(map[string]interface{})
	// elem := reflect.ValueOf(data).Elem()
	// size := elem.NumField()

	// type information []Information
	// var lenInfo information

	// fmt.Println(len(lenInfo))

	// Structのjsonタグを取得
	s := Information{}
	t := reflect.TypeOf(s)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Println(field.Tag.Get("json"))
	}

	// JSONファイル読み込み
	bytes, err := ioutil.ReadFile("./input.json")
	if err != nil {
		log.Fatal(err)
	}

	// JSONデコード
	Informations := []Information{}
	if err := json.Unmarshal(bytes, &Informations); err != nil {
		log.Fatal(err)
	}

	// JSONデコードした結果を表示
	// fmt.Println(Informations)

	//CSVファイル作成
	file, err := os.OpenFile("./output.csv", os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	// Structをmapに変換する

	// mapに変換したJSONのキーを取得して出力する

	// ヘッダーテキストを書き込む
	_, err = file.WriteString("id, subject, start_at\n")
	if err != nil {
		log.Println(err)
	}

	// デコードしたデータを表示
	for _, info := range Informations {
		fmt.Fprintln(file, info.ID, ",", info.Subject, ",", info.StartAt)
	}
	fmt.Println("処理終了")
}

// 	result := make(map[string]interface{})
// 	elem := reflect.ValueOf(data).Elem()
// 	size := elem.NumField()

// 	for i := 0; i < size; i++ {
// 	  field := elem.Type().Field(i).Tag.Get("json")
// 	  value := elem.Field(i).Interface()
// 	  result[field] = value

//   c := StructToJsonTagMap(&Informations)
//   fmt.Println(c)

// 実行コマンドショートカット
// go run convertJtoC.go
