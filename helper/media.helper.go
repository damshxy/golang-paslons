package helper

import (
	"context"
	"go-paslons-crud/database"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func ImageUploadHelper(input interface{}) (string, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    cld, err := cloudinary.NewFromParams(database.EnvCloudName(), database.EnvCloudAPIKey(), database.EnvCloudAPISecret())
    if err != nil {
        return "", err
    }

    uploadParam, err := cld.Upload.Upload(ctx, input, uploader.UploadParams{Folder: database.EnvCloudUploadFolder()})
    if err != nil {
        return "", err
    }
    return uploadParam.SecureURL, nil
}