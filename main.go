package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/fatih/color"
)

type Entry struct {
	Text string
}

type Entries struct {
	Entries []Entry `json:"entries"`
}

func main() {
	filename := "exports/Downtime.json"
	jsonFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("successfully opened %s", filename)
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var entries Entries

	json.Unmarshal(byteValue, &entries)

	fmt.Println(len(entries.Entries))

	color.Cyan("Prints text in cyan.")
}
