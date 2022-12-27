package entry

import (
	"dod/pkg/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Functions for working with multiple Day One Entries
func GetEntriesByTag(entries []model.Entry, tag string) []model.Entry {
	filteredEntries := []model.Entry{}
	for _, entry := range entries {
		if entryHasTag(entry, tag) {
			filteredEntries = append(filteredEntries, entry)
		}
	}

	return filteredEntries
}

func GetEntriesByTagAndYear(entries []model.Entry, tag string, year int) []model.Entry {
	filteredEntries := []model.Entry{}
	for _, entry := range entries {
		if entryHasTag(entry, tag) && entryDuringYear(entry, year) {
			filteredEntries = append(filteredEntries, entry)
		}
	}

	return filteredEntries
}

// Given: A journal name and a constant working directory
//   - Build a file name based on on `workingdir/JournalName.json`
//   - Retrieve all Day One Entries from the export
//   - Return the nested entries array to the caller
func RetrieveEntriesFromJson(journalName string, cfg model.Config) ([]model.Entry, error) {
	journalExport := filepath.Join(cfg.TmpDir, fmt.Sprintf("%s.json", journalName))
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
