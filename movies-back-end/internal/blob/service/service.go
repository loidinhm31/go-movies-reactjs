package service

import (
	"context"
	"fmt"
	"github.com/cloudinary/cloudinary-go/v2/api"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"log"
	"mime/multipart"
	"movies-service/config"
	"movies-service/internal/blob"
	config2 "movies-service/internal/config"
	"movies-service/internal/control"
	"movies-service/internal/errors"
	"movies-service/internal/middlewares"
	"strconv"
	"time"
)

type blobService struct {
	cfg  *config.Config
	ctrl control.Service
}

func NewBlobService(cfg *config.Config, ctrl control.Service) blob.Service {
	return &blobService{
		cfg,
		ctrl,
	}
}

func (is *blobService) UploadVideo(ctx context.Context, file multipart.File) (string, error) {
	log.Println("processing uploading file...")
	log.Println("checking role...")
	username := fmt.Sprintf("%s", ctx.Value(middlewares.CtxUserKey))
	if !is.ctrl.CheckPrivilege(username) {
		return "", errors.ErrUnAuthorized
	}

	object := config2.CloudinaryObject{}
	cld := object.GetCloudinaryObject(is.cfg.Cloudinary)
	log.Println("uploading file...")
	fileKey := is.cfg.Cloudinary.VideoFolderPath + strconv.FormatInt(time.Now().Unix(), 10)
	res, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{
		PublicID:   fileKey,
		EagerAsync: api.Bool(true),
	})
	if err != nil {
		return "", err
	}
	log.Printf("complete upload file %s\n", fileKey)
	return fileKey + "." + res.Format, nil
}

func (is *blobService) DeleteVideo(ctx context.Context, fileId string) (string, error) {
	log.Println("checking role...")
	username := fmt.Sprintf("%s", ctx.Value(middlewares.CtxUserKey))
	if !is.ctrl.CheckPrivilege(username) {
		return "", errors.ErrUnAuthorized
	}

	object := config2.CloudinaryObject{}
	cld := object.GetCloudinaryObject(is.cfg.Cloudinary)
	log.Println("deleting file...")
	filePath := is.cfg.Cloudinary.VideoFolderPath + fileId
	res, err := cld.Upload.Destroy(ctx, uploader.DestroyParams{
		PublicID:     filePath,
		ResourceType: "video",
	})
	if err != nil {
		return "", err
	}
	log.Printf("deleted file %s\n", filePath)
	return res.Result, nil
}
