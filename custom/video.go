package custom

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func ConvertVideo(path string, format string) string {

	// Split path to find folder and file parts
	dir := filepath.Dir(path)
	file := filepath.Base(path)

	// Define the new folder for images
	convert_path := filepath.Join(dir, "Videos")
	err_convert_path := os.MkdirAll(convert_path, os.ModePerm)
	if err_convert_path != nil {
		log.Fatal("Error creating converting videos folder \n", err_convert_path)
	}
	// Get the file without its extension and create a new path
	filename_without_ext := strings.TrimSuffix(file, filepath.Ext(file))
	new_path := filepath.Join(convert_path, filename_without_ext+"."+format)

	// Check if the format is a video format
	if filepath.Ext(path) == ".mp4" || filepath.Ext(path) == ".avi" || filepath.Ext(path) == ".webm" {
		// Check video format and convert
		if filepath.Ext(path) != "."+format {
			// Convertion
			if format == "mp4" {
				err := ffmpeg.Input(path).Output(new_path, ffmpeg.KwArgs{"loglevel": "panic"}).OverWriteOutput().ErrorToStdOut().Run()
				if err != nil {
				}
			} else if format == "avi" {
				err := ffmpeg.Input(path).Output(new_path, ffmpeg.KwArgs{"loglevel": "panic"}).OverWriteOutput().ErrorToStdOut().Run()
				if err != nil {
				}
			} else if format == "webm" {
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
		// Move original file to Videos folder if no conversion needed
		err_rename := os.Rename(path, filepath.Join(convert_path, file))
		if err_rename != nil {
			log.Fatal("Error moving original file: ", err_rename)
		}
		return filepath.Join(convert_path, file)
	}
	return ""

}
