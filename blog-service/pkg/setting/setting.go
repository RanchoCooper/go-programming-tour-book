package setting

import (
    "github.com/spf13/viper"
)

/**
 * @author Rancho
 * @date 2021/11/26
 */

type Setting struct {
    vp *viper.Viper
}

func NewSetting() (*Setting, error) {
    vp := viper.New()
    vp.AddConfigPath("blog-service/configs")
    vp.SetConfigName("config")
    vp.SetConfigType("yaml")

    err := vp.ReadInConfig()
    if err != nil {
        return nil, err
    }

    return &Setting{vp: vp}, nil
}

