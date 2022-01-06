package entity

/**
 * @author Rancho
 * @date 2022/1/5
 */

type Tag struct {
    ID         uint32 `gorm:"primary_key" json:"id"`
    Name       string `json:"name"`
    State      uint8  `json:"state"`
    CreatedBy  string `json:"created_by"`
    ModifiedBy string `json:"modified_by"`
    CreatedOn  uint32 `json:"created_on"`
    ModifiedOn uint32 `json:"modified_on"`
    DeletedOn  uint32 `json:"deleted_on"`
    IsDel      uint32 `json:"is_del"`
    changeMap  map[string]interface{}
}

func (t Tag) TableName() string {
    return "blog_tag"
}

func (t *Tag) GetChangeMap() map[string]interface{} {
    return t.changeMap
}
