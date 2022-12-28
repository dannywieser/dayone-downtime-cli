package cmd

import (
	"dod/pkg/config"
	"dod/pkg/entry"
	"dod/pkg/review"
	"dod/pkg/utils"
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(yirCmd)
}

var yirCmd = &cobra.Command{
	Use:   "yir",
	Short: "Year in Review",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.LoadConfig()
		cobra.CheckErr(err)

		year, err := strconv.Atoi(args[0])
		cobra.CheckErr(err)

		// extract and process export
		fmt.Print(utils.YearInReviewTitle(year))
		fmt.Printf("  - processing `%s`...\n", cfg.ExportFile)

		err = utils.Unzip(*cfg)
		cobra.CheckErr(err)

		entries, err := entry.RetrieveEntriesFromJson(*cfg)
		cobra.CheckErr(err)

		reviewString := review.CreateReviewBody(entries, *cfg, year)
		err = entry.CreateEntry(cfg.TargetJournal, strings.Join(cfg.ReviewEntryTags, " "), reviewString)
		cobra.CheckErr(err)

		fmt.Print(review.CreateCliReport(entries, *cfg, year))
	},
}
