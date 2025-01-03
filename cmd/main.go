package main

import (
	"log"

	"github.com/lawnmower-74/psd_uploader/db"
	"github.com/lawnmower-74/psd_uploader/model"
)

func main()  {
	db := db.ConnectDB()
	if err := db.AutoMigrate(&model.PSDFile{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}
}