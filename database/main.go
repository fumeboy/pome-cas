package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"good/build"
	"log"
	"time"
)

var database *gorm.DB

type ModelCom struct {
	ID int `gorm:"primary_key"`

	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`
}

func init() {
	var err error
	dbInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		build.DB_HOST,
		build.DB_PORT,
		build.DB_USER,
		build.DB_PASSWORD,
		build.DB_NAME,
	)
	database, err = gorm.Open(build.DB, dbInfo)
	if err != nil {
		log.Panic("连接数据库不成功", err, dbInfo)
	}
	database.LogMode(true)
	// 全局禁用表名复数
	database.SingularTable(true)
	database.DB().SetMaxIdleConns(10)
	database.DB().SetMaxOpenConns(100)

	database.AutoMigrate(
		&Comment{},
		&Relation{},
		&TagLikes{},
		&Tags{},
		&Topic{},
		&TopicLikes{},
		&User{},
		&Test{},
	)
}

func CloseDB() {
	defer database.Close()
}
