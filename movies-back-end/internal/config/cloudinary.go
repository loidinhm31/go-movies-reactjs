package config

import (
	"github.com/cloudinary/cloudinary-go/v2"
	"movies-service/config"
	"sync"
)

type CloudinaryGetter interface {
	GetCloudinaryObject(cloudinaryConfig config.CloudinaryConfig) *cloudinary.Cloudinary
}

type CloudinaryObject struct {
}

var (
	cloudinaryObject *cloudinary.Cloudinary
	once             sync.Once
)

func credentials(cloudinaryConfig config.CloudinaryConfig) *cloudinary.Cloudinary {
	// Add your Cloudinary credentials, set configuration parameter
	cld, _ := cloudinary.NewFromParams(
		cloudinaryConfig.CloudName,
		cloudinaryConfig.ApiKey,
		cloudinaryConfig.ApiSecret,
	)
	cld.Config.URL.Secure = true
	return cld
}

func loadCloudinaryObject(cloudinaryConfig config.CloudinaryConfig) {
	object := credentials(cloudinaryConfig)
	cloudinaryObject = object
}

func (lo *CloudinaryObject) GetCloudinaryObject(cloudinaryConfig config.CloudinaryConfig) *cloudinary.Cloudinary {
	once.Do(func() {
		loadCloudinaryObject(cloudinaryConfig)
	}) // loadCloudinaryObject will be called only once

	return cloudinaryObject
}
