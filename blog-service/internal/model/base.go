package model

import (
    "fmt"

    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
    "gorm.io/gorm/schema"

    "go-programming-tour-book/blog-service/global"
    "go-programming-tour-book/blog-service/pkg/setting"
)

/**
 * @author Rancho
 * @date 2021/11/26
 */

type Model struct {
    ID         uint32 `gorm:"primary_key" json:"id"`
    CreatedBy  string `json:"created_by"`
    ModifiedBy string `json:"modified_by"`
    CreatedOn  uint32 `json:"created_on"`
    ModifiedOn uint32 `json:"modified_on"`
    DeletedOn  uint32 `json:"deleted_on"`
    IsDel      uint8  `json:"is_del"`
}

func NewDBEngine(databaseSetting *setting.DatabaseSettingS) (*gorm.DB, error) {
    dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
        databaseSetting.UserName,
        databaseSetting.Password,
        databaseSetting.Host,
        databaseSetting.DBName,
        databaseSetting.Charset,
        databaseSetting.ParseTime,
    )
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
        NamingStrategy: schema.NamingStrategy{
            SingularTable: true,
        },
    })
    if err != nil {
        return nil, err
    }

    if global.ServerSetting.RunMode == "debug" {
        db.Logger.LogMode(logger.Info)
    }
    sqlDB, err := db.DB()
    if err != nil {
        return nil, err
    }
    sqlDB.SetMaxIdleConns(databaseSetting.MaxIdleConns)
    sqlDB.SetMaxOpenConns(databaseSetting.MaxOpenConns)

    return db, nil
}
