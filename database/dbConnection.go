package database

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//数据库连接，获取数据库连接就可以操作数据库表
var Conn *gorm.DB

//数据库连接串，这里屏蔽掉生产的数据库
var testDBName = "host=localhost port=5432 user=joe.zhang dbname=mydb password=19950209 sslmode=disable"
var prodDBName = os.Getenv("DATABASE_URL")
var DBName = prodDBName

var DBEngine = "postgres"

func GetDBConnection() (*gorm.DB, error) {
	db, err := gorm.Open(DBEngine, DBName)
	if err != nil {
		panic("Failed to connect to database " + err.Error())
	}
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(10)

	Conn = db
	return db, err
}
