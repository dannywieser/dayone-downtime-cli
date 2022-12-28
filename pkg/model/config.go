package model

type Config struct {
	ExportFile      string
	TargetJournal   string
	TmpDir          string
	TypeTags        []string
	ReviewEntryTags []string
	GenreTags       []string
	FavoriteTag     string
	DidNotFinishTag string
	Ratings         []string
}
