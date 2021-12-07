package repository

import (
    "fmt"

    driver "gorm.io/driver/mysql"
    "gorm.io/gorm"
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
    Auth *mysql.AuthRepo
    db   *gorm.DB
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
        config.Config.Database.DBType,
        config.Config.Database.Charset,
        config.Config.Database.ParseTime,
        config.Config.Database.TimeZone,
    )

    db, err := gorm.Open(driver.Open(dsn), &gorm.Config{
        NamingStrategy: schema.NamingStrategy{
            SingularTable: true,
        },
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
        Auth: mysql.NewAuthRepository(db),
        db:   db,
    }

    return MySQL
}
