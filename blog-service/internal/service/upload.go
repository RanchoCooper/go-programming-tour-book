package service

import (
    "errors"
    "mime/multipart"
    "os"

    "go-programming-tour-book/blog-service/global"
    "go-programming-tour-book/blog-service/pkg/upload"
)

/**
 * @author Rancho
 * @date 2021/11/30
 */

type FileInfo struct {
	Name      string
	AccessUrl string
}

func (svc *Service) UploadFile(fileType upload.FileType, file multipart.File, fileHeader *multipart.FileHeader) (*FileInfo, error) {
    fileName := upload.GetFileName(fileHeader.Filename)
    savePath := upload.GetSavePath()
    dst := savePath + "/" + fileName
    if !upload.CheckContainExt(fileType, fileName) {
        return nil, errors.New("file suffix is not supported")
    }
    if upload.CheckSavePath(savePath) {
        err := upload.CreateSavePath(savePath, os.ModePerm)
        if err != nil {
            return nil, errors.New("failed to create save directory")
        }
    }
    if upload.CheckOverMaxSize(fileType, file) {
        return nil, errors.New("exceeded maximum file limit")
    }
    if upload.NoPermission(savePath) {
        return nil, errors.New("insufficient file permission")
    }

    if err := upload.SaveFile(fileHeader, dst); err != nil {
        return nil, err
    }
    accessUrl := global.AppSetting.UploadServerURl + "/" + fileName

    return &FileInfo{
        Name:      fileName,
        AccessUrl: accessUrl,
    }, nil
}
