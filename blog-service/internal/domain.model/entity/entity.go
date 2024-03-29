package entity

type Model struct {
    ID         uint32 `gorm:"primary_key" json:"id"`
    CreatedBy  string `json:"created_by"`
    ModifiedBy string `json:"modified_by"`
    CreatedOn  uint32 `json:"created_on"`
    ModifiedOn uint32 `json:"modified_on"`
    DeletedOn  uint32 `json:"deleted_on"`
    IsDel      uint32 `json:"is_del"`
}
