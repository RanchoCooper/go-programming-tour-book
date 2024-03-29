package cmd

import (
    "log"
    "strings"

    "go-programming-tour-book/tour/internal/word"

    "github.com/spf13/cobra"
)

/**
 * @author Rancho
 * @date 2021/11/4
 */

const (
    ModeUpper = iota + 1
    ModeLower
    ModeUnderscoreToUpperCamelCase
    ModeUnderscoreToLowerCamelCase
    ModeCamelCaseToUnderscore
)

var str string
var mode int8

 var wordCmd = &cobra.Command{
     Use:   "word",
     Short: "单词格式转换",
     Long:  desc,
     Run: func(cmd *cobra.Command, args []string) {
         var content string
         switch mode {
         case ModeUpper:
             content = word.ToUpper(str)
         case ModeLower:
             content = word.ToLower(str)
         case ModeUnderscoreToUpperCamelCase:
             content = word.UnderscoreToUpperCamelCase(str)
         case ModeUnderscoreToLowerCamelCase:
             content = word.UnderscoreToLowerCamelCase(str)
         case ModeCamelCaseToUnderscore:
             content = word.CamelCaseToUnderscore(str)
         default:
             log.Fatalf("暂不支持该转换模式，请执行 help word 查看帮助文档")
         }

         log.Printf("输出结果: %s", content)
     },
 }

var desc = strings.Join([]string{
    "该子命令支持各种单词格式转换，模式如下：",
    "1：全部转大写",
    "2：全部转小写",
    "3：下划线转大写驼峰",
    "4：下划线转小写驼峰",
    "5：驼峰转下划线",
}, "\n")

 func init() {
    wordCmd.Flags().StringVarP(&str, "str", "s", "", "请输入单词内容")
    wordCmd.Flags().Int8VarP(&mode, "mode", "m", 0, "请输入单词转换模式")
 }
