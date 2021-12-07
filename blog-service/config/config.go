package config

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "path/filepath"
    "runtime"

    "gopkg.in/yaml.v3"
)

/**
 * @author Rancho
 * @date 2021/12/8
 */

const (
    configFilePath        = "/config.yaml"
    privateConfigFilePath = "/config.private.yaml"
)

var Config = &config{}

type config struct {
    Server   ServerSetting   `yaml:"Server"`
    App      AppSetting      `yaml:"App"`
    JWT      JWTSetting      `yaml:"JWT"`
    Database DatabaseSetting `yaml:"Database"`
}

type ServerSetting struct {
    RunMode      string `yaml:"RunMode"`
    HTTPPort     string `yaml:"HTTPPort"`
    ReadTimeout  int    `yaml:"ReadTimeout"`
    WriteTimeout int    `yaml:"WriteTimeout"`
}

type AppSetting struct {
    DefaultPageSize      int      `yaml:"DefaultPageSize"`
    MaxPageSize          int      `yaml:"MaxPageSize"`
    LogSavePath          string   `yaml:"LogSavePath"`
    LogFileName          string   `yaml:"LogFileName"`
    LogFileExt           string   `yaml:"LogFileExt"`
    UploadSavePath       string   `yaml:"UploadSavePath"`
    UploadServerURL      string   `yaml:"UploadServerUrl"`
    UploadImageMaxSize   int      `yaml:"UploadImageMaxSize"`
    UploadImageAllowExts []string `yaml:"UploadImageAllowExts"`
}

type JWTSetting struct {
    Secret string `yaml:"Secret"`
    Issuer string `yaml:"Issuer"`
    Expire int    `yaml:"Expire"`
}

type DatabaseSetting struct {
    DBType       string `yaml:"DBType"`
    UserName     string `yaml:"UserName"`
    Password     string `yaml:"Password"`
    Host         string `yaml:"Host"`
    DBName       string `yaml:"DBName"`
    TablePrefix  string `yaml:"TablePrefix"`
    Charset      string `yaml:"Charset"`
    ParseTime    bool   `yaml:"ParseTime"`
    TimeZone     string `yaml:"TimeZone"`
    MaxIdleConns int    `yaml:"MaxIdleConns"`
    MaxOpenConns int    `yaml:"MaxOpenConns"`
}

func readYamlConfig(configPath string) {
    yamlFile, err := filepath.Abs(configPath)
    if err != nil {
        log.Fatalf("invalid config file path, err: %v", err)
    }
    content, err := ioutil.ReadFile(yamlFile)
    if err != nil {
        log.Printf("read config file fail, err: %v", err)
    }
    err = yaml.Unmarshal(content, Config)
    if err != nil {
        log.Printf("config file unmarshal fail, err: %v", err)
    }

}

func init() {
    _, file, _, _ := runtime.Caller(0)
    configPath := filepath.Dir(file)

    readYamlConfig(configPath + configFilePath)
    if Config.Database.Password == "" {
        // read private config
        readYamlConfig(configPath + privateConfigFilePath)
    }

    bf, _ := json.MarshalIndent(Config, "", "    ")
    fmt.Printf("Config:\n%s\n", string(bf))
}
