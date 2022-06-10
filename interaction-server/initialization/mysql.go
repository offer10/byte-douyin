package initialization

import (
	"fmt"
	"log"

	"github.com/offer10/byte-douyin/basic-server/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func RegisterMySQL() {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%s&loc=%s",
		//"root", "1234", "127.0.0.1", 3306, "byte-douyin", "utf8mb4", "true", "Local",
		"root", "1234", "127.0.0.1", 3306, "byte-douyin", "utf8mb4", "true", "Local",
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panicf("初始化 MySQL 异常: %v", err)
	}
	conf.MySQL = db
}
