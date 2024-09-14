package custom

import (
	"log"
	"os"
	"strings"

	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func ConvertVideo(path string, format string) string {

	// * Create the new path
	file_path := strings.SplitAfter(path, "/")
	var convert_path string
	for i := 0; i < len(file_path)-1; i++ {
		convert_path += file_path[i]
	}
	convert_path += "Videos/"
	err_convert_path := os.MkdirAll(convert_path, os.ModePerm)
	if err_convert_path != nil {
		log.Fatal("Error creating converting videos folder \n", err_convert_path)
	}
	convert_path += file_path[len(file_path)-1] // New path with old file

	path_format := strings.SplitAfter(convert_path, ".")

	var new_path string
	for i := 0; i < len(path_format)-1; i++ {
		new_path += path_format[i]
	}
	new_path += format // New path with new file

	// * Check if the format is a video format
	if path_format[len(path_format)-1] == "mp4" || path_format[len(path_format)-1] == "avi" || path_format[len(path_format)-1] == "webm" {
		if path_format[len(path_format)-1] != format {
			// * Convertion
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
	} else {
		return ""
	}
}
