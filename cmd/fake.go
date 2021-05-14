package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// fakeCmd represents the fake command
var fakeCmd = &cobra.Command{
	Use:   "fake",
	Short: "Generate fake generators in a specified settings instance.",
	Long: `The fake command is used to generate fake generators such as users, pages, or an entire space.`,
	Run: func(cmd * cobra.Command, args[] string) { fmt.Println("fake")},
}

func init() {
	RootCmd.AddCommand(fakeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fakeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fakeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
