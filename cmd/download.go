package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"

	"ChanLoader/cmd/ui/textinput"
	"ChanLoader/custom"
)

var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Download a thread",
	Long: `Download a thread by simply:
- Calling the command
- Pasting the link of thread you want
- Declare if you want a verbose or "quiet" download`,
	Run: download,
}

// ? Structs to save configurations of the program
type Save struct {
	Url *textinput.Output
}
type cfg struct {
	Url, Path, ThreadName, ThreadID, Board string
}

// ? Struct to save each individual post in the thread
type Post struct {
	OP                                          bool
	Subject, Media, Name, Date, PostID, Message string
}

func init() {
	rootCmd.AddCommand(downloadCmd)

	downloadCmd.Flags().StringP("url", "u", "", "URL of the thread to save")
}

func download(cmd *cobra.Command, args []string) {
	// Get the current working directory
	currentPath, err_c := os.Getwd()
	if err_c != nil {
		log.Fatal(err_c)
	}

	// Get the absolute path
	absPath, err_path := filepath.Abs(currentPath)
	if err_path != nil {
		log.Fatal(err_path)
	}
	// * Init
	flagURL := cmd.Flag("url").Value.String()
	config_path := absPath + "/config/"
	err_convert_path := os.MkdirAll(config_path, os.ModePerm)
	if err_convert_path != nil {
		log.Fatal("Error creating converting config folder \n", err_convert_path)
	}
	config_path += "ConfigData.json"

	save := Save{
		Url: &textinput.Output{},
	}

	cfg := &cfg{
		Url: flagURL,
	}

	// ! Save URL in local variable
	tprogram_url := tea.NewProgram(textinput.InitialTextInputModel(save.Url, "Send the URL of the thread you want to save"))
	if _, err := tprogram_url.Run(); err != nil {
		cobra.CheckErr(err)
	}

	cfg.Url = save.Url.Output
	err := cmd.Flag("url").Value.Set(cfg.Url)
	if err != nil {
		log.Fatal("failed to set the url flag value \n", err)
	}

	// * Check if it is a valid 4Chan thread link
	if !strings.Contains(cfg.Url, "boards.4chan.org") {
		log.Fatal("Url given is not a valid 4chan thread. (Invalid link)")
	} else if !strings.Contains(cfg.Url, "thread") {
		log.Fatal("Url given is not a valid 4chan thread. (Not a thread)")
	}

	split_url := strings.Split(cfg.Url, "/")

	// ! Save Board and Thread ID info
	for i := 0; i < len(split_url); i++ {
		check := split_url[i]
		if check == "boards.4chan.org" {
			cfg.Board = split_url[i+1]
		} else if check == "thread" {
			cfg.ThreadID = split_url[i+1]
		} else {
			continue
		}
	}

	// ! Read and write config path in local variable
	content, err := os.ReadFile(config_path)
	if err != nil {
		log.Fatal("Couldn't get file path correctly \n", err)
	}
	var config Config
	if err := json.Unmarshal(content, &config); err != nil {
		fmt.Println("Couldn't read json file", err)
	}
	cfg.Path = config.Path

	// ! Create board and thread folder
	posts := scrapeurl(cfg.Url, cfg.ThreadID)
	cfg.ThreadName = strings.ReplaceAll(posts[0].Subject, "/", "!")
	var thread_path string
	if cfg.ThreadName != "" {
		thread_path = cfg.Path + "/" + cfg.Board + "/" + cfg.ThreadID + " - " + cfg.ThreadName
	} else {
		thread_path = cfg.Path + "/" + cfg.Board + "/" + cfg.ThreadID + " - " + "Untitled Thread"
	}
	err_thread_path := os.MkdirAll(thread_path, os.ModePerm)
	if err_thread_path != nil {
		log.Fatal("Error creating thread folder \n", err_thread_path)
	}

	// * Scraping state
	// TODO: Make useful name types
	var files []string
	for i := 0; i < len(posts)-1; i++ {
		post := posts[i]
		file := custom.NameFiles(thread_path, post.Media, config.Name.String(), post.PostID, i+1)
		files = append(files, file)
	}

	// * Convert images and videos
	var images []string
	for i := 0; i < len(files); i++ {
		file := files[i]
		image := custom.ConvertImage(file, config.IExtension.String())
		images = append(images, image)
	}
	var videos []string
	for i := 0; i < len(files); i++ {
		file := files[i]
		video := custom.ConvertVideo(file, config.VExtension.String())
		videos = append(videos, video)
	}

	fmt.Println("All files have been downloaded successfully!")
}
