package spoty

import "fmt"

// Result represent the result of a status call
type Result struct {
	ClientVersion string `json:"client_version"`
	Version       int    `json:"version"`

	Running bool `json:"running"`
	Playing bool `json:"playing"`
	Shuffle bool `json:"shuffle"`
	Repeat  bool `json:"repeat"`

	Error struct {
		Type    string `json:"type"`
		Message string `json:"message"`
	}

	Track Track `json:"track"`
}

type Track struct {
	TrackResource struct {
		Name     string `json:"name"`
		URI      string `json:"uri"`
		Location struct {
			OG string `json:"og"`
		} `json:"location"`
	} `json:"track_resource"`
	ArtistResource struct {
		Name     string `json:"name"`
		URI      string `json:"uri"`
		Location struct {
			OG string `json:"og"`
		} `json:"location"`
	} `json:"artist_resource"`
	AlbumResource struct {
		Name     string `json:"name"`
		URI      string `json:"uri"`
		Location struct {
			OG string `json:"og"`
		} `json:"location"`
	} `json:"album_resource"`
}

func (r *Result) CurrentSongURI() string {
	if r != nil {
		return r.Track.TrackResource.URI
	}

	return ""
}

func (r *Result) CurrentSongTitle() string {
	if r != nil {
		return fmt.Sprintf("%s - %s by %s", r.Track.TrackResource.Name, r.Track.AlbumResource.Name, r.Track.ArtistResource.Name)
	}

	return ""
}
