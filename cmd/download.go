package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"

	"ChanLoader/cmd/ui/textinput"
)

var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "",
	Long:  "",
	Run:   download,
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
	Subject, Image, Name, Date, PostID, Message string
}

func init() {
	rootCmd.AddCommand(downloadCmd)

	downloadCmd.Flags().StringP("url", "u", "", "URL of the thread to save")
}

func download(cmd *cobra.Command, args []string) {

	// * Init
	flagURL := cmd.Flag("url").Value.String()
	config_path := "config/ConfigData.json"

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
	fmt.Println(cfg)

	// ! Create board and thread folder
	thread_path := cfg.Path + "/" + cfg.Board + "/" + cfg.ThreadID
	err_thread_path := os.MkdirAll(thread_path, os.ModePerm)
	if err_thread_path != nil {
		log.Fatal("Error creating thread folder \n", err_thread_path)
	}

	// * Scraping state
	test := scrapeurl(cfg.Url, cfg.ThreadID)
	for i := 0; i < len(test); i++ {
		// TODO: Name -> Image -> Video
		// TODO: Create name downloader and video converter
	}
}
