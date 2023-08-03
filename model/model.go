package model

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	sqlDB   *gorm.DB
	redisDB *redis.Client
}

type UserData struct {
	User_id   int    `gorm:"primaryKey"`
	User_name string `json:"username"`
	Password  string `json:"password"`
	Salt      string
}

var db *DB

func init() {
	fmt.Println("Model init.")
	dsn := "host=localhost user=oneecho password=1 dbname=hustmaths port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	sDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic(err)
	}

	rClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	db = &DB{sqlDB: sDB, redisDB: rClient}
}

func GetDBInstance() *DB {
	return db
}

func (db *DB) InsertUserData(uData *UserData) error {
	res := db.sqlDB.Table("users").Create(uData)
	if res.Error != nil {
		log.Println(res.Error)
		// 未知错误
		return res.Error
	}
	return nil
}

// Return uName在数据库内的数量，true为有大于0个，false为0个
func (db *DB) CheckUserNameNum(uName string) bool {
	rows := db.sqlDB.Table("users").Where("user_name=?", uName).Find(&UserData{}).RowsAffected
	if rows > 0 {
		return true
	} else {
		return false
	}
}

func (db *DB) GetUserDataByUserName(uName string, userData *UserData) {
	err := db.sqlDB.Table("users").Where("user_name=?", uName).Find(userData).Error
	if err != nil {
		log.Println(err)
	}
}

func (db *DB) Get(key string) string {
	val, err := db.redisDB.Get(context.Background(), key).Result()
	if err != nil {
		log.Println(err)
	}
	return val
}

func (db *DB) Set(key string, value string, exp time.Duration) {
	err := db.redisDB.Set(context.Background(), key, value, exp).Err()
	if err != nil {
		log.Panic(err)
	}
}

func (db *DB) Remove(key string) {
	err := db.redisDB.Del(context.Background(), key).Err()
	if err != nil {
		log.Panic(err)
	}
}
