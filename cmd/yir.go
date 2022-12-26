// package cmd

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"os"

// 	"model"

// 	"github.com/fatih/color"
// )

// func main() {
// 	filename := "exports/Downtime.json"
// 	jsonFile, err := os.Open(filename)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Printf("successfully opened %s", filename)
// 	defer jsonFile.Close()

// 	byteValue, _ := ioutil.ReadAll(jsonFile)

// 	var entries model.Entries

// 	json.Unmarshal(byteValue, &entries)

// 	fmt.Println(len(entries.Entries))

// 	color.Cyan("Prints text in cyan.")
// }

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "hugo",
	Short: "Hugo is a very fast static site generator",
	Long: `A Fast and Flexible Static Site Generator built with
				  love by spf13 and friends in Go.
				  Complete documentation is available at https://gohugo.io/documentation/`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		fmt.Println("start")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
