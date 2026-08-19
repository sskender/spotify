package spotify

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type SpotifyProfileImage struct {
	Height int    `json:"height"`
	Url    string `json:"url"`
	Width  int    `json:"width"`
}

type SpotifyProfile struct {
	DisplayName string                `json:"display_name"`
	Id          string                `json:"id"`
	Images      []SpotifyProfileImage `json:"images"`
}

type SpotifyTrackArtist struct {
	Name string `json:"name"`
}

type SpotifyTrackItem struct {
	Name         string               `json:"name"`
	Artists      []SpotifyTrackArtist `json:"artists"`
	ExternalURLs map[string]string    `json:"external_urls"`
}

type SpotifyTrackData struct {
	IsPlaying bool             `json:"is_playing"`
	Item      SpotifyTrackItem `json:"item"`
}

type SpotifyClient struct {
	client *http.Client
}

func NewSpotifyClient(client *http.Client) *SpotifyClient {
	return &SpotifyClient{client: client}
}

func (s *SpotifyClient) GetProfile() (*SpotifyProfile, error) {
	resp, err := s.client.Get("https://api.spotify.com/v1/me")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNoContent {
		return nil, nil
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	decoder := json.NewDecoder(resp.Body)

	var profile SpotifyProfile
	if err := decoder.Decode(&profile); err != nil {
		return nil, err
	}

	return &profile, nil
}

func (s *SpotifyClient) GetCurrentlyPlaying() (*SpotifyTrackData, error) {
	resp, err := s.client.Get("https://api.spotify.com/v1/me/player/currently-playing")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNoContent {
		return nil, nil
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	decoder := json.NewDecoder(resp.Body)

	var trackData SpotifyTrackData
	if err := decoder.Decode(&trackData); err != nil {
		return nil, err
	}

	return &trackData, nil
}
