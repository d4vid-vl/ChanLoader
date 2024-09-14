package custom

import (
	"log"
	"os"
	"strings"

	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func ConvertImage(path string, format string) string {

	// * Create the new path
	file_path := strings.SplitAfter(path, "/")
	var convert_path string
	for i := 0; i < len(file_path)-1; i++ {
		convert_path += file_path[i]
	}
	convert_path += "Images/"
	err_convert_path := os.MkdirAll(convert_path, os.ModePerm)
	if err_convert_path != nil {
		log.Fatal("Error creating converting images folder \n", err_convert_path)
	}
	convert_path += file_path[len(file_path)-1] // New path with old file

	path_format := strings.SplitAfter(convert_path, ".")
	var new_path string
	for i := 0; i < len(path_format)-1; i++ {
		new_path += path_format[i]
	}
	new_path += format // New path with new file

	// * Check if the format is an image format
	if path_format[len(path_format)-1] == "png" || path_format[len(path_format)-1] == "jpg" || path_format[len(path_format)-1] == "webp" || path_format[len(path_format)-1] == "jpeg" {
		// * Check if the file format is the same as the converted one
		if path_format[len(path_format)-1] != format {

			// * Convertion
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

			// * Remove the original file
			if err := os.Remove(path); err != nil {
			}
			// TODO: Remove the println
			return new_path

		} else {
			// * Move original file to images
			if err := os.Rename(path, convert_path); err != nil {
			}
			return convert_path
		}
	} else if path_format[len(path_format)-1] == "gif" {
		if err := os.Rename(path, convert_path); err != nil {
		}
		return convert_path
	} else {
		return ""
	}
}
