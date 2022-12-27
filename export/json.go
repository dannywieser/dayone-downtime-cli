package export

import (
	"dod/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func RetrieveEntriesFromJson(journalName string) ([]model.Entry, error) {
	journalExport := filepath.Join(exportWorkingDir, fmt.Sprintf("%s.json", journalName))
	jsonFile, err := os.Open(journalExport)
	if err != nil {
		return nil, err
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var entries model.Entries
	err = json.Unmarshal(byteValue, &entries)
	if err != nil {
		return nil, err
	}
	return entries.Entries, nil
}
