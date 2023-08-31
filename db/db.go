package db

import (
	"fmt"
	"gotest/config"
	"gotest/types"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MariaDBDataStore struct {
	db *gorm.DB
}

func autoMigrateStructs(db *gorm.DB) {
	db.AutoMigrate(&types.Grocery{})
	db.AutoMigrate(&types.Role{})
	db.AutoMigrate(&types.User{})
	db.AutoMigrate(&types.UserGrocery{})
}

func NewMariaDBDataStore(c config.MariaDBConfig) *MariaDBDataStore {
	ds := MariaDBDataStore{}
	var err error
	ds.db, err = gorm.Open(mysql.Open(c.DSN()))
	autoMigrateStructs(ds.db)
	//autoMigrateStructs(ds.db)
	if err != nil {
		fmt.Errorf("Connection failed %v", err)
	}
	return &ds
}
