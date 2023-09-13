package object

import "time"

type Anime struct {
	Loop      bool
	Type      int
	InitValue float64
	StartTime float64
	Duration  float64
	Curve     func(t float64) float32
	Time      time.Time
	End       func()
}

const (
	ANIME_ALPHA = iota
	ANIME_ROTATE
	ANIME_XPOS
	ANIME_YPOS
)
