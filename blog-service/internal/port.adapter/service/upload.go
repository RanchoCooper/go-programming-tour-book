package service

import (
    "errors"
    "mime/multipart"
    "os"

    "go-programming-tour-book/blog-service/config"
    "go-programming-tour-book/blog-service/util"
)

/**
 * @author Rancho
 * @date 2021/11/30
 */

type FileInfo struct {
    Name      string
    AccessUrl string
}

func UploadFile(fileType util.FileType, file multipart.File, fileHeader *multipart.FileHeader) (*FileInfo, error) {
    fileName := util.GetFileName(fileHeader.Filename)
    savePath := config.Config.App.UploadSavePath
    dst := savePath + "/" + fileName
    if !util.CheckContainExt(fileType, fileName) {
        return nil, errors.New("file suffix is not supported")
    }
    if util.CheckSavePath(savePath) {
        err := util.CreateSavePath(savePath, os.ModePerm)
        if err != nil {
            return nil, errors.New("failed to create save directory")
        }
    }
    if util.CheckOverMaxSize(fileType, file) {
        return nil, errors.New("exceeded maximum file limit")
    }
    if util.NoPermission(savePath) {
        return nil, errors.New("insufficient file permission")
    }

    if err := util.SaveFile(fileHeader, dst); err != nil {
        return nil, err
    }
    accessUrl := config.Config.App.UploadServerURL + "/" + fileName

    return &FileInfo{
        Name:      fileName,
        AccessUrl: accessUrl,
    }, nil
}
