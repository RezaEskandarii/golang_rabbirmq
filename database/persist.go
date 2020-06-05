package database

import "../models"

func Persist(entity interface{}) {
	db, _ := GetConnection()
	defer db.Close()
	db.Model(entity).Save(entity)
}

// migrate models
func init() {
	db, _ := GetConnection()
	defer db.Close()
	db.AutoMigrate(&models.Message{})
}
