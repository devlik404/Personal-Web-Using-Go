package midleware

import (
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func MiddleFile(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		// //ambil value dari name input file
		file, err := c.FormFile("image")
		if err != nil {
			return c.String(http.StatusBadRequest, "Error uploading file")
		}

		// Buka file yang akan diunggah
		src, err := file.Open()
		if err != nil {
			return c.String(http.StatusInternalServerError, "Error opening file")
		}
		defer src.Close()
		//generate file unik
		// Buat UUID baru
		uuid := uuid.New()
		// Dapatkan ekstensi file dari nama file yang diunggah
		extension := filepath.Ext(file.Filename)
		// Cek apakah ekstensi file sesuai dengan yang diizinkan (jpg atau png)
		if extension != ".jpg" && extension != ".png" {
			return c.String(http.StatusBadRequest, "error exstension")
		}
		// package UUID sebagai nama file unik dan gabungkan dengan exstensinya
		uniqueFileName := uuid.String() + extension
		// Dapatkan ekstensi file dari nama file yang diunggah

		// Buat file baru di server untuk menyimpan gambar
		dstPath := filepath.Join("upload", uniqueFileName)
		dst, err := os.Create(dstPath)
		if err != nil {
			return c.String(http.StatusInternalServerError, "Error creating destination file")
		}
		defer dst.Close()

		// Salin isi file dari src ke dst (unggah)
		if _, err := io.Copy(dst, src); err != nil {
			return c.String(http.StatusInternalServerError, "Error file")
		}

		data := dst.Name()
		NameFile := data[7:]

		c.Set("DataFile", NameFile)

		return next(c)
	}
}
