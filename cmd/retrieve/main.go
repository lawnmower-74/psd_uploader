package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/lawnmower-74/psd_uploader/model"
	"github.com/lawnmower-74/psd_uploader/db"
)

func RetrievePSDFileFromDB(id uint, outputDir string) {
	db := db.ConnectDB()

	var psdFile model.PSDFile
	if err := db.First(&psdFile, id).Error; err != nil {
		log.Fatal("Failed to retrieve file from database:", err)
	}

	// ダウンロード先のディレクトリがない場合に備えて
	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		log.Fatal("Failed to create output directory:", err)
	}

	outputFilePath := fmt.Sprintf("%s/%s", outputDir, psdFile.Name)

	err := os.WriteFile(outputFilePath, psdFile.Data, 0644)
	if err != nil {
		log.Fatal("Failed to write file:", err)
	}

	fmt.Printf("File retrieved and saved to %s\n", outputFilePath)
}

func main() {
	if len(os.Args) < 3 {
		log.Fatal("Please specify the file ID and output directory.")
	}

	idStr := os.Args[1]
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		log.Fatal("Invalid ID format:", err)
	}

	outputDir := os.Args[2]

	RetrievePSDFileFromDB(uint(id), outputDir)
}
