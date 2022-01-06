package dto

/**
 * @author Rancho
 * @date 2022/1/7
 */

type AuthRequest struct {
    AppKey    string `json:"app_key" binding:"required"`
    AppSecret string `json:"app_secret" binding:"required"`
}
