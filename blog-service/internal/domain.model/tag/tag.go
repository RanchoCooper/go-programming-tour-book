package tag

import (
    "gorm.io/gorm"
)

/**
 * @author Rancho
 * @date 2021/11/26
 */

type Tag struct {
    gorm.Model
    Name  string `json:"name"`
    State uint8  `json:"state"`
}

func (t Tag) TableName() string {
    return "blog_tag"
}
