package export

import (
	"dod/pkg/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Given: A journal name and a constant working directory
//   - Build a file name based on on `workingdir/JournalName.json`
//   - Retrieve all Day One Entries from the export
//   - Return the nested entries array to the caller
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
