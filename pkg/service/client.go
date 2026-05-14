package service

import "github.com/sskender/spotify/pkg/spotify"

type Client interface {
	GetCurrentlyPlaying() (*spotify.SpotifyTrackData, error)
}
