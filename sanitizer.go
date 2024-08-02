package main

import (
	"strings"
	"unicode"
)

func sanitizePlaylist(playlist tracks) tracks {
    for i, item := range playlist.Tracks.Items {
        playlist.Tracks.Items[i].Track.Album.Name = sanitizeString(item.Track.Album.Name)
        for j, artist := range item.Track.Artists {
            playlist.Tracks.Items[i].Track.Artists[j].Name = sanitizeString(artist.Name)
        }
        playlist.Tracks.Items[i].Track.Album.Name = sanitizeString(item.Track.Album.Name)
    }
    return playlist
}

func sanitizeString(s string) string {
    s = strings.TrimSpace(s)
    s = strings.Title(strings.ToLower(s))
    return strings.Map(func(r rune) rune {
        if unicode.IsLetter(r) || unicode.IsNumber(r) || unicode.IsSpace(r) {
            return r
        }
        return -1
    }, s)
}
