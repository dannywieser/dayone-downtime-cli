package export

import (
	"dod/model"
	"fmt"
	"time"
)

func entryHasTag(entry model.Entry, findTag string) bool {
	for _, tag := range entry.Tags {
		if tag == findTag {
			return true
		}
	}
	return false
}

func entryDuringYear(entry model.Entry, year int) bool {
	startYear := time.Date(year, 1, 1, 0, 0, 0, 0, time.Local)
	endYear := time.Date(year, 12, 31, 12, 59, 59, 0, time.Local)

	return entry.CreationDate.After(startYear) && entry.CreationDate.Before(endYear)
}

func GetEntriesByTagAndYear(entries []model.Entry, tag string, year int) []model.Entry {
	for _, entry := range entries {
		if entryHasTag(entry, tag) && entryDuringYear(entry, year) {
			fmt.Println(entry.CreationDate)
		}
	}

	return nil
}
