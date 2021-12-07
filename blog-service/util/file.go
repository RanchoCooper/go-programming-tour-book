package util

import (
    "io"
    "io/ioutil"
    "mime/multipart"
    "os"
    "path"
    "strings"

    "go-programming-tour-book/blog-service/config"
)

/**
 * @author Rancho
 * @date 2021/11/30
 */

type FileType int

const (
    TypeImage FileType = iota + 1
    TypeExcel
    TypeTxt
)

func GetFileExt(name string) string {
    return path.Ext(name)
}

func GetFileName(name string) string {
    ext := GetFileExt(name)
    fileName := strings.TrimSuffix(name, ext)
    fileName = EncodeMD5(fileName)

    return fileName + ext
}

func CheckSavePath(dst string) bool {
    _, err := os.Stat(dst)
    return os.IsNotExist(err)
}

func CheckContainExt(t FileType, name string) bool {
    ext := GetFileExt(name)
    ext = strings.ToUpper(ext)

    switch t {
    case TypeImage:
        for _, allowExt := range config.Config.App.UploadImageAllowExts {
            if strings.ToUpper(allowExt) == ext {
                return true
            }
        }
    }

    return false
}

func CheckOverMaxSize(t FileType, f multipart.File) bool {
    content, _ := ioutil.ReadAll(f)
    size := len(content)
    switch t {
    case TypeImage:
        if size >= config.Config.App.UploadImageMaxSize*1024*1024 {
            return true
        }
    }

    return false
}

func NoPermission(dst string) bool {
    _, err := os.Stat(dst)
    return os.IsPermission(err)
}

func CreateSavePath(dst string, perm os.FileMode) error {
    err := os.MkdirAll(dst, perm)
    if err != nil {
        return err
    }

    return nil
}

func SaveFile(fh *multipart.FileHeader, dst string) error {
    file, err := fh.Open()
    if err != nil {
        return err
    }
    defer file.Close()

    out, err := os.Create(dst)
    if err != nil {
        return err
    }
    defer out.Close()

    _, err = io.Copy(out, file)
    return err
}
