package dal

import (
	"sync"
	model "tiktok-video/gen/dal/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	//gormopentracing "gorm.io/plugin/opentracing"
)

var DB *gorm.DB
var once sync.Once

func Init() {
	once.Do(func() {
		DB = ConnctDB().Debug()
		_ = DB.AutoMigrate(&model.Vedio{}, &model.User{}, &model.UserFavorite{})
	})
}

func ConnctDB() (conn *gorm.DB) {
	conn, err := gorm.Open(mysql.Open("root:123456@tcp(192.168.1.102:3306)/gorm?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// er := DB.Use(gormopentracing.New())
	// if er != nil {
	// 	panic(er)
	// }
	return conn
}
