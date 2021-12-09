package handle

import (
    "github.com/gin-gonic/gin"
    "github.com/spf13/cast"

    "go-programming-tour-book/blog-service/api/http/DTO"
    "go-programming-tour-book/blog-service/api/http/errcode"
    "go-programming-tour-book/blog-service/internal/domain.model/tag"
    "go-programming-tour-book/blog-service/util/logger"
)

/**
 * @author Rancho
 * @date 2021/11/26
 */

// GetTag get a single tag list
// @Summary 获取单个标签
// @Produce json
// @Param id path int true "标签ID"
// @Success 200 {object} model.Tag "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /tags/{id} [get]
func GetTag(c *gin.Context) {
    response := NewResponse(c)
    t := &tag.Tag{
        ID: cast.ToUint(c.Param("id")),
    }
    t, err := t.GetTag()
    if err != nil {
        logger.Log.Errorf(c, "domain.GetTag err: %v", err.Error())
        response.ToErrorResponse(errcode.DBError)
    }
    response.ToResponse(t)
    return
}

// ListTag get tag list
// @Summary 获取多个标签
// @Produce json
// @Param name query string false "标签名称" maxlength(100)
// @Param state query int false "状态" Enums(0, 1) default(1)
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} model.TagSwagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /tags [get]
func ListTag(c *gin.Context) {
    param := DTO.TagListRequest{}
    valid, errs := BindAndValid(c, &param)
    response := NewResponse(c)
    if !valid {
        logger.Log.Errorf(c, "BindAndValid errs: %v", errs.Errors())
        response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
        return
    }
    pager := Pager{
        Page:       GetPage(c),
        PageSize:   GetPageSize(c),
        PageOffset: GetPageOffset(GetPage(c), GetPageSize(c)),
    }
    t := &tag.Tag{Name: param.Name, State: &param.State}
    totalRows, err := t.CountTag()
    if err != nil {
        logger.Log.Errorf(c, "domain.CountTag err: %v", err.Error())
        response.ToErrorResponse(errcode.DBError)
        return
    }
    tags, err := t.GetTagList(pager.PageOffset, pager.PageSize)
    if err != nil {
        logger.Log.Errorf(c, "domain.GetTagList err: %v", err.Error())
        response.ToErrorResponse(errcode.DBError)
        return
    }
    response.ToResponseList(tags, totalRows)
    return
}

// CreateTag create a tag
// @Summary 新增标签
// @Produce json
// @Param name body string true "标签名称" minlength(3) maxlength(100)
// @Param state body int false "状态" Enums(0, 1) default(1)
// @Param created_by body string true "创建者" minlength(3) maxlength(100)
// @Success 200 {object} model.Tag "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /tags [post]
func CreateTag(c *gin.Context) {
    param := DTO.CreateTagRequest{}
    valid, errs := BindAndValid(c, &param)
    response := NewResponse(c)
    if !valid {
        logger.Log.Errorf(c, "BindAndValid errs: %v", errs.Errors())
        response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
        return
    }
    t := &tag.Tag{
        Name:      param.Name,
        State:     &param.State,
        CreatedBy: param.CreatedBy,
    }
    t, err := t.CreateTag()
    if err != nil {
        logger.Log.Errorf(c, "domain.CreateTag err: %v", err.Error())
        response.ToErrorResponse(errcode.DBError)
        return
    }
    response.ToResponse(t)
    return
}

// UpdateTag update a tag
// @Summary 更新标签
// @Produce json
// @Param id path int true "标签ID"
// @Param name body string true "标签名称" minlength(3) maxlength(100)
// @Param state body int false "状态" Enums(0, 1) default(1)
// @Param created_by body string true "修改者" minlength(3) maxlength(100)
// @Success 200 {object} model.Tag "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /tags/{id} [put]
func UpdateTag(c *gin.Context) {
    response := NewResponse(c)
    param := DTO.UpdateTagRequest{}
    valid, errs := BindAndValid(c, &param)
    if !valid {
        logger.Log.Errorf(c, "BindAndValid errs: %v", errs.Errors())
        response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
        return
    }
    t := &tag.Tag{
        ID:        cast.ToUint(c.Param("id")),
        Name:      param.Name,
        State:     &param.State,
        CreatedBy: param.ModifiedBy,
    }
    err := t.UpdateTag()
    if err != nil {
        logger.Log.Errorf(c, "domain.UpdateTag err: %v", err.Error())
        response.ToErrorResponse(errcode.DBError)
        return
    }
    response.ToResponse(t)
    return
}

// DeleteTag delete a tag
// @Summary 删除标签
// @Produce json
// @Param id path int true "标签ID"
// @Success 200 {object} model.Tag "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /tags/{id} [delete]
func DeleteTag(c *gin.Context) {
    param := DTO.DeleteTagRequest{}
    response := NewResponse(c)
    valid, errs := BindAndValid(c, &param)
    if !valid {
        logger.Log.Errorf(c, "app.BindAndValid errs: %v", errs)
        errResp := errcode.InvalidParams.WithDetails(errs.Errors()...)
        response.ToErrorResponse(errResp)
        return
    }
    t := &tag.Tag{
        ID: uint(param.ID),
    }
    err := t.DeleteTag()
    if err != nil {
        logger.Log.Errorf(c, "domain.DeleteTag err: %v", err.Error())
        response.ToErrorResponse(errcode.DBError)
        return
    }

    response.ToResponse(gin.H{})
    return
}
