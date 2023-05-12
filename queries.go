package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"os"
)

func search(query string) []SearchResult {
	var searchResults Search

	url := "https://doubledouble.top/search?service=odesli&q=" + url.QueryEscape(query)

	resp, err := http.Get(url)
	if err != nil {
		return searchResults.SearchResult
	}

	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(&searchResults)
	return searchResults.SearchResult
}

func getDownloadOptions(id string) (options [][2]string) {
	url := "https://doubledouble.top/resolve?url=" + url.QueryEscape(id)

	resp, err := http.Get(url)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	var download_results DownloadOptions
	json.NewDecoder(resp.Body).Decode(&download_results)

	if download_results.Spotify != "" {
		options = append(options, [2]string{
			"Spotify",
			download_results.Spotify,
		})
	}

	if download_results.Tidal != "" {
		options = append(options, [2]string{
			"Tidal",
			download_results.Tidal,
		})
	}

	if download_results.Deezer != "" {
		options = append(options, [2]string{
			"Deezer",
			download_results.Deezer,
		})
	}

	if download_results.Soundcloud != "" {
		options = append(options, [2]string{
			"Soundcloud",
			download_results.Soundcloud,
		})
	}

	if download_results.Napster != "" {
		options = append(options, [2]string{
			"Napster",
			download_results.Napster,
		})
	}

	if download_results.Youtube != "" {
		options = append(options, [2]string{
			"Youtube Music",
			download_results.Youtube,
		})
	}

	return
}

func requestDownload(d_url string) (id string) {
	url := "https://doubledouble.top/dl?url=" + url.QueryEscape(d_url)

	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	var download_request DownloadRequest
	json.NewDecoder(resp.Body).Decode(&download_request)

	return download_request.ID
}

func getDownloadStatus(id string) (status DownloadStatus) {
	url := "https://doubledouble.top/dl/" + url.QueryEscape(id)

	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(&status)

	return
}

func startDownload(id string) (err error) {
	file, err := os.Create(id + ".zip")
	if err != nil {
		return err
	}
	defer file.Close()

	url := "https://doubledouble.top/temp/" + url.QueryEscape(id) + ".zip"
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
