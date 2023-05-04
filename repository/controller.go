package repository

import (
	"log"

	"github.com/yusufwib/arvigo-backend/pkg/database"
	"gorm.io/gorm"
)

func Database() (db *gorm.DB) {
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal("Error connecting database")
		return
	}
	return
}
