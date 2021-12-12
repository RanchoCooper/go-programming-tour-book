package tag

import (
    "time"

    "gorm.io/gorm"
)

/**
 * @author Rancho
 * @date 2021/11/26
 */

type Tag struct {
    ID        uint   `gorm:"primarykey"`
    Name      string `json:"name"`
    State     *uint8 `json:"state"`
    CreatedBy string `json:"created_by"`
    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (t Tag) TableName() string {
    return "blog_tag"
}
