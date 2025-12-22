package utils

import (
	"context"
	"fmt"
	"last-project/app/config/cloudinary_config"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func UploadFotoToCloudinary(filepath string) (string, error) {
	cld, err := cloudinary.NewFromParams(
		cloudinary_config.CloudinaryConfig().CLOUDINARY_CLOUD_NAME,
		cloudinary_config.CloudinaryConfig().CLOUDINARY_API_KEY,
		cloudinary_config.CloudinaryConfig().CLOUDINARY_API_SECRET,
	)

	if err != nil {
		fmt.Println("An error occurred while cloudinary. " + err.Error())
	}

	uploadResult, errUpload := cld.Upload.Upload(
		context.Background(),
		filepath,
		uploader.UploadParams{
			Folder: "rakamin/toko",
		},
	)

	if errUpload != nil {
		fmt.Println("An error occurred while uploading the image. " + errUpload.Error())
	}

	return uploadResult.SecureURL, nil
}
