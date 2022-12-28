package entry

import (
	"dod/pkg/config"
	"dod/pkg/model"
	"dod/pkg/utils"
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

// Given a set of entries, this function will look through the tags for those entries, and:
//   - collect all the tags currently defined in the configuration for various purposes
//   - locate tags that exist that are not covered by this group
//   - report on those tags as _possibly_ being genres not included in config
func AuditGenreTags(entries []model.Entry, cfg model.Config) []string {
	var possibleGenres []string
	for _, checkEntry := range entries {
		for _, checkTag := range checkEntry.Tags {
			inConfig := config.TagExistsInConfig(checkTag, cfg)
			if !inConfig {
				possibleGenres = append(possibleGenres, checkTag)
			}
		}
	}
	return utils.RemoveDuplicates(possibleGenres)
}
