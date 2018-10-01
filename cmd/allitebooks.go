package cmd

import (
	"os"

	"github.com/PGo-Projects/bore/internal/allitebooks"
	"github.com/PGo-Projects/bore/internal/allitebooks/scraper"
	"github.com/PGo-Projects/bore/internal/allitebooks/utils"
	tm "github.com/buger/goterm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var homepage = "http://www.allitebooks.com"
var allitebooksCmd = &cobra.Command{
	Use:   "allitebooks",
	Short: "Download pdf of books or search",
	Long:  "Search through www.allitebooks.com for specific books and/or download them",
	Args:  cobra.ExactArgs(0),
	Run:   allitebooksBorer,
}

func init() {
	allitebooksCmd.Flags().Int("startpage", 0, "the page to start searching/downloading from")
	allitebooksCmd.Flags().String("starturl", "", "the url to start searching/downloading from")

	viper.BindPFlag("allitebooks-startpage", allitebooksCmd.Flags().Lookup("startpage"))
	viper.BindPFlag("allitebooks-starturl", allitebooksCmd.Flags().Lookup("starturl"))

	startPage, err := scraper.GetTotalPages(homepage)
	if err != nil {
		utils.DisplayMessage("Unable to get value for start page, did not pass health check", tm.RED)
		os.Exit(1)
	}
	startURL, err := scraper.GetLastURLForPage(homepage, startPage)
	if err != nil {
		utils.DisplayMessage("Unable to get value for start url, did not pass health check", tm.RED)
		os.Exit(1)
	}
	viper.SetDefault("allitebooks-startpage", startPage)
	viper.SetDefault("allitebooks-starturl", startURL)

	RootCmd.AddCommand(allitebooksCmd)
}

func allitebooksBorer(cmd *cobra.Command, args []string) {
	startPage := viper.GetInt("allitebooks-startpage")
	startURL := viper.GetString("allitebooks-starturl")
	if cmd.Flags().Changed("startpage") && !cmd.Flags().Changed("starturl") {
		startURL = getNewStartURL(startPage)
	}
	borer := allitebooks.
		New().
		WithStartPage(startPage).
		WithStartURL(startURL).
		Build()
	borer.GetAll()
}

func getNewStartURL(startPage int) string {
	url, err := scraper.GetLastURLForPage(homepage, startPage)
	if err != nil {
		utils.DisplayMessage("Unable to retrieve the starting url", tm.RED)
		os.Exit(1)
	}
	viper.Set("allitebooks-starturl", url)
	return url
}
