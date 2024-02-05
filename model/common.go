package model

import (
	"database/sql/driver"
	"encoding/json"
	"log"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type TableType interface {
	TableName() string
}

type DsnConfig struct {
	UserName string
	PassWord string
	Ip       string
	Database string
}

func getConfig() DsnConfig {
	viper.SetConfigFile("./conf/config.json")
	err := viper.ReadInConfig()
	if err != nil {
		panic("config not found")
	}
	res := DsnConfig{
		UserName: viper.GetString("userName"),
		PassWord: viper.GetString("passWord"),
		Ip:       viper.GetString("ip"),
		Database: viper.GetString("database"),
	}
	return res
}

func GetDsnDefault() string {
	conf := getConfig()
	dsn := conf.UserName + ":" + conf.PassWord + "@tcp(" + conf.Ip + ")/" + conf.Database + "?charset=utf8mb4&timeout=30s"
	return dsn
}

func GetDsn(database string) string {
	conf := getConfig()
	dsn := conf.UserName + ":" + conf.PassWord + "@tcp(" + conf.Ip + ")/" + database + "?charset=utf8mb4&timeout=30s"
	return dsn
}

func NewDB(dsn string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	sqlDb, _ := db.DB()
	sqlDb.SetMaxIdleConns(10)
	sqlDb.SetMaxOpenConns(50)
	sqlDb.SetConnMaxLifetime(time.Hour)
	return db
}

func CreateTable(tableType TableType, db *gorm.DB) {
	if db.Migrator().HasTable(tableType) {
		log.Println(tableType.TableName() + " alreadly create")
	} else {
		err := db.AutoMigrate(tableType)
		log.Println(tableType.TableName()+" create success", err)
	}
}

func CompoundScan(val interface{}, data interface{}) error {
	err := json.Unmarshal(val.([]byte), data)
	return err
}

func CompoundValue(data interface{}) (driver.Value, error) {
	str, err := json.Marshal(data)
	return str, err
}
