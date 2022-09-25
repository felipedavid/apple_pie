package data

import (
	"time"
)

type Game struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	Title     string    `json:"title"`
	Year      int32     `json:"year"`
	Playtime  playtime  `json:"playtime,omitempty"`
	Genre     []string  `json:"genre,omitempty"`
	Version   int32     `json:"version"`
}
