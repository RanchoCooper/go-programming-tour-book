package entity

/**
 * @author Rancho
 * @date 2022/1/7
 */

type Auth struct {
    ID         uint32 `gorm:"primary_key" json:"id"`
    AppKey     string `json:"app_key"`
    AppSecret  string `json:"app_secret"`
    CreatedBy  string `json:"created_by"`
    ModifiedBy string `json:"modified_by"`
    CreatedOn  uint32 `json:"created_on"`
    ModifiedOn uint32 `json:"modified_on"`
    DeletedOn  uint32 `json:"deleted_on"`
    IsDel      uint32 `json:"is_del"`
    changeMap  map[string]interface{}
}

func (a Auth) TableName() string {
    return "blog_auth"
}

func (a *Auth) GetChangeMap() map[string]interface{} {
    return a.changeMap
}
