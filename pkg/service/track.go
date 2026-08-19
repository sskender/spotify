package service

import (
	"encoding/json"
	"fmt"
)

type NowPlayingTrack struct {
	IsPlaying bool   `json:"is_playing"`
	Title     string `json:"title,omitempty"`
	Author    string `json:"author,omitempty"`
	Link      string `json:"link,omitempty"`
}

func (t *NowPlayingTrack) GenerateJson() string {
	res, _ := json.Marshal(t)
	return string(res)
}

func (t *NowPlayingTrack) GenerateSvg() string {
	if !t.IsPlaying {
		return SVG_TEMPLATE_NOT_PLAYING
	}
	return fmt.Sprintf(SVG_TEMPLATE, t.Title, t.Author)
}

const SVG_TEMPLATE string = `
<svg width="460" height="88" viewBox="0 0 460 88" xmlns="http://www.w3.org/2000/svg" role="img">
  <title>Spotify Now Playing</title>

  <defs>
    <linearGradient id="bg" x1="0" y1="0" x2="1" y2="1">
      <stop offset="0%%" stop-color="#121212"/>
        <stop offset="100%%" stop-color="#1DB954" stop-opacity="1"/>
    </linearGradient>

    <clipPath id="textClip">
      <rect x="80" y="0" width="310" height="88"/>
    </clipPath>
  </defs>

  <rect width="460" height="88" rx="12" fill="url(#bg)" />

  <g transform="translate(20,20) scale(0.21)">
    <path
      d="M177.707 98.987c-35.992-21.375-95.36-23.34-129.719-12.912-5.519 1.674-11.353-1.44-13.024-6.958-1.672-5.521 1.439-11.352 6.96-13.029 39.443-11.972 105.008-9.66 146.443 14.936 4.964 2.947 6.59 9.356 3.649 14.31-2.944 4.963-9.359 6.6-14.31 3.653m-1.178 31.658c-2.525 4.098-7.883 5.383-11.975 2.867-30.005-18.444-75.762-23.788-111.262-13.012-4.603 1.39-9.466-1.204-10.864-5.8a8.717 8.717 0 015.805-10.856c40.553-12.307 90.968-6.347 125.432 14.833 4.092 2.52 5.38 7.88 2.864 11.968m-13.663 30.404a6.954 6.954 0 01-9.569 2.316c-26.22-16.025-59.223-19.644-98.09-10.766a6.955 6.955 0 01-8.331-5.232 6.95 6.95 0 015.233-8.334c42.533-9.722 79.017-5.538 108.448 12.446a6.96 6.96 0 012.31 9.57M111.656 0C49.992 0 0 49.99 0 111.656c0 61.672 49.992 111.66 111.657 111.66 61.668 0 111.659-49.988 111.659-111.66C223.316 49.991 173.326 0 111.657 0"
      fill="#1ED660"
    />
  </g>

  <g clip-path="url(#textClip)">
    <text x="80" y="44"
          fill="#FFFFFF"
          font-family="-apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Ubuntu, sans-serif"
          font-size="14"
          font-weight="600">
		  %s
    </text>
    <text x="80" y="62"
          fill="#B3B3B3"
          font-family="-apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Ubuntu, sans-serif"
          font-size="12">
      %s
    </text>
  </g>

    <g transform="translate(420, 26)">
      <rect x="0" y="10" width="3" height="15" fill="#1DB954" rx="1">
          <animate attributeName="height" values="10;20;10" dur="1s" repeatCount="indefinite" />
          <animate attributeName="y" values="15;5;15" dur="1s" repeatCount="indefinite" />
      </rect>
      <rect x="6" y="5" width="3" height="25" fill="#1DB954" rx="1">
          <animate attributeName="height" values="20;30;20" dur="1.2s" repeatCount="indefinite" />
          <animate attributeName="y" values="10;0;10" dur="1.2s" repeatCount="indefinite" />
      </rect>
      <rect x="12" y="15" width="3" height="10" fill="#1DB954" rx="1">
          <animate attributeName="height" values="10;20;10" dur="0.8s" repeatCount="indefinite" />
          <animate attributeName="y" values="15;5;15" dur="0.8s" repeatCount="indefinite" />
      </rect>
  </g>
</svg>
`

const SVG_TEMPLATE_NOT_PLAYING string = `
<svg width="460" height="88" viewBox="0 0 460 88" xmlns="http://www.w3.org/2000/svg" role="img">
  <title>Spotify Now Playing</title>

  <defs>
    <linearGradient id="bg" x1="0" y1="0" x2="1" y2="1">
      <stop offset="0%" stop-color="#121212"/>
      <stop offset="100%" stop-color="#1DB954" stop-opacity="0.35"/>
    </linearGradient>

    <clipPath id="textClip">
      <rect x="80" y="0" width="340" height="88"/>
    </clipPath>
  </defs>

  <rect width="460" height="88" rx="12" fill="url(#bg)" />

  <g transform="translate(20,20) scale(0.21)" opacity="0.85">
    <path
      d="M177.707 98.987c-35.992-21.375-95.36-23.34-129.719-12.912-5.519 1.674-11.353-1.44-13.024-6.958-1.672-5.521 1.439-11.352 6.96-13.029 39.443-11.972 105.008-9.66 146.443 14.936 4.964 2.947 6.59 9.356 3.649 14.31-2.944 4.963-9.359 6.6-14.31 3.653m-1.178 31.658c-2.525 4.098-7.883 5.383-11.975 2.867-30.005-18.444-75.762-23.788-111.262-13.012-4.603 1.39-9.466-1.204-10.864-5.8a8.717 8.717 0 015.805-10.856c40.553-12.307 90.968-6.347 125.432 14.833 4.092 2.52 5.38 7.88 2.864 11.968m-13.663 30.404a6.954 6.954 0 01-9.569 2.316c-26.22-16.025-59.223-19.644-98.09-10.766a6.955 6.955 0 01-8.331-5.232 6.95 6.95 0 015.233-8.334c42.533-9.722 79.017-5.538 108.448 12.446a6.96 6.96 0 012.31 9.57M111.656 0C49.992 0 0 49.99 0 111.656c0 61.672 49.992 111.66 111.657 111.66 61.668 0 111.659-49.988 111.659-111.66C223.316 49.991 173.326 0 111.657 0"
      fill="#1ED660"
    />
  </g>

  <g clip-path="url(#textClip)">
    <text x="80"
          y="48"
          fill="#E6E6E6"
          font-family="-apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Ubuntu, sans-serif"
          font-size="14"
          font-weight="500"
          dominant-baseline="middle">
      Not playing anything at the moment
    </text>
  </g>
</svg>
`
