package entity

/**
 * @author Rancho
 * @date 2022/1/5
 */

type Tag struct {
    *Model
    Name      string `json:"name"`
    State     uint8  `json:"state"`
    changeMap map[string]interface{}
}

func (t Tag) TableName() string {
    return "blog_tag"
}

func (t *Tag) GetChangeMap() map[string]interface{} {
    return t.changeMap
}
