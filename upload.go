package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/lawnmower-74/psd_uploader/domain/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


func SavePSDFileToDB(psdFilePath string) error {
	// PSDファイルを読み込む
	fileData, err := os.ReadFile(psdFilePath) // ioutil.ReadFileは非推奨
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	// ファイル名だけを取得
	fileName := filepath.Base(psdFilePath)

	// データベース接続情報
	dsn := "root:password@tcp(db)/psd_uploader?charset=utf8mb4&parseTime=True&loc=Local"

	// データベース接続
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	// テーブルのマイグレーション
	if err := db.AutoMigrate(&model.PSDFile{}); err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}

	// ファイル情報を保存
	psdFile := model.PSDFile{
		Name: fileName,
		Data: fileData,
	}

	// データベースに保存
	if err := db.Create(&psdFile).Error; err != nil {
		return fmt.Errorf("failed to save file to database: %w", err)
	}

	return nil
}

func main() {
	// コマンドライン引数の確認
	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s <path_to_psd_file>", os.Args[0])
	}

	// PSDファイルのパスを取得
	psdFilePath := os.Args[1]

	// PSDファイルをDBに保存
	err := SavePSDFileToDB(psdFilePath)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Printf("File '%s' saved to database successfully.\n", psdFilePath)
}
