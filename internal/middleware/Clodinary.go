package middleware

import (
	"log"
	"net/http"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/gofiber/fiber/v2"
)

var cloudinaryClient *cloudinary.Cloudinary

func init() {
	var err error
	cloudinaryClient, err = cloudinary.NewFromParams("dgdvfpzrw", "666481357989953", "qiTWQejmUoN4iariFo44_2e8pCI")
	if err != nil {
		log.Fatalf("Gagal Menghubungkan Ke Cloudinary: %v", err)
	}
}

func UploadImageMiddleware(c *fiber.Ctx) error {

	file, err := c.FormFile("image")
	if err != nil {
		return c.Next()
	}

	fileContent, err := file.Open()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Gagal Untuk Mengupload Gambar",
		})
	}
	defer fileContent.Close()

	uploadResp, err := cloudinaryClient.Upload.Upload(c.Context(), fileContent, uploader.UploadParams{
		Folder: "product_images",
	})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Gagal Untuk Mengupload Gambar",
		})
	}

	c.Locals("imageURL", uploadResp.SecureURL)
	return c.Next()
}
