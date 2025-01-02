package model

import "time"

type PSDFile struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"size:255;not null"`
	Data      []byte `gorm:"type:longblob;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
