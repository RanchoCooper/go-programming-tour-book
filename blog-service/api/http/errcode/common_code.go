package errcode

/**
 * @author Rancho
 * @date 2021/11/26
 */

var (
    Success                   = NewError(0, "成功")
    ServerError               = NewError(10000000, "内部错误")
    InvalidParams             = NewError(10000001, "入参错误")
    NotFound                  = NewError(10000002, "找不到")
    UnauthorizedAuthNotExist  = NewError(10000003, "鉴权失败，找不到对应的AppKey和AppSecret")
    UnauthorizedTokenError    = NewError(10000004, "鉴权失败，Token错误")
    UnauthorizedTokenTimeout  = NewError(10000005, "鉴权失败，Token超时")
    UnauthorizedTokenGenerate = NewError(10000006, "鉴权失败，Token生成失败")
    TooManyRequests           = NewError(10000007, "请求过多")
    DBError                   = NewError(10000008, "数据库错误")
)
