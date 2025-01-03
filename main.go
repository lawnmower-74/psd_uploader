package main

import (
	"fmt"

	"github.com/lawnmower-74/psd_uploader/db"
	"github.com/lawnmower-74/psd_uploader/model"
)

func main()  {
	db := db.ConnectDB()
	if err := db.AutoMigrate(&model.PSDFile{}); err != nil {
		fmt.Errorf("failed to migrate database: %w", err)
	}
}