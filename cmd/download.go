package cmd

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"

	"log"

	"ChanLoader/cmd/ui/textinput"
)

var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "",
	Long:  "",
	Run:   download,
}

type Save struct {
	Url        *textinput.Output
	Path       string
	ThreadName string
	ThreadID   string
	Board      string
}

type cfg struct {
	Url        string
	Path       string
	ThreadName string
	ThreadID   string
	Board      string
}

func init() {
	rootCmd.AddCommand(downloadCmd)

	downloadCmd.Flags().StringP("url", "u", "", "URL of the thread to save")
}

func download(cmd *cobra.Command, args []string) {

	flagURL := cmd.Flag("url").Value.String()

	save := Save{
		Url: &textinput.Output{},
	}

	cfg := &cfg{
		Url: flagURL,
	}

	tprogram_url := tea.NewProgram(textinput.InitialTextInputModel(save.Url, "Send the URL of the thread you want to save"))
	if _, err := tprogram_url.Run(); err != nil {
		cobra.CheckErr(err)
	}

	cfg.Url = save.Url.Output
	err := cmd.Flag("url").Value.Set(cfg.Url)
	if err != nil {
		log.Fatal("failed to set the url flag value \n", err)
	}

	split_url := strings.Split(cfg.Url, "/")
	fmt.Println(split_url)

	// ! Save Board and Thread ID info
	for i := 0; i < len(split_url); i++ {
		check := split_url[i]
		if check == "board.4chan.org" {
			save.Board = split_url[i+1]
		} else if check == "thread" {
			save.ThreadID = split_url[i+1]
		} else {
			continue
		}
	}

}
