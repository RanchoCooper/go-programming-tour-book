package global

import (
	"github.com/RanchoCooper/go-programming-tour-book/blog-service/configs"
	"github.com/RanchoCooper/go-programming-tour-book/blog-service/pkg/logger"
)

var (
	ServerSetting   *configs.ServerSettingS
	AppSetting      *configs.AppSettingS
	DatabaseSetting *configs.DatabaseSettingS
	JWTSetting *configs.JWTSettingS

	Logger *logger.Logger
)
