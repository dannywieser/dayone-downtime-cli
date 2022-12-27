package entry

import (
	"dod/pkg/model"
	"fmt"
)

// Functions for working with multiple Day One Entries

func GetEntriesByTagAndYear(entries []model.Entry, tag string, year int) []model.Entry {
	for _, entry := range entries {
		if entryHasTag(entry, tag) && entryDuringYear(entry, year) {
			fmt.Println(entry.CreationDate)
		}
	}

	return nil
}
