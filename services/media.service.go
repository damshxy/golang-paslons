package services

import (
	"go-paslons-crud/helper"
	"go-paslons-crud/models"

	"github.com/go-playground/validator/v10"
)

var (
    validate = validator.New()
)

type mediaUpload interface {
    FileUpload(file models.File) (string, error)
}

type media struct {}

func NewMediaUpload() mediaUpload {
    return &media{}
}

func (*media) FileUpload(file models.File) (string, error) {
    err := validate.Struct(file)
    if err != nil {
        return "", err
    }

    uploadUrl, err := helper.ImageUploadHelper(file.File)
    if err != nil {
        return "", err
    }
    return uploadUrl, nil
}