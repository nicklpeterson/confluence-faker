package cmd

import (
	"github.com/nicklpeterson/confluence-faker/generators"
	"github.com/nicklpeterson/confluence-faker/logging"
	"github.com/nicklpeterson/confluence-faker/ui"
	"os"
	"sync"

	"github.com/spf13/cobra"
)

// blogCmd represents the blog command
var blogCmd = &cobra.Command{
	Use:   "blog",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		blogs, _ := cmd.Flags().GetInt("blogs")
		space, _ := cmd.Flags().GetString("space")
		url, _ := cmd.Flags().GetString("url")
		verbose, _ := cmd.Flags().GetBool("verbose")
		addFakeBlogs(blogs, space, url, logging.Logger{Verbose: verbose})
	},
}

func init() {
	fakeCmd.AddCommand(blogCmd)
	// Persistent Flags
	blogCmd.PersistentFlags().String("url", "", "Confluence host url")

	// local flags
	blogCmd.Flags().Int("blogs", 10, "number of pages to create in target space")
	blogCmd.Flags().String("space", "", "target space for new pages")
}

func addFakeBlogs(numBlogs int, space string, url string, logger logging.Logger) {
	host := ui.GetConfluenceHost(url)
	if space == "" {
		var err error
		space, err = ui.GetConfluenceSpace(host)
		if err != nil {
			logger.Info("Unable to connect to " + url + "\nPlease check your settings and try again.")
			os.Exit(-1)
		}
	}
	spinner := ui.NewSpinner("Generating and Uploading Blog Posts")
	if !logger.Verbose {
		spinner.Start()
	}

	var wg sync.WaitGroup
	for i := 0; i < numBlogs; i++ {
		wg.Add(1)
		go generators.ContentWorker(i, &wg, &logger, host, space, generators.NewFakeBlog)
	}
	wg.Wait()

	if spinner.Active() {
		spinner.Stop()
	}
	logger.Info("Done adding blogs!\n")
}

