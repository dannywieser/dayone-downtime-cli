package cmd

import (
	"dod/pkg/config"
	"dod/pkg/entry"
	"dod/pkg/utils"
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(auditGenreCmd)
	// TODO: how to not have this be a global var?
	auditGenreCmd.Flags().StringVarP(&ExportFile, "export", "e", "", "Path to export file (JSON zip format)")
	auditGenreCmd.MarkFlagRequired("export")
	auditGenreCmd.Flags().StringVarP(&Journal, "journal", "j", "", "Journal Name")
	auditGenreCmd.MarkFlagRequired("journal")
}

var auditGenreCmd = &cobra.Command{
	Use:   "audit-genre",
	Short: "Audit Genre Tags",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.LoadConfig()
		cobra.CheckErr(err)

		// extract and process export
		fmt.Println("Processing...")
		err = utils.Unzip(ExportFile, *cfg)
		cobra.CheckErr(err)

		entries, err := entry.RetrieveEntriesFromJson(Journal, *cfg)
		cobra.CheckErr(err)
		checkTags := entry.AuditGenreTags(entries, *cfg)

		if len(checkTags) > 0 {
			fmt.Println("The following tags exists that may need to be reviewed, as they are not defined as types or genres in the targeted entries.")

			for _, reviewTag := range checkTags {
				fmt.Printf("   - %s\n", reviewTag)
			}
		}
	},
}
