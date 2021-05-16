package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/nicklpeterson/confluence-faker/confluence"
	"github.com/nicklpeterson/confluence-faker/generators"
	"github.com/nicklpeterson/confluence-faker/logging"
	"github.com/nicklpeterson/confluence-faker/ui"
	"github.com/spf13/cobra"
	"os"
	"sync"
)

// pageCmd represents the page command
var pageCmd = &cobra.Command{
	Use:   "page",
	Short: "Generate pages in your Confluence Cloud Host",
	Long: `Generate random pages in a Confluence Cloud Host

		Examples:
			confluence-faker fake page --pages=20 #Generates 20 new pages
			confluence-faker fake page --space=PHP #Generate 10 new pages in the space PHP
			confluence-fakes fake page --url=example.atlassian.net #Generate 10 new pages in the specified confluence

			Currently all pages are created at the top level.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("page called")
		pages, _ := cmd.Flags().GetInt("pages")
		space, _ := cmd.Flags().GetString("space")
		url, _ := cmd.Flags().GetString("url")
		verbose, _ := cmd.Flags().GetBool("verbose")
		addFakePages(pages, space, url, logging.Logger{Verbose: verbose})
	},
}

func init() {
	fakeCmd.AddCommand(pageCmd)
	// Persistent Flags
	pageCmd.PersistentFlags().String("url", "", "Confluence host url")

	// local flags
	pageCmd.Flags().Int("pages", 10, "number of pages to create in target space")
	pageCmd.Flags().String("space", "", "target space for new pages")
}

func addFakePages(numPages int, space string, url string, logger logging.Logger) {
	selectedHost := ui.GetConfluenceHost(url)
	if space == "" {
		spaceList, err := selectedHost.GetSpaces()
		if err != nil {
			logger.Info("Unable to connect to " + url + "\nPlease check your settings and try again.")
			os.Exit(-1)
		}
		items := make([]string, len(*spaceList))
		for i, space := range *spaceList {
			items[i] = space.Name
		}
		index, _, err := ui.SelectFromList(items, "Please select a target space")
		if err != nil {
			logger.Info("Unable to connect to " + url + "\nPlease check your settings and try again.")
			os.Exit(-1)
		}
		space = (*spaceList)[index].Key
	}

	fakePageArray, err := generators.NewFakePageArray(space, numPages)
	if err != nil {
		logger.Info("Unable to generate data: %v\n", err)
	}

	spinner := ui.NewSpinner("Uploading Pages")

	if !logger.Verbose {
		spinner.Start()
	}
	logger.Debug("Uploading Pages ...\n")

	var wg sync.WaitGroup

	for index, page := range fakePageArray {
		jsonBody, err := json.Marshal(page)
		if err == nil {
			wg.Add(1)
			go worker(index, &wg, &logger, selectedHost, jsonBody)
		} else {
			logger.Info("Unable to create page %v, skipping", index)
		}
	}

	wg.Wait()
	if !logger.Verbose {
		spinner.Stop()
	}
	logger.Info("%d pages were added successfully!\n", numPages)
}

func worker(id int, wg *sync.WaitGroup, logger *logging.Logger, host *confluence.Host, body []byte) {
	defer wg.Done()
	status, _, err := host.Post("/wiki/rest/api/content", body)
	logger.Debug("Worker %d: http response: %v\n", id, status)
	if err != nil {
		logger.Info("Unable to create page %v, skipping", id)
	}
}
