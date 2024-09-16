package custom

import (
	"math/rand"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/imroc/req"
	"github.com/schollz/progressbar/v3"
)

func NameFiles(path string, url string, nameformat string, id string, time int) string {
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
			times := strconv.Itoa(time)
			name += times + "."
		} else if nameformat == "alphabetic" {
			name += _alphabetic(time) + "."
		} else if nameformat == "id" {
			name += id + "."
		} else if nameformat == "random" {
			charset := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890"
			temp := make([]byte, 12)
			for i := range temp {
				temp[i] = charset[rand.Intn(len(charset))]
			}
			name += string(temp[:]) + "."
		} else if nameformat == "original" {
			name += file_split[len(file_split)-1]
		}

		new_path := filepath.Join(path, name+url_split[len(url_split)-1])
		working_url := "https:" + url

		r, _ := req.Get(working_url, req.DownloadProgress(progress))
		r.ToFile(new_path)
		return new_path
	} else {
		return "No media found in post: " + id
	}
}

func _alphabetic(n int) string {
	result := ""
	for n > 0 {
		remainder := (n - 1) % 26
		result = string('a'+remainder) + result
		n = (n - 1) / 26
	}
	return result
}
