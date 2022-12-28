package review

import (
	"dod/pkg/entry"
	"dod/pkg/model"
	"fmt"
	"sort"
	"strings"
)

func reportGenre(entries []model.Entry, cfg model.Config) string {
	var sb strings.Builder
	var genres []model.GenreStats
	for _, genreTag := range cfg.GenreTags {
		genreEntries := entry.GetEntriesByTag(entries, genreTag)
		if len(genreEntries) > 0 {
			genres = append(genres, model.GenreStats{Genre: genreTag, Count: len(genreEntries)})
		}
	}

	if len(genres) > 0 {
		sort.SliceStable(genres, func(i, j int) bool {
			return genres[i].Count > genres[j].Count
		})

		sb.WriteString(fmt.Sprintf("**Top Genre: %s**\n", genres[0].Genre))
		for _, genreItem := range genres {
			sb.WriteString(fmt.Sprintf(" - %s: %v\n", genreItem.Genre, genreItem.Count))
		}
	}
	return sb.String()
}

func reportFavorites(entries []model.Entry, cfg model.Config) string {
	var sb strings.Builder
	favoriteEntries := entry.GetEntriesByTag(entries, cfg.FavoriteTag)
	if len(favoriteEntries) > 0 {
		sb.WriteString("\n**Favorites**\n")
	}
	for _, faveEntry := range favoriteEntries {
		sb.WriteString(fmt.Sprintf(" - *%s*\n", entry.GetEntryTitle(faveEntry)))
	}

	return sb.String()
}

func reportDidNotFinish(entries []model.Entry, cfg model.Config) string {
	var sb strings.Builder
	didNotFinishEntries := entry.GetEntriesByTag(entries, cfg.DidNotFinishTag)
	if len(didNotFinishEntries) > 0 {
		sb.WriteString("\n**Did Not Finish**\n")
	}
	for _, dnfEntry := range didNotFinishEntries {
		sb.WriteString(fmt.Sprintf(" - *%s*\n", entry.GetEntryTitle(dnfEntry)))
	}

	return sb.String()
}

func CreateReviewBody(entries []model.Entry, cfg model.Config, year int) string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("# 🥳 %v 🥳 : Year In Review\n", year))
	for _, typeTag := range cfg.TypeTags {
		entries := entry.GetEntriesByTagAndYear(entries, typeTag, year)
		sb.WriteString(fmt.Sprintf("# %s \n", typeTag))
		sb.WriteString(fmt.Sprintf("**Total: %d**\n", len(entries)))
		sb.WriteString(reportGenre(entries, cfg))
		sb.WriteString(reportFavorites(entries, cfg))
		sb.WriteString(reportDidNotFinish(entries, cfg))
		sb.WriteString("---\n")
	}
	return sb.String()
}
