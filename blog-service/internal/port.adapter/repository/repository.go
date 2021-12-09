package repository

import (
    "fmt"

    driver "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
    "gorm.io/gorm/schema"

    "go-programming-tour-book/blog-service/config"
    "go-programming-tour-book/blog-service/internal/port.adapter/repository/mysql"
)

/**
 * @author Rancho
 * @date 2021/12/8
 */

var MySQL *MySQLRepository

type MySQLRepository struct {
    db   *gorm.DB
    Auth *mysql.AuthRepo
    Tag  *mysql.TagRepo
}

func init() {
    if MySQL == nil {
        MySQL = NewMySQLRepository()
    }
}

func NewMySQLRepository() *MySQLRepository {
    dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=%s",
        config.Config.Database.UserName,
        config.Config.Database.Password,
        config.Config.Database.Host,
        config.Config.Database.DBName,
        config.Config.Database.Charset,
        config.Config.Database.ParseTime,
        config.Config.Database.TimeZone,
    )

    db, err := gorm.Open(driver.Open(dsn), &gorm.Config{
        NamingStrategy: schema.NamingStrategy{
            SingularTable: true,
        },
        Logger: logger.Default.LogMode(logger.Info),
    })
    if err != nil {
        panic("init DB fail, err: " + err.Error())
    }

    sqlDB, err := db.DB()
    if err != nil {
        panic("get sqlDB fail, err: " + err.Error())
    }
    sqlDB.SetMaxIdleConns(config.Config.Database.MaxIdleConns)
    sqlDB.SetMaxOpenConns(config.Config.Database.MaxOpenConns)

    MySQL = &MySQLRepository{
        db:   db,
        Auth: mysql.NewAuthRepository(db),
        Tag:  mysql.NewTagRepository(db),
    }

    return MySQL
}
