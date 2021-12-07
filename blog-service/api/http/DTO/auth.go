package DTO

/**
 * @author Rancho
 * @date 2021/12/8
 */

type AuthRequest struct {
    AppKey    string `form:"app_key" binding:"required"`
    AppSecret string `form:"app_secret" binding:"required"`
}
