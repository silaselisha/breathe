package main

import (
	"strings"
)

func checkPlagiarism(playlist tracks) []string {
	var plagiarizedTracks []string
	trackMap := make(map[string]bool)

	for _, item := range playlist.Tracks.Items {
		key := strings.ToLower(item.Track.Album.Name + " - " + item.Track.Artists[0].Name)
		if trackMap[key] {
			plagiarizedTracks = append(plagiarizedTracks, item.Track.Album.Name+" by "+item.Track.Artists[0].Name)
		} else {
			trackMap[key] = true
		}
	}

	return plagiarizedTracks
}
