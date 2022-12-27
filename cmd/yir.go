package cmd

import (
	"dod/pkg/config"
	"dod/pkg/entry"
	"dod/pkg/utils"
	util "dod/pkg/utils"
	"fmt"

	"github.com/spf13/cobra"
)

var ExportFile string
var Year int
var Journal string

func init() {
	rootCmd.AddCommand(yirCmd)
	yirCmd.Flags().StringVarP(&ExportFile, "export", "e", "", "Path to export file (JSON zip format)")
	yirCmd.MarkFlagRequired("export")

	yirCmd.Flags().IntVarP(&Year, "year", "y", 0, "Year")
	yirCmd.MarkFlagRequired("year")

	yirCmd.Flags().StringVarP(&Journal, "journal", "j", "", "Journal Name")
	yirCmd.MarkFlagRequired("journal")
}

var yirCmd = &cobra.Command{
	Use:   "yir",
	Short: "Year in Review",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.LoadConfig()
		cobra.CheckErr(err)

		// extract and process export
		fmt.Println("Processing...")
		err = utils.Unzip(ExportFile, *cfg)
		cobra.CheckErr(err)

		entries, err := entry.RetrieveEntriesFromJson(Journal, *cfg)
		cobra.CheckErr(err)

		fmt.Print(util.YearInReviewTitle(Year))

		for _, typeTag := range cfg.TypeTags {
			entries := entry.GetEntriesByTagAndYear(entries, typeTag, Year)
			fmt.Print(utils.TypeTitle(typeTag, len(entries)))
		}
	},
}
