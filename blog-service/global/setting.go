package global

import (
    "go-programming-tour-book/blog-service/pkg/logger"
    "go-programming-tour-book/blog-service/pkg/setting"
)

/**
 * @author Rancho
 * @date 2021/11/26
 */

var (
    ServerSetting   *setting.ServerSettingS
    AppSetting      *setting.AppSettingS
    DatabaseSetting *setting.DatabaseSettingS
    Logger          *logger.Logger
    JWTSetting      *setting.JWTSettingS
)
