package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/lawnmower-74/psd_uploader/model"
	"github.com/lawnmower-74/psd_uploader/db"
)


func SavePSDFileToDB(psdFilePath string) error {
	db := db.ConnectDB()

	fileName := filepath.Base(psdFilePath)
	fileData, err := os.ReadFile(psdFilePath)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}
	
	psdFile := model.PSDFile{
		Name: fileName,
		Data: fileData,
	}

	if err := db.Create(&psdFile).Error; err != nil {
		return fmt.Errorf("failed to save file to database: %w", err)
	}

	return nil
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s <path_to_psd_file>", os.Args[0])
	}

	psdFilePath := os.Args[1]
	
	err := SavePSDFileToDB(psdFilePath)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Printf("File '%s' saved to database successfully.\n", psdFilePath)
}
