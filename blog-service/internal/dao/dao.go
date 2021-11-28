package dao

import (
    "gorm.io/gorm"
)

/**
 * @author Rancho
 * @date 2021/11/28
 */

type Dao struct {
    engine *gorm.DB
}

func New(engine *gorm.DB) *Dao {
    return &Dao{engine: engine}
}