package model

import "time"

type Entry struct {
	Text         string
	Tags         []string
	CreationDate time.Time
}

type Entries struct {
	Entries []Entry `json:"entries"`
}
