package model

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/RanchoCooper/go-programming-tour-book/blog-service/configs"
	"github.com/RanchoCooper/go-programming-tour-book/blog-service/global"
)

type Model struct {
	ID         uint32 `gorm:"primary_key" json:"id"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	CreatedOn  uint32 `json:"created_on"`
	ModifiedOn uint32 `json:"modified_on"`
	DeletedOn  uint32 `json:"deleted_on"`
	IsDel      uint8  `json:"id_del"`
}

func NewDBEngine(databaseSetting configs.DatabaseSettingS) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		databaseSetting.UserName,
		databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.DBName,
		databaseSetting.Charset,
		databaseSetting.ParseTime,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetConnMaxIdleTime(time.Duration(global.DatabaseSetting.MaxIdleConns))
	sqlDB.SetMaxOpenConns(global.DatabaseSetting.MaxOpenConns)
	return db, nil
}

func updateTimeStampForCreateCallback(db *gorm.DB) *gorm.DB {
	// FIXME
	return db
}
