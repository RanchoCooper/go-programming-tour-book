package setting

import (
	"time"
)

/**
 * @author Rancho
 * @date 2021/11/26
 */

type ServerSettingS struct {
	RunMode      string
	HTTPPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type AppSettingS struct {
	DefaultPageSize int
	MaxPageSize     int
	LogSavePath     string
	LogFileName     string
	LogFileExt      string
    UploadSavePath string
    UploadServerURl string
    UploadiMageMaxSize int
    UploadImageAllowExts []string
}

type DatabaseSettingS struct {
    DBType       string
    UserName     string
    Password     string
    Host         string
    DBName       string
    TablePrefix  string
    Charset      string
    ParseTime    bool
    MaxIdleConns int
    MaxOpenConns int
}

func (s *Setting) ReadSection(k string, v interface{}) error {
    err := s.vp.UnmarshalKey(k, v)
    if err != nil {
        return err
    }

    return nil
}