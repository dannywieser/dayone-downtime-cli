package model

type Entry struct {
	Text string
}

type Entries struct {
	Entries []Entry `json:"entries"`
}
