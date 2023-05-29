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

func (is *blobService) UploadFile(ctx context.Context, file multipart.File, fileType string) (string, error) {
	log.Println("processing uploading file...")
	log.Println("checking role...")
	username := fmt.Sprintf("%s", ctx.Value(middlewares.CtxUserKey))
	if !is.ctrl.CheckPrivilege(username) {
		return "", errors.ErrUnAuthorized
	}

	if fileType == "" {
		return "", errors.ErrInvalidInput
	}

	object := config2.CloudinaryObject{}
	cld := object.GetCloudinaryObject(is.cfg.Cloudinary)
	log.Printf("uploading %s file...\n", fileType)

	var fileKey string
	if fileType == "video" {
		fileKey = fmt.Sprintf("%s/%s/%s", is.cfg.Cloudinary.FolderPath, "videos", strconv.FormatInt(time.Now().Unix(), 10))
	} else if fileType == "image" {
		fileKey = fmt.Sprintf("%s/%s/%s", is.cfg.Cloudinary.FolderPath, "images", strconv.FormatInt(time.Now().Unix(), 10))
	}

	res, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{
		PublicID:   fileKey,
		EagerAsync: api.Bool(true),
	})
	if err != nil {
		return "", err
	}

	var fileResult string
	if fileType == "video" {
		fileResult = fileKey + "." + res.Format
	} else if fileType == "image" {
		fileResult = fmt.Sprintf("%s/%s/%s/%s.%s",
			"https://res.cloudinary.com",
			is.cfg.Cloudinary.CloudName,
			"image/upload/c_scale,w_200,h_300",
			fileKey, res.Format,
		)
	}
	log.Printf("complete upload file %s\n", fileKey)
	return fileResult, nil
}

func (is *blobService) DeleteFile(ctx context.Context, fileId string, fileType string) (string, error) {
	log.Println("checking role...")
	username := fmt.Sprintf("%s", ctx.Value(middlewares.CtxUserKey))
	if !is.ctrl.CheckPrivilege(username) {
		return "", errors.ErrUnAuthorized
	}

	object := config2.CloudinaryObject{}
	cld := object.GetCloudinaryObject(is.cfg.Cloudinary)
	log.Printf("deleting %s file...\n", fileType)

	var res *uploader.DestroyResult
	var err error
	var filePath string
	if fileType == "video" {
		filePath = fmt.Sprintf("%s/%s/%s", is.cfg.Cloudinary.FolderPath, "videos", fileId)
		res, err = cld.Upload.Destroy(ctx, uploader.DestroyParams{
			PublicID:     filePath,
			ResourceType: "video",
		})
	} else if fileType == "image" {
		filePath = fmt.Sprintf("%s/%s/%s", is.cfg.Cloudinary.FolderPath, "images", fileId)
		res, err = cld.Upload.Destroy(ctx, uploader.DestroyParams{
			PublicID:     filePath,
			ResourceType: "image",
		})
	}

	if err != nil {
		return "", err
	}
	log.Printf("deleted file %s\n", filePath)
	return res.Result, nil
}
