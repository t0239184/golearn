package database

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/t0239184/golearn/internal/model"
)

type GormDatabase struct {
	DB *gorm.DB
}

func InitDatabase() *GormDatabase {
	const (
		UserName     string = "admin"
		Password     string = "password"
		Addr         string = "127.0.0.1"
		Port         int    = 3306
		Database     string = "golearn"
		MaxLifetime  int    = 10
		MaxOpenConns int    = 10
		MaxIdleConns int    = 10
	)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True", UserName, Password, Addr, Port, Database)
	gormConfig := &gorm.Config{}
	db, err := New(dsn, gormConfig)
	if err != nil {
		logrus.Fatalf("[main] database.New failed: %v", err)
	}
	db.AutoMigrate()
	return db
}

func New(dsn string, gormConfig *gorm.Config) (*GormDatabase, error) {
	db, err := gorm.Open(mysql.Open(dsn), gormConfig)
	if err != nil {
		return nil, err
	}

	return &GormDatabase{DB: db}, nil
}

func (d *GormDatabase) AutoMigrate() {

	if err := d.DB.AutoMigrate(
		&model.User{},
	); err != nil {
		logrus.Fatal(err.Error())
	}
}

func (d *GormDatabase) DropAllTables() {
	if err := d.DB.Migrator().DropTable(
		"users",
	); err != nil {
		logrus.Fatal(err.Error())
	}
}
