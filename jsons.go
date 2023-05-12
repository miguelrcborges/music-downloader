package main

type Search struct {
	SearchResult []SearchResult `json:"results"`
}

type SearchResult struct {
	Type   string `json:"type"`
	Link   string `json:"links"`
	Name   string `json:"name"`
	Artist string `json:"artist"`
	// Album  string `json:"album"`
}

type DownloadOptions struct {
	Spotify    string `json:"spotify"`
	Tidal      string `json:"tidal"`
	Deezer     string `json:"deezer"`
	Soundcloud string `json:"soundcloud"`
	Napster    string `json:"napster"`
	Youtube    string `json:"youtube"`
}

type DownloadRequest struct {
	ID string `json:"id"`
}

type DownloadStatus struct {
	Status  string `json:"friendlyStatus"`
	Percent int    `json:"percent"`
}
