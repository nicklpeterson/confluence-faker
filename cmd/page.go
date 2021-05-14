package cmd

import (
	"fmt"
	"github.com/nicklpeterson/confluence-faker/generators"
	"github.com/nicklpeterson/confluence-faker/ui"
	"github.com/spf13/cobra"
	"log"
	"os"
)

// pageCmd represents the page command
var pageCmd = &cobra.Command{
	Use:   "page",
	Short: "Generate pages in your Confluence Cloud Instance",
	Long: `Generate random pages in a Confluence Cloud Instance

		Examples:
			confluence-faker fake page --pages=20 #Generates 20 new pages
			confluence-faker fake page --space=PHP #Generate 10 new pages in the space PHP
			confluence-fakes fake page --url=example.atlassian.net #Generate 10 new pages in the specified confluence instance

			Currently all pages are created at the top level.`,
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
	// Persistent Flags
	pageCmd.PersistentFlags().String("url", "", "url of target confluence instance")

	// local flags
	pageCmd.Flags().Int("pages", 10, "number of pages to create in target space")
	pageCmd.Flags().String("space", "", "target space for new pages")
}

func addFakePages(numPages int, space string, url string) {
	selectedInstance := ui.GetConfluenceInstance(url)
	if space == "" {
		spaceList, err := selectedInstance.GetSpaces()
		if err != nil {
			log.Printf("Unable to connect to " + url + "\nPlease check your settings and try again.")
			os.Exit(-1)
		}
		items := make([]string, len(*spaceList))
		for i, space := range *spaceList {
			items[i] = space.Name
		}
		index, _, err := ui.SelectFromList(items, "Please select a target space")
		if err != nil {
			log.Printf("Unable to connect to " + url + "\nPlease check your settings and try again.")
			os.Exit(-1)
		}
		space = (*spaceList)[index].Key
	}

	fakePageArray, err := generators.NewFakePageArray(space, numPages)
	if err != nil {
		log.Printf("Unable to generate data: %v\n", err)
	}
	log.Printf("%v\n", fakePageArray)
}
