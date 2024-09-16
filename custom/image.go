package custom

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func ConvertImage(path string, format string) string {

	// Split path to find folder and file parts
	dir := filepath.Dir(path)
	file := filepath.Base(path)

	// Define the new folder for images
	convert_path := filepath.Join(dir, "Images")
	err_convert_path := os.MkdirAll(convert_path, os.ModePerm)
	if err_convert_path != nil {
		log.Fatal("Error creating converting images folder \n", err_convert_path)
	}

	// Get the file without its extension and create a new path
	filename_without_ext := strings.TrimSuffix(file, filepath.Ext(file))
	new_path := filepath.Join(convert_path, filename_without_ext+"."+format)

	if filepath.Ext(path) == ".png" || filepath.Ext(path) == ".jpg" || filepath.Ext(path) == ".webp" || filepath.Ext(path) == ".jpeg" {
		// Check image format and convert
		if filepath.Ext(path) != "."+format {
			// Convertion
			if format == "png" {
				err := ffmpeg.Input(path).Output(new_path, ffmpeg.KwArgs{"loglevel": "panic"}).OverWriteOutput().ErrorToStdOut().Run()
				if err != nil {
				}
			} else if format == "jpeg" {
				err := ffmpeg.Input(path).Output(new_path, ffmpeg.KwArgs{"loglevel": "panic"}).OverWriteOutput().ErrorToStdOut().Run()
				if err != nil {
				}
			} else if format == "webp" {
				err := ffmpeg.Input(path).Output(new_path, ffmpeg.KwArgs{"loglevel": "panic"}).OverWriteOutput().ErrorToStdOut().Run()
				if err != nil {
				}
			}

			// Remove the original file
			err_remove := os.Remove(path)
			if err_remove != nil {
				log.Fatal("Error removing original file: ", err_remove)
			}
			return new_path
		}

		// Move original file to Images folder if no conversion needed
		err_rename := os.Rename(path, filepath.Join(convert_path, file))
		if err_rename != nil {
			log.Fatal("Error moving original file: ", err_rename)
		}
		return filepath.Join(convert_path, file)
	}
	return ""
}
