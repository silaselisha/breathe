package main

type request_params struct {
	GrantType    string `json:"grant_type"`
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type access_token_params struct {
	AccessToken string `json:"access_token"`
	Token_Type  string `json:"token_type"`
	ExpiresIn   int32  `json:"expires_in"`
}

type artist struct {
	ExternalUrls struct {
		Spotify string `json:"spotify"`
	} `json:"external_urls"`
	Followers struct {
		Href  string `json:"href"`
		Total int32  `json:"total"`
	} `json:"followers"`
	Id     string   `json:"id"`
	Geners []string `json:"geners"`
	Href   string   `json:"href"`
	Images []struct {
		Url    string `json:"url"`
		Height int32  `json:"height"`
		Width  string `json:"width"`
	}
	Name       string `json:"name"`
	Popularity string `json:"popularity"`
	Type       string `json:"type"`
	Uri        string `json:"uri"`
}

type album struct {
	AlbumType   string `json:"album_type"`
	TotalTracks int32  `json:"total_tracks"`
	// AvailableMarkets []string `json:"available_markets"`
	ExternalUrls struct {
		Spotify string `json:"spotify"`
	} `json:"external_urls"`
	Href   string `json:"href"`
	Id     string `json:"id"`
	Images []struct {
		Url    string `json:"url"`
		Height int32  `json:"height"`
		Width  int32  `json:"width"`
	} `json:"images"`
	Name                 string `json:"name"`
	ReleaseDate          string `json:"release-date"`
	ReleaseDatePrecision string `json:"release_date_precision"`
	Restriction          struct {
		Reason string `json:"reason"`
	} `json:"restriction"`
	Type    string `json:"type"`
	Uri     string `json:"uri"`
	Artists []struct {
		ExternalUrls struct {
			Spotify string `json:"spotify"`
		} `json:"external_urls"`
		Href string `json:"href"`
		Id   string `json:"id"`
		Type string `json:"type"`
		Uri  string `json:"uri"`
		Name string `json:"name"`
	}
}

type added_by struct {
	Followers followers `json:"followers"`
	Href      string    `json:"href"`
	Id        string    `json:"id"`
	Type      string    `json:"type"`
	Uri       string    `json:"uri"`
}

type followers struct {
	Href  string `json:"href"`
	Total int32  `json:"total"`
}

type track struct {
	Album   album    `json:"album"`
	Artists []artist `json:"artists"`
}

type item struct {
	AddedAt string   `json:"added_at"`
	AddedBy added_by `json:"added_by"`
	IsLocal bool     `json:"is_local"`
	Track   track    `json:"track"`
}

type tracks struct {
	Tracks struct {
		Total int32  `json:"total"`
		Items []item `json:"items"`
	} `json:"tracks"`
}
