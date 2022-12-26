package cmd

import (
	"dod/display"
	"fmt"

	"github.com/spf13/cobra"
)

var ExportFile string

func init() {
	rootCmd.AddCommand(yirCmd)
	yirCmd.Flags().StringVarP(&ExportFile, "export", "e", "", "Path to export file (JSON zip format)")
	yirCmd.MarkFlagRequired("export")
}

var yirCmd = &cobra.Command{
	Use:   "yir",
	Short: "Year in Review",
	Run: func(cmd *cobra.Command, args []string) {
		display.Line(40)
		fmt.Printf("Generate Year in Review %s\n", ExportFile)
	},
}
