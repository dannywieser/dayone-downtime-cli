package entry

import (
	"dod/pkg/model"
	"strings"
	"time"
)

// Functions relating to working with a Day One `Entry`

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

func GetEntryTitle(entry model.Entry) string {
	titleStart := strings.Index(entry.Text, "#")
	titleEnd := strings.Index(entry.Text, "\n")
	titleLen := titleEnd - titleStart
	return entry.Text[titleStart+2 : titleLen]
}
