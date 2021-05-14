package cmd

import (
	"fmt"
	"github.com/nicklpeterson/confluence-faker/settings"
	"github.com/nicklpeterson/confluence-faker/ui"
	"github.com/spf13/cobra"
	"log"
)

// pageCmd represents the page command
var pageCmd = &cobra.Command{
	Use:   "page",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("page called")
		pages, _ := cmd.Flags().GetInt("pages")
		space, _ := cmd.Flags().GetString("space")
		url, _ := cmd.Flags().GetString("url")
		addFakePages(pages, space, url)
	},
}

func init() {
	fakeCmd.AddCommand(pageCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pageCmd.PersistentFlags().String("foo", "", "A help for foo")
	pageCmd.PersistentFlags().String("url", "", "url of target confluence instance")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pageCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	pageCmd.Flags().Int("pages", 10, "number of pages to create in target space")
	pageCmd.Flags().String("space", "", "target space for new pages")
}

func addFakePages(numPages int, space string, url string) {
	userSettings, err := settings.GetSettings()
	selectedInstance := &settings.Instance{}
	if err != nil ||  len(userSettings.Instances) == 0 {
		log.Printf("%v", err)
		// Todo: Prompt user to enter confluence information
		selectedInstance = ui.PromptUserForConfluenceInstance("Please add a confluence instance.")
		settings.AddNewConfluenceInstance(selectedInstance)

	} else {
		selectedInstance, err = ui.SelectConfluenceInstance(userSettings)
		if err != nil {
			log.Panicf("%v\n", err)
		}
	}


}
