package dto

/**
 * @author Rancho
 * @date 2022/1/6
 */

type CountTagRequest struct {
    Name  string `form:"name" binding:"max=100"`
    State uint8  `form:"state" binding:"oneof=0 1"`
}

type TagListRequest struct {
    Name  string `form:"name" binding:"max=100"`
    State uint8  `form:"state" binding:"oneof=0 1"`
}

type CreateTagRequest struct {
    Name      string `json:"name" binding:"required,min=3,max=100"`
    CreatedBy string `json:"created_by" validate:"required,min=3,max=100"`
    State     uint8  `json:"state" validate:"oneof=0 1"`
}

type UpdateTagRequest struct {
    ID         uint32 `form:"id" binding:"required,gte=1"`
    Name       string `json:"name" binding:"max=100"`
    State      uint8  `json:"state" binding:"oneof=0 1"`
    ModifiedBy string `json:"modified_by" binding:"required,min=3,max=100"`
}

type DeleteTagRequest struct {
    ID uint32 `uri:"id" binding:"required,gte=1"`
}
