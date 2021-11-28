package routers

import (
    "go-programming-tour-book/blog-service/global"
    "go-programming-tour-book/blog-service/internal/service"
    "go-programming-tour-book/blog-service/pkg/app"
    "go-programming-tour-book/blog-service/pkg/errcode"

    "github.com/gin-gonic/gin"
)

/**
 * @author Rancho
 * @date 2021/11/26
 */

type Tag struct {

}

func NewTag() Tag {
    return Tag{}
}

// Get get a single tag list
// @Summary 获取单个标签
// @Produce json
// @Param id path int true "标签ID"
// @Success 200 {object} model.Tag "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /tags/{id} [get]
func (t Tag) Get(c *gin.Context) {
    app.NewResponse(c).ToErrorResponse(errcode.ServerError)
}

// List get tag list
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
func (t Tag) List(c *gin.Context) {
    param := service.TagListRequest{}
    response := app.NewResponse(c)
    valid, errs := app.BindAndValid(c, &param)
    if !valid {
        global.Logger.Errorf("app.BindAndValid errs: %v", errs)
        errResp := errcode.InvalidParams.WithDetails(errs.Errors()...)
        response.ToErrorResponse(errResp)
        return
    }

    svc := service.New(c.Request.Context())
    pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
    totalRows, err := svc.CountTag(&service.CountTagRequest{
        Name: param.Name,
        State: param.State,
    })
    if err != nil {
        global.Logger.Errorf("svc.CountTag err: %v", err)
        response.ToErrorResponse(errcode.ErrorCountTagFail)
        return
    }

    tags, err := svc.GetTagList(&param, &pager)
    if err != nil {
        global.Logger.Errorf("svc.GetTagList err: %v", err)
        response.ToErrorResponse(errcode.ErrorGetTagListFail)
        return
    }

    response.ToResponseList(tags, totalRows)
    return
}

// Create create a tag
// @Summary 新增标签
// @Produce json
// @Param name body string true "标签名称" minlength(3) maxlength(100)
// @Param state body int false "状态" Enums(0, 1) default(1)
// @Param created_by body string true "创建者" minlength(3) maxlength(100)
// @Success 200 {object} model.Tag "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /tags [post]
func (t Tag) Create(c *gin.Context) {
    param := service.CreateTagRequest{}
    response := app.NewResponse(c)
    valid, errs := app.BindAndValid(c, &param)
    if !valid {
        global.Logger.Errorf("app.BindAndValid errs: %v", errs)
        errResp := errcode.InvalidParams.WithDetails(errs.Errors()...)
        response.ToErrorResponse(errResp)
        return
    }

    svc := service.New(c.Request.Context())
    err := svc.CreateTag(&param)
    if err != nil {
        global.Logger.Errorf("svc.CreateTag err: %v", err)
        response.ToErrorResponse(errcode.ErrorCreateTagFail)
        return
    }

    response.ToResponse(gin.H{})
    return
}

// Update update a tag
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
func (t Tag) Update(c *gin.Context) {
    param := service.UpdateTagRequest{}
    response := app.NewResponse(c)
    valid, errs := app.BindAndValid(c, &param)
    if !valid {
        global.Logger.Errorf("app.BindAndValid errs: %v", errs)
        errResp := errcode.InvalidParams.WithDetails(errs.Errors()...)
        response.ToErrorResponse(errResp)
        return
    }

    svc := service.New(c.Request.Context())
    err := svc.UpdateTag(&param)
    if err != nil {
        global.Logger.Errorf("svc.UpdateTag err: %v", err)
        response.ToErrorResponse(errcode.ErrorUpdateTagFail)
        return
    }

    response.ToResponse(gin.H{})
    return
}

// Delete delete a tag
// @Summary 删除标签
// @Produce json
// @Param id path int true "标签ID"
// @Success 200 {object} model.Tag "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /tags/{id} [delete]
func (t Tag) Delete(c *gin.Context) {
    param := service.DeleteTagRequest{}
    response := app.NewResponse(c)
    valid, errs := app.BindAndValid(c, &param)
    if !valid {
        global.Logger.Errorf("app.BindAndValid errs: %v", errs)
        errResp := errcode.InvalidParams.WithDetails(errs.Errors()...)
        response.ToErrorResponse(errResp)
        return
    }

    svc := service.New(c.Request.Context())
    err := svc.DeleteTag(&param)
    if err != nil {
        global.Logger.Errorf("svc.DeleteTag err: %v", err)
        response.ToErrorResponse(errcode.ErrorDeleteTagFail)
        return
    }

    response.ToResponse(gin.H{})
    return
}
