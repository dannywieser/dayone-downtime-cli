package cmd

import (
	"dod/display"
	"dod/export"
	"fmt"
	"strings"

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
		var sb strings.Builder
		sb.WriteString(display.YearInReviewTitle(Year))
		fmt.Print(sb.String())

		// extract and process export
		fmt.Println("Please wait, processing export file")
		err := export.Unzip(ExportFile)
		cobra.CheckErr(err)

		entries, err := export.RetrieveEntriesFromJson(Journal)
		cobra.CheckErr(err)

		export.GetEntriesByTagAndYear(entries, "books", Year)
	},
}
