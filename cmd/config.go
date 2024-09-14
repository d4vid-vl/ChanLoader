package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"

	"ChanLoader/cmd/flags"
	"ChanLoader/cmd/steps"
	"ChanLoader/cmd/ui/multioption"
	"ChanLoader/cmd/ui/textinput"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

type Config struct {
	Path       string
	IExtension flags.IExtension
	VExtension flags.VExtension
	Name       flags.Name
}

type forjson struct {
	Path       string `json:"path"`
	IExtension string `json:"iextension"`
	VExtension string `json:"vextension"`
	Name       string `json:"name"`
}

type Options struct {
	Path       *textinput.Output
	IExtension *multioption.Selection
	VExtension *multioption.Selection
	Name       *multioption.Selection
}

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configuration for files",
	Long:  `Configurate files path, name, extension, etc...`,
	Run:   configurate,
}

func init() {
	var flagIExtension flags.IExtension
	var flagVExtension flags.VExtension
	var flagName flags.Name

	rootCmd.AddCommand(configCmd)

	configCmd.Flags().StringP("path", "p", "", "The path where the files will be saved")
	configCmd.Flags().VarP(&flagIExtension, "image", "i", fmt.Sprintf("File image extension to use. Allowed values: %s", strings.Join(flags.AllowedIExtensions, ", ")))
	configCmd.Flags().VarP(&flagVExtension, "video", "v", fmt.Sprintf("File video extension to use. Allowed values: %s", strings.Join(flags.AllowedVExtensions, ", ")))
	configCmd.Flags().VarP(&flagName, "name", "n", fmt.Sprintf("Categorized file names to use. Allowed values: %s", strings.Join(flags.AllowedVExtensions, ", ")))

}

func configurate(cmd *cobra.Command, args []string) {
	var separator string
	var configDataFile string
	if runtime.GOOS == "windows" {
		separator = "\\"
		configDataFile = separator + "config" + separator + "ConfigData.json"
	} else {
		separator = "/"
		configDataFile = "config" + separator + "ConfigData.json"
	}

	flagPath := cmd.Flag("path").Value.String()
	flagIExtension := flags.IExtension(cmd.Flag("image").Value.String())
	flagVExtension := flags.VExtension(cmd.Flag("video").Value.String())
	flagName := flags.Name(cmd.Flag("name").Value.String())

	steps := steps.InitSteps(flagIExtension, flagVExtension, flagName)

	step_iextension := steps.Steps["i_extension"]
	step_vextension := steps.Steps["v_extension"]
	step_name := steps.Steps["name"]

	options := Options{
		Path:       &textinput.Output{},
		IExtension: &multioption.Selection{},
		VExtension: &multioption.Selection{},
		Name:       &multioption.Selection{},
	}

	config := &Config{
		Path:       flagPath,
		IExtension: flagIExtension,
		VExtension: flagVExtension,
		Name:       flagName,
	}

	// ! File save path config
	tprogram_path := tea.NewProgram(textinput.InitialTextInputModel(options.Path, "What is the path to save the files? (Copy the whole path link)"))
	if _, err := tprogram_path.Run(); err != nil {
		cobra.CheckErr(err)
	}
	if _, err_path := os.Stat(options.Path.Output); err_path != nil {
		log.Fatal("Path given is not valid \n", err_path)
	}

	config.Path = options.Path.Output
	err := cmd.Flag("path").Value.Set(config.Path)
	if err != nil {
		log.Fatal("failed to set the path flag value \n", err)
	}

	// ! File image extension config
	tprogram_iextension := tea.NewProgram(multioption.InitialModelMulti(step_iextension.Options, options.IExtension, step_iextension.Headers))
	if _, err := tprogram_iextension.Run(); err != nil {
		cobra.CheckErr(err)
	}
	step_iextension.Field = options.IExtension.Choice
	config.IExtension = flags.IExtension(strings.ToLower(options.IExtension.Choice))
	err_i := cmd.Flag("image").Value.Set(config.IExtension.String())
	if err_i != nil {
		log.Fatal("failed to set the image flag value \n", err_i)
	}

	// ! File video extension config
	tprogram_vextension := tea.NewProgram(multioption.InitialModelMulti(step_vextension.Options, options.VExtension, step_vextension.Headers))
	if _, err := tprogram_vextension.Run(); err != nil {
		cobra.CheckErr(err)
	}
	step_vextension.Field = options.VExtension.Choice
	config.VExtension = flags.VExtension(strings.ToLower(options.VExtension.Choice))
	err_v := cmd.Flag("video").Value.Set(config.VExtension.String())
	if err_v != nil {
		log.Fatal("failed to set the image flag value \n", err_v)
	}

	// ! File name config
	tprogram_name := tea.NewProgram(multioption.InitialModelMulti(step_name.Options, options.Name, step_name.Headers))
	if _, err := tprogram_name.Run(); err != nil {
		cobra.CheckErr(err)
	}
	step_name.Field = options.Name.Choice
	config.Name = flags.Name(strings.ToLower(options.Name.Choice))
	err_n := cmd.Flag("name").Value.Set(config.Name.String())
	if err_n != nil {
		log.Fatal("failed to set the name flag value \n", err_n)
	}

	ConfigData := forjson{
		Path:       config.Path,
		IExtension: config.IExtension.String(),
		VExtension: config.VExtension.String(),
		Name:       config.Name.String(),
	}

	bytes, _ := json.MarshalIndent(ConfigData, "", "  ")

	// Revisa si el archivo JSON está creado
	if _, err := os.Stat(configDataFile); err == nil { // En caso de estar creado
		// Abre el archivo y borra los contenidos de este
		_, err := os.OpenFile(configDataFile, os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
			fmt.Println("Error opening the JSON: ", err)
		}
		// Escribe la información en el JSON
		if err := os.WriteFile(configDataFile, bytes, 0644); err != nil {
			fmt.Println("Error writing in the JSON: ", err)
			return
		}
	} else { // En caso de no estar creado
		// Crea el archivo
		file, err := os.Create(configDataFile)
		if err != nil {
			fmt.Println("Error creating the JSON: ", err)
			return
		}
		defer file.Close()

		// Escribe la información en el JSON
		if err := os.WriteFile(configDataFile, bytes, 0644); err != nil {
			fmt.Println("Error writing in the JSON: ", err)
			return
		}
	}

	fmt.Println("Config saved in:", configDataFile)
}
