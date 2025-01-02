package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/lawnmower-74/psd_uploader/domain/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB接続設定
var db *gorm.DB
var err error

func init() {
	// MySQLデータベースに接続
	dsn := "root:password@tcp(db)/psd_uploader?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
}

// RetrievePSDFileFromDB 関数は、指定したIDのPSDファイルをDBから取得し、保存します
func RetrievePSDFileFromDB(id uint, outputDir string) {
	// DBからPSDファイルを取得
	var psdFile model.PSDFile
	if err := db.First(&psdFile, id).Error; err != nil {
		log.Fatal("Failed to retrieve file from database:", err)
	}

	// outputDirディレクトリを作成
	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		log.Fatal("Failed to create output directory:", err)
	}

	// ファイルを保存するパス
	outputFilePath := fmt.Sprintf("%s/%s", outputDir, psdFile.Name)

	// 取得したバイナリデータをファイルとして保存
	err := os.WriteFile(outputFilePath, psdFile.Data, 0644)
	if err != nil {
		log.Fatal("Failed to write file:", err)
	}

	fmt.Printf("File retrieved and saved to %s\n", outputFilePath)
}

func main() {
	// コマンドライン引数を受け取り、IDと保存先ディレクトリを指定してファイルを取得
	if len(os.Args) < 3 {
		log.Fatal("Please specify the file ID and output directory.")
	}

	idStr := os.Args[1]
	id, err := strconv.ParseUint(idStr, 10, 32) // 文字列を uint 型に変換
	if err != nil {
		log.Fatal("Invalid ID format:", err)
	}
	outputDir := os.Args[2]
	fmt.Printf("ココまでは来てるよね？")
	RetrievePSDFileFromDB(uint(id), outputDir)
}
