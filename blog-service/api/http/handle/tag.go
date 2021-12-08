package handle

import (
    "github.com/gin-gonic/gin"

    "go-programming-tour-book/blog-service/api/http/DTO"
    "go-programming-tour-book/blog-service/api/http/errcode"
    "go-programming-tour-book/blog-service/api/http/router"
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
    router.NewResponse(c).ToErrorResponse(errcode.ServerError)
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
    response := router.NewResponse(c)
    valid, errs := BindAndValid(c, &param)
    if !valid {
        logger.Log.Errorf(c, "app.BindAndValid errs: %v", errs)
        errResp := errcode.InvalidParams.WithDetails(errs.Errors()...)
        response.ToErrorResponse(errResp)
        return
    }

    response.ToResponse(gin.H{})
    return
}
