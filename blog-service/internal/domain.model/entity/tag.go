package entity

/**
 * @author Rancho
 * @date 2022/1/5
 */

type Tag struct {
    ID         uint32                 `gorm:"primary_key" json:"id"`
    Name       string                 `gorm:"name" json:"name" structs:",omitempty,underline"`
    State      uint8                  `json:"state" structs:",omitempty,underline"`
    CreatedBy  string                 `json:"created_by" structs:",omitempty,underline"`
    ModifiedBy string                 `json:"modified_by" structs:",omitempty,underline"`
    CreatedOn  uint32                 `json:"created_on" structs:",omitempty,underline"`
    ModifiedOn uint32                 `json:"modified_on" structs:",omitempty,underline"`
    DeletedOn  uint32                 `json:"deleted_on" structs:",omitempty,underline"`
    IsDel      uint32                 `json:"is_del" structs:",omitempty,underline"`
    ChangeMap  map[string]interface{} `json:"-" structs:",omitempty,underline"`
}

func (t Tag) TableName() string {
    return "blog_tag"
}
