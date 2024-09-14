package custom

import (
	"fmt"
	"strings"

	"github.com/imroc/req"
	"github.com/schollz/progressbar/v3"
)

func NameFiles(path string, url string, nameformat string, id string) string {
	if url != "" {

		bar := progressbar.DefaultBytes(-1, "Downloading...")
		progress := func(current, total int64) {
			//fmt.Println(float32(current)/float32(total)*100, "%")
			bar.ChangeMax64(total)
			bar.Set64(current)
		}
		name := "/"
		url_split := strings.SplitAfter(url, ".")
		file_split := strings.Split(url_split[len(url_split)-2], "/")

		if nameformat == "numeric" {
			name += url_split[len(url_split)-2]
		} else if nameformat == "alphabetic" {
			name += url_split[len(url_split)-2]
		} else if nameformat == "id" {
			name += id + "."
		} else if nameformat == "random" {
			name += url_split[len(url_split)-2]
		} else if nameformat == "original" {
			name += file_split[len(file_split)-1]
		}

		new_path := path + name + url_split[len(url_split)-1]
		working_url := "https:" + url

		r, _ := req.Get(working_url, req.DownloadProgress(progress))
		r.ToFile(new_path)
		fmt.Println("File saved in: " + new_path)
		return new_path
	} else {
		fmt.Println("No media found in post: " + id)
		return "No media found in post: " + id
	}
}
