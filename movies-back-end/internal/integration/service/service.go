package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"log"
	"mime/multipart"
	"movies-service/config"
	"movies-service/internal/control"
	"movies-service/internal/integration"
	"movies-service/internal/middlewares"
	"movies-service/internal/movies"
	"strconv"
	"time"
)

type integrationService struct {
	cfg             *config.Config
	ctrl            control.Service
	movieRepository movies.MovieRepository
}

func NewIntegrationService(cfg *config.Config, ctrl control.Service, movieRepository movies.MovieRepository) integration.Service {
	return &integrationService{
		cfg,
		ctrl,
		movieRepository,
	}
}

func (is *integrationService) credentials() *cloudinary.Cloudinary {
	// Add your Cloudinary credentials, set configuration parameter
	cld, _ := cloudinary.NewFromParams(
		is.cfg.Cloudinary.CloudName,
		is.cfg.Cloudinary.ApiKey,
		is.cfg.Cloudinary.ApiSecret,
	)
	cld.Config.URL.Secure = true
	return cld
}

func (is *integrationService) UploadVideo(ctx context.Context, file multipart.File) (string, error) {
	log.Println("checking role...")
	username := fmt.Sprintf("%s", ctx.Value(middlewares.CtxUserKey))
	if !is.ctrl.CheckUser(username) {
		return "", errors.New("user not found")
	}

	cld := is.credentials()
	log.Println("uploading file...")
	fileKey := is.cfg.Cloudinary.FolderPath + strconv.FormatInt(time.Now().Unix(), 10)
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

func (is *integrationService) DeleteVideo(ctx context.Context, fileId string) (string, error) {
	log.Println("checking role...")
	username := fmt.Sprintf("%s", ctx.Value(middlewares.CtxUserKey))
	if !is.ctrl.CheckUser(username) {
		return "", errors.New("user not found")
	}

	cld := is.credentials()

	log.Println("deleting file...")
	filePath := is.cfg.Cloudinary.FolderPath + fileId
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
