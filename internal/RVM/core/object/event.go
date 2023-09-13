package object

import (
	"internal/RVM/core/system/game/event"
)

type Key struct {
	Down, Up func(e *event.EVENT_Key)
}

func (k *Key) screenObj() {}

type Button struct {
	MainImageName  string
	HoverImageName string
	T              Transform
	Anime          []*Anime
	Down           func(e *event.EVENT_MouseButton)
	Up             func(e *event.EVENT_MouseButton)
	Hover          func(e *event.EVENT_MouseMotion)
	UnHover        func(e *event.EVENT_MouseMotion)
}

func (b *Button) screenObj() {}

type Bar struct {
	FrameImageName       string
	CursorImageName      string
	CursorHoverImageName string
	GaugeImageName       string

	FrameImageT  Transform
	CursorSize   Vector2
	StartPadding float32
	EndPadding   float32
	SidePadding  float32
	IsVertical   bool

	MaxValue  float32
	MinValue  float32
	InitValue float32

	Down   func(e *event.EVENT_MouseButton, value float32)
	Up     func(e *event.EVENT_MouseButton, value float32)
	Scroll func(e *event.EVENT_MouseMotion, value float32)
}

func (b *Bar) screenObj() {}
