package custom

import (
	"io"
	"log"
	"strings"

	"github.com/sunshineplan/imgconv"
)

func ConvertImage(path string, format string) string {
	src, err := imgconv.Open(path)
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}

	path_format := strings.SplitAfter(path, ".")

	// * Check if the file format is the same as the converted one
	if path_format[len(path_format)] != format {
		// * Convertion
		if format == "png" {
			if err := imgconv.Write(io.Discard, src, &imgconv.FormatOption{Format: imgconv.PNG}); err != nil {
				log.Fatalf("failed to write image: %v", err)
			}
		} else if format == "jpeg" {
			if err := imgconv.Write(io.Discard, src, &imgconv.FormatOption{Format: imgconv.JPEG}); err != nil {
				log.Fatalf("failed to write image: %v", err)
			}
		} else if format == "webp" {
			if err := imgconv.Write(io.Discard, src, &imgconv.FormatOption{Format: imgconv.PNG}); err != nil {
				log.Fatalf("failed to write image: %v", err)
			}
		}
	}
	var new_path string
	for i := 0; i < len(path_format)-1; i++ {
		new_path += path_format[i]
	}
	new_path += format
	return new_path

	// TODO: Verify if is working
}
