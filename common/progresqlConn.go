package common

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Database struct {
	*gorm.DB
}

var DB *gorm.DB

// Opening a database and save the reference to `Database` struct.
func Init() *gorm.DB {
	//db, err := gorm.Open("postgres", "host=localhost dbschema=ledger_hub user=ledger_hub dbname=ledger_hub sslmode=disable password=ledger_hub")
	//db, err := gorm.Open("postgres", "host="+config.GetConfig().Database.DbHost+" port="+config.GetConfig().Database.DbPort+" user="+config.GetConfig().Database.DbUsername+" dbname="+config.GetConfig().Database.DbSchemaName+" sslmode=disable password="+config.GetConfig().Database.DbPwd)
	db, err := gorm.Open("postgres", "host=localhost user=mike dbname=postgres sslmode=disable password=a123456")
	//defer DB.Close()

	if err != nil {
		fmt.Println(err.Error())
	}
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.LogMode(true)
	DB = db
	return DB
}

// Using this function to get a connection, you can create your connection pool here.
func GetDB() *gorm.DB {
	return DB
}

func SaveOne(data interface{}) error {
	db := GetDB()
	err := db.Save(data).Error
	return err
}
