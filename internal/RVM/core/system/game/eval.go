package game

import (
	"internal/RVM/core/globaltype"
	obj "internal/RVM/core/object"
	event "internal/RVM/core/system/game/event"
	"log"
	"unicode/utf8"
)

func (g *Game) evalShow(s *obj.Show, name string, bps int) {

	if s.T.Size.X != 0 && s.T.Size.Y != 0 {
		s.T = g.echoTransform(s.T, s.T.Size.X, s.T.Size.Y)
	} else {
		s.T = g.echoTransform(s.T, float32(g.Graphic.Image_Manager.GetImageWidth(s.Name)), float32(g.Graphic.Image_Manager.GetImageHeight(s.Name)))
	}

	g.Graphic.AddScreenTextureRenderBuffer(
		bps,
		g.Graphic.Image_Manager.GetImageTexture(s.Name),
		s.T,
	)

	if len(s.Anime) != 0 {
		for _, anime := range s.Anime {
			switch anime.Type {
			case obj.ANIME_ALPHA:
				g.Graphic.Image_Manager.SetImageAlpha(s.Name, uint8(anime.InitValue))
				g.Graphic.AddAnimation(name, anime, bps, g.Graphic.GetCurrentTopScreenIndexByBps(bps))
			case obj.ANIME_ROTATE:
				g.Graphic.SetRotateByBps(bps, g.Graphic.GetCurrentTopScreenIndexByBps(bps), float32(anime.InitValue))
				g.Graphic.AddAnimation(name, anime, bps, g.Graphic.GetCurrentTopScreenIndexByBps(bps))
			case obj.ANIME_XPOS:
				g.Graphic.SetCurrentTextureXPosition(bps, g.Graphic.GetCurrentTopScreenIndexByBps(bps), float32(anime.InitValue))
				g.Graphic.AddAnimation(name, anime, bps, g.Graphic.GetCurrentTopScreenIndexByBps(bps))
			case obj.ANIME_YPOS:
				g.Graphic.SetCurrentTextureYPosition(bps, g.Graphic.GetCurrentTopScreenIndexByBps(bps), float32(anime.InitValue))
				g.Graphic.AddAnimation(name, anime, bps, g.Graphic.GetCurrentTopScreenIndexByBps(bps))
			}
		}
	}
}

func (g *Game) evalPlayVideo(pv *obj.PlayVideo, name string, bps int) {
	g.Graphic.AddScreenTextureRenderBuffer(
		bps,
		g.Graphic.GetVideoTexture(pv.Name),
		pv.T,
	)
	g.Graphic.VideoStart(name, pv.Name, pv.Loop)

	if len(pv.Anime) != 0 {
		for _, anime := range pv.Anime {
			switch anime.Type {
			case obj.ANIME_ALPHA:
				g.Graphic.SetVideoAlphaByName(pv.Name, uint8(anime.InitValue))
				g.Graphic.AddAnimation(name, anime, bps, g.Graphic.GetCurrentTopScreenIndexByBps(bps))
			case obj.ANIME_ROTATE:
				g.Graphic.SetRotateByBps(bps, g.Graphic.GetCurrentTopScreenIndexByBps(bps), float32(anime.InitValue))
				g.Graphic.AddAnimation(name, anime, bps, g.Graphic.GetCurrentTopScreenIndexByBps(bps))
			case obj.ANIME_XPOS:
				g.Graphic.SetCurrentTextureXPosition(bps, g.Graphic.GetCurrentTopScreenIndexByBps(bps), float32(anime.InitValue))
				g.Graphic.AddAnimation(name, anime, bps, g.Graphic.GetCurrentTopScreenIndexByBps(bps))
			case obj.ANIME_YPOS:
				g.Graphic.SetCurrentTextureYPosition(bps, g.Graphic.GetCurrentTopScreenIndexByBps(bps), float32(anime.InitValue))
				g.Graphic.AddAnimation(name, anime, bps, g.Graphic.GetCurrentTopScreenIndexByBps(bps))
			}
		}
	}
}

func (g *Game) evalText(t *obj.Text, name string, bps int) {
	if t.Text == "" {
		return
	}

	if t.TypingFX {
		var data []struct {
			Texture   *globaltype.SDL_Texture
			Transform obj.Transform
		}
		for index, runeValue := range t.Text {
			texture, width, height := g.Graphic.GetTextTexture(t.Text[0:index]+string(runeValue), t.FontName, t.Color)
			g.Graphic.RegisterTextMemPool(name, texture)
			t.T = g.echoTransform(t.T, float32(width), float32(height))
		}
		g.Graphic.AddScreenTextureRenderBuffer(
			bps,
			data[0].Texture,
			data[0].Transform,
		)
		g.Graphic.RegisterTypingFX(data, name, float64(utf8.RuneCountInString(t.Text))/g.TextSpeed, bps, g.Graphic.GetCurrentTopScreenIndexByBps(bps))
		return
	}

	texture, width, height := g.Graphic.GetTextTexture(t.Text, t.FontName, t.Color)
	g.Graphic.RegisterTextMemPool(name, texture)
	t.T = g.echoTransform(t.T, float32(width), float32(height))

	g.Graphic.AddScreenTextureRenderBuffer(
		bps,
		texture,
		t.T,
	)
}

func (g *Game) evalTextPointer(t *obj.TextPointer, name string, bps int) {
	if *t.Text == "" {
		return
	}

	if t.TypingFX {
		var data []struct {
			Texture   *globaltype.SDL_Texture
			Transform obj.Transform
		}
		for index, runeValue := range *t.Text {
			var transform obj.Transform
			texture, width, height := g.Graphic.GetTextTexture((*t.Text)[0:index]+string(runeValue), t.FontName, t.Color)
			g.Graphic.RegisterTextMemPool(name, texture)
			transform = g.echoTransform(t.T, float32(width), float32(height))
			data = append(data, struct {
				Texture   *globaltype.SDL_Texture
				Transform obj.Transform
			}{
				Texture:   texture,
				Transform: transform,
			})
		}
		g.Graphic.AddScreenTextureRenderBuffer(
			bps,
			data[0].Texture,
			data[0].Transform,
		)
		g.Graphic.RegisterTypingFX(data, name, float64(utf8.RuneCountInString(*t.Text))/g.TextSpeed, bps, g.Graphic.GetCurrentTopScreenIndexByBps(bps))
		return
	}

	texture, width, height := g.Graphic.GetTextTexture(*t.Text, t.FontName, t.Color)
	g.Graphic.RegisterTextMemPool(name, texture)
	t.T = g.echoTransform(t.T, float32(width), float32(height))

	g.Graphic.AddScreenTextureRenderBuffer(
		bps,
		texture,
		t.T,
	)
}

func (g *Game) evalButton(b *obj.Button, name string, bps int) {
	if b.T.Size.X != 0 && b.T.Size.Y != 0 {
		b.T = g.echoTransform(b.T, b.T.Size.X, b.T.Size.Y)
	} else {
		b.T = g.echoTransform(b.T, float32(g.Graphic.Image_Manager.GetImageWidth(b.MainImageName)), float32(g.Graphic.Image_Manager.GetImageHeight(b.MainImageName)))
	}

	g.Graphic.AddScreenTextureRenderBuffer(
		bps,
		g.Graphic.Image_Manager.GetImageTexture(b.MainImageName),
		b.T,
	)

	index := g.Graphic.GetCurrentTopScreenIndexByBps(bps)

	if len(b.Anime) != 0 {
		for _, anime := range b.Anime {
			switch anime.Type {
			case obj.ANIME_ALPHA:
				g.Graphic.Image_Manager.SetImageAlpha(b.MainImageName, uint8(anime.InitValue))
				g.Graphic.AddAnimation(name, anime, bps, g.Graphic.GetCurrentTopScreenIndexByBps(bps))
			case obj.ANIME_ROTATE:
				g.Graphic.SetRotateByBps(bps, g.Graphic.GetCurrentTopScreenIndexByBps(bps), float32(anime.InitValue))
				g.Graphic.AddAnimation(name, anime, bps, g.Graphic.GetCurrentTopScreenIndexByBps(bps))
			case obj.ANIME_XPOS:
				g.Graphic.SetCurrentTextureXPosition(bps, g.Graphic.GetCurrentTopScreenIndexByBps(bps), float32(anime.InitValue))
				g.Graphic.AddAnimation(name, anime, bps, g.Graphic.GetCurrentTopScreenIndexByBps(bps))
			case obj.ANIME_YPOS:
				g.Graphic.SetCurrentTextureYPosition(bps, g.Graphic.GetCurrentTopScreenIndexByBps(bps), float32(anime.InitValue))
				g.Graphic.AddAnimation(name, anime, bps, g.Graphic.GetCurrentTopScreenIndexByBps(bps))
			}
		}
	}

	var e event.ButtonEvent

	if b.Down != nil && b.Up != nil {
		e.Down = func(e *event.EVENT_MouseButton) bool {
			var xpos, ypos float32 = g.Graphic.GetCurrentTexturePosition(g.screenBps[name], index)
			var xsize, ysize float32 = g.Graphic.GetCurrentTextureSize(g.screenBps[name], index)
			if g.Graphic.GetFixedRealXSize(e.X) > xpos &&
				g.Graphic.GetFixedRealYSize(e.Y) > ypos &&
				g.Graphic.GetFixedRealXSize(e.X) < xpos+xsize &&
				g.Graphic.GetFixedRealYSize(e.Y) < ypos+ysize &&
				e.Button == event.RENG_BUTTON_LEFT {
				b.Down(e)
				return true
			}
			return false
		}

		e.Up = func(e *event.EVENT_MouseButton) {
			b.Up(e)
		}

		if b.Hover != nil && b.UnHover != nil {
			e.Hover = func(e *event.EVENT_MouseMotion) {
				var xpos, ypos float32 = g.Graphic.GetCurrentTexturePosition(g.screenBps[name], index)
				var xsize, ysize float32 = g.Graphic.GetCurrentTextureSize(g.screenBps[name], index)
				if g.Graphic.GetFixedRealXSize(e.X) >= xpos &&
					g.Graphic.GetFixedRealYSize(e.Y) >= ypos &&
					g.Graphic.GetFixedRealXSize(e.X) <= xpos+xsize &&
					g.Graphic.GetFixedRealYSize(e.Y) <= ypos+ysize {
					if b.HoverImageName != "" {
						g.Graphic.ChangeTextureByBpsUseImageManager(g.screenBps[name], index, b.HoverImageName)
					}
					b.Hover(e)
				} else {
					if b.HoverImageName != "" {
						g.Graphic.ChangeTextureByBpsUseImageManager(g.screenBps[name], index, b.MainImageName)
					}
					b.UnHover(e)
				}
			}
		}

		g.Event.AddButtonEvent(name, e)
	}
}

func (g *Game) evalBar(b *obj.Bar, name string, bps int) {
	if b.MaxValue <= b.MinValue {
		log.Fatalf("Error : MaxValue <= MinValue")
		return
	}

	percent := float64(b.InitValue-b.MinValue) / float64(b.MaxValue-b.MinValue)

	if percent < 0 {
		log.Fatalf("Error : InitValue is smaller than MinValue")
		return
	}

	if b.FrameImageT.Size.X != 0 && b.FrameImageT.Size.Y != 0 {
		b.FrameImageT = g.echoTransform(b.FrameImageT, b.FrameImageT.Size.X, b.FrameImageT.Size.Y)
	} else {
		b.FrameImageT = g.echoTransform(b.FrameImageT, float32(g.Graphic.Image_Manager.GetImageWidth(b.FrameImageName)), float32(g.Graphic.Image_Manager.GetImageHeight(b.FrameImageName)))
	}

	g.Graphic.AddScreenTextureRenderBuffer(
		bps,
		g.Graphic.Image_Manager.GetImageTexture(b.FrameImageName),
		b.FrameImageT,
	)

	frameIndex := g.Graphic.GetCurrentTopScreenIndexByBps(bps)

	if b.IsVertical {
		b.FrameImageT.Pos.Y += b.StartPadding
		b.FrameImageT.Flip.X = float32(g.Graphic.Image_Manager.GetImageWidth(b.GaugeImageName))
		b.FrameImageT.Flip.Y = b.FrameImageT.Size.Y * float32(percent)
		b.FrameImageT.Size.Y = b.FrameImageT.Size.Y * float32(percent)
		b.FrameImageT.Pos.X += b.SidePadding
		b.FrameImageT.Size.X -= b.SidePadding * 2
	} else {
		b.FrameImageT.Pos.X += b.StartPadding
		b.FrameImageT.Flip.X = b.FrameImageT.Size.X * float32(percent)
		b.FrameImageT.Flip.Y = float32(g.Graphic.Image_Manager.GetImageHeight(b.GaugeImageName))
		b.FrameImageT.Size.X = b.FrameImageT.Size.X * float32(percent)
		b.FrameImageT.Pos.Y += b.SidePadding
		b.FrameImageT.Size.Y -= b.SidePadding * 2
	}

	g.Graphic.AddScreenTextureRenderBuffer(
		bps,
		g.Graphic.Image_Manager.GetImageTexture(b.GaugeImageName),
		b.FrameImageT,
	)

	gaugeIndex := g.Graphic.GetCurrentTopScreenIndexByBps(bps)

	if b.IsVertical {
		b.FrameImageT.Pos.X = b.FrameImageT.Pos.X + ((b.FrameImageT.Flip.X) / 2) - float32(g.Graphic.Image_Manager.GetImageWidth(b.CursorImageName))/2
		b.FrameImageT.Pos.Y = b.FrameImageT.Pos.Y + b.FrameImageT.Flip.Y - float32(g.Graphic.Image_Manager.GetImageHeight(b.CursorImageName))/2
	} else {
		b.FrameImageT.Pos.Y = b.FrameImageT.Pos.Y + ((b.FrameImageT.Size.Y) / 2) - float32(g.Graphic.Image_Manager.GetImageHeight(b.CursorImageName))/2
		b.FrameImageT.Pos.X = b.FrameImageT.Pos.X + b.FrameImageT.Flip.X - float32(g.Graphic.Image_Manager.GetImageWidth(b.CursorImageName))/2
	}

	b.FrameImageT.Flip.X = 0
	b.FrameImageT.Flip.Y = 0

	b.FrameImageT.Size.X = float32(g.Graphic.Image_Manager.GetImageWidth(b.CursorImageName))
	b.FrameImageT.Size.Y = float32(g.Graphic.Image_Manager.GetImageHeight(b.CursorImageName))

	g.Graphic.AddScreenTextureRenderBuffer(
		bps,
		g.Graphic.Image_Manager.GetImageTexture(b.CursorImageName),
		b.FrameImageT,
	)

	cursorIndex := g.Graphic.GetCurrentTopScreenIndexByBps(bps)

	if b.CursorHoverImageName != "" {
		if b.Down != nil && b.Up != nil && b.Scroll != nil {
			g.Event.AddBarEvent(
				name,
				event.BarEvent{
					IsNowDown: false,
					Down: func(e *event.EVENT_MouseButton) bool {
						bps := g.screenBps[name]
						var xpos, ypos float32 = g.Graphic.GetCurrentTexturePosition(g.screenBps[name], cursorIndex)
						var xsize, ysize float32 = g.Graphic.GetCurrentTextureSize(g.screenBps[name], cursorIndex)
						if g.Graphic.GetFixedRealXSize(e.X) > xpos+xsize/2-b.CursorSize.X/2 &&
							g.Graphic.GetFixedRealYSize(e.Y) > ypos+ysize/2-b.CursorSize.Y/2 &&
							g.Graphic.GetFixedRealXSize(e.X) < xpos+xsize/2+b.CursorSize.X/2 &&
							g.Graphic.GetFixedRealYSize(e.Y) < ypos+ysize/2+b.CursorSize.Y/2 &&
							e.Button == event.RENG_BUTTON_LEFT {
							g.Graphic.ChangeTextureByBpsUseImageManager(g.screenBps[name], cursorIndex, b.CursorHoverImageName)
							if b.IsVertical {
								if g.Graphic.GetFixedRealYSize(e.Y) >= g.Graphic.GetCurrentTextureYPosition(bps, frameIndex)+g.Graphic.GetCurrentTextureYSize(bps, frameIndex)-b.EndPadding {
									g.Graphic.SetCurrentTextureYFlip(bps, gaugeIndex, g.Graphic.GetCurrentTextureYSize(bps, frameIndex)-(b.StartPadding+b.EndPadding))
									g.Graphic.SetCurrentTextureYPosition(bps, cursorIndex, g.Graphic.GetCurrentTextureYPosition(bps, gaugeIndex)+g.Graphic.GetCurrentTextureYFlip(bps, gaugeIndex)-float32(g.Graphic.Image_Manager.GetImageHeight(b.CursorImageName))/2)
									b.Down(e, b.MaxValue)
								} else if g.Graphic.GetFixedRealYSize(e.Y) <= g.Graphic.GetCurrentTextureYPosition(bps, gaugeIndex) {
									g.Graphic.SetCurrentTextureYFlip(bps, gaugeIndex, 0)
									g.Graphic.SetCurrentTextureYPosition(bps, cursorIndex, g.Graphic.GetCurrentTextureYPosition(bps, gaugeIndex)-float32(g.Graphic.Image_Manager.GetImageHeight(b.CursorImageName))/2)
									b.Down(e, b.MinValue)
								} else {
									g.Graphic.SetCurrentTextureYFlip(bps, gaugeIndex, g.Graphic.GetFixedRealYSize(e.Y)-g.Graphic.GetCurrentTextureYPosition(bps, gaugeIndex))
									g.Graphic.SetCurrentTextureYPosition(bps, cursorIndex, g.Graphic.GetCurrentTextureYPosition(bps, gaugeIndex)+g.Graphic.GetCurrentTextureYFlip(bps, gaugeIndex)-float32(g.Graphic.Image_Manager.GetImageWidth(b.CursorImageName))/2)
									b.Down(e, b.MinValue+b.MaxValue-b.MinValue*((g.Graphic.GetFixedRealYSize(e.Y)-g.Graphic.GetCurrentTextureYPosition(bps, gaugeIndex))/(g.Graphic.GetCurrentTextureYSize(bps, frameIndex)-(b.StartPadding+b.EndPadding))))
								}
							} else {
								if g.Graphic.GetFixedRealXSize(e.X) >= g.Graphic.GetCurrentTextureXPosition(bps, frameIndex)+g.Graphic.GetCurrentTextureXSize(bps, frameIndex)-b.EndPadding {
									g.Graphic.SetCurrentTextureXFlip(bps, gaugeIndex, g.Graphic.GetCurrentTextureXSize(bps, frameIndex)-(b.StartPadding+b.EndPadding))
									g.Graphic.SetCurrentTextureXSize(bps, gaugeIndex, g.Graphic.GetCurrentTextureXSize(bps, frameIndex)-(b.StartPadding+b.EndPadding))
									g.Graphic.SetCurrentTextureXPosition(bps, cursorIndex, g.Graphic.GetCurrentTextureXPosition(bps, gaugeIndex)+g.Graphic.GetCurrentTextureXFlip(bps, gaugeIndex)-float32(g.Graphic.Image_Manager.GetImageWidth(b.CursorImageName))/2)
									b.Down(e, b.MaxValue)
								} else if g.Graphic.GetFixedRealXSize(e.X) <= g.Graphic.GetCurrentTextureXPosition(bps, gaugeIndex) {
									g.Graphic.SetCurrentTextureXFlip(bps, gaugeIndex, 0)
									g.Graphic.SetCurrentTextureXSize(bps, gaugeIndex, 0)
									g.Graphic.SetCurrentTextureXPosition(bps, cursorIndex, g.Graphic.GetCurrentTextureXPosition(bps, gaugeIndex)-float32(g.Graphic.Image_Manager.GetImageWidth(b.CursorImageName)/2))
									b.Down(e, b.MinValue)
								} else {
									g.Graphic.SetCurrentTextureXFlip(bps, gaugeIndex, g.Graphic.GetFixedRealXSize(e.X)-g.Graphic.GetCurrentTextureXPosition(bps, gaugeIndex))
									g.Graphic.SetCurrentTextureXSize(bps, gaugeIndex, g.Graphic.GetFixedRealXSize(e.X)-g.Graphic.GetCurrentTextureXPosition(bps, gaugeIndex))
									g.Graphic.SetCurrentTextureXPosition(bps, cursorIndex, g.Graphic.GetCurrentTextureXPosition(bps, gaugeIndex)+g.Graphic.GetCurrentTextureXFlip(bps, gaugeIndex)-float32(g.Graphic.Image_Manager.GetImageWidth(b.CursorImageName))/2)
									b.Down(e, b.MinValue+(b.MaxValue-b.MinValue)*((g.Graphic.GetFixedRealXSize(e.X)-g.Graphic.GetCurrentTextureXPosition(bps, gaugeIndex))/(g.Graphic.GetCurrentTextureXSize(bps, frameIndex)-(b.StartPadding+b.EndPadding))))
								}
							}
							return true
						}
						return false
					},
					Up: func(e *event.EVENT_MouseButton) {
						bps := g.screenBps[name]
						g.Graphic.ChangeTextureByBpsUseImageManager(g.screenBps[name], cursorIndex, b.CursorImageName)
						if b.IsVertical {
							if g.Graphic.GetFixedRealYSize(e.Y) >= g.Graphic.GetCurrentTextureYPosition(bps, frameIndex)+g.Graphic.GetCurrentTextureYSize(bps, frameIndex)-b.EndPadding {
								g.Graphic.SetCurrentTextureYFlip(bps, gaugeIndex, g.Graphic.GetCurrentTextureYSize(bps, frameIndex)-(b.StartPadding+b.EndPadding))
								g.Graphic.SetCurrentTextureYPosition(bps, cursorIndex, g.Graphic.GetCurrentTextureYPosition(bps, gaugeIndex)+g.Graphic.GetCurrentTextureYFlip(bps, gaugeIndex)-float32(g.Graphic.Image_Manager.GetImageHeight(b.CursorImageName))/2)
								b.Up(e, b.MaxValue)
							} else if g.Graphic.GetFixedRealYSize(e.Y) <= g.Graphic.GetCurrentTextureYPosition(bps, gaugeIndex) {
								g.Graphic.SetCurrentTextureYFlip(bps, gaugeIndex, 0)
								g.Graphic.SetCurrentTextureYPosition(bps, cursorIndex, g.Graphic.GetCurrentTextureYPosition(bps, gaugeIndex)-float32(g.Graphic.Image_Manager.GetImageHeight(b.CursorImageName))/2)
								b.Up(e, b.MinValue)
							} else {
								g.Graphic.SetCurrentTextureYFlip(bps, gaugeIndex, g.Graphic.GetFixedRealYSize(e.Y)-g.Graphic.GetCurrentTextureYPosition(bps, gaugeIndex))
								g.Graphic.SetCurrentTextureYPosition(bps, cursorIndex, g.Graphic.GetCurrentTextureYPosition(bps, gaugeIndex)+g.Graphic.GetCurrentTextureYFlip(bps, gaugeIndex)-float32(g.Graphic.Image_Manager.GetImageHeight(b.CursorImageName))/2)
								b.Up(e, b.MinValue+((b.MaxValue-b.MinValue)*((g.Graphic.GetFixedRealYSize(e.Y)-g.Graphic.GetCurrentTextureYPosition(bps, gaugeIndex))/g.Graphic.GetCurrentTextureYSize(bps, frameIndex)-(b.StartPadding+b.EndPadding))))
							}
						} else {
							if g.Graphic.GetFixedRealXSize(e.X) >= g.Graphic.GetCurrentTextureXPosition(bps, frameIndex)+g.Graphic.GetCurrentTextureXSize(bps, frameIndex)-b.EndPadding {
								g.Graphic.SetCurrentTextureXFlip(bps, gaugeIndex, g.Graphic.GetCurrentTextureXSize(bps, frameIndex)-(b.StartPadding+b.EndPadding))
								g.Graphic.SetCurrentTextureXSize(bps, gaugeIndex, g.Graphic.GetCurrentTextureXSize(bps, frameIndex)-(b.StartPadding+b.EndPadding))
								g.Graphic.SetCurrentTextureXPosition(bps, cursorIndex, g.Graphic.GetCurrentTextureXPosition(bps, gaugeIndex)+g.Graphic.GetCurrentTextureXFlip(bps, gaugeIndex)-float32(g.Graphic.Image_Manager.GetImageWidth(b.CursorImageName))/2)
								b.Up(e, b.MaxValue)
							} else if g.Graphic.GetFixedRealXSize(e.X) <= g.Graphic.GetCurrentTextureXPosition(bps, gaugeIndex) {
								g.Graphic.SetCurrentTextureXFlip(bps, gaugeIndex, 0)
								g.Graphic.SetCurrentTextureXSize(bps, gaugeIndex, 0)
								g.Graphic.SetCurrentTextureXPosition(bps, cursorIndex, g.Graphic.GetCurrentTextureXPosition(bps, gaugeIndex)-float32(g.Graphic.Image_Manager.GetImageWidth(b.CursorImageName))/2)
								b.Up(e, b.MinValue)
							} else {
								g.Graphic.SetCurrentTextureXFlip(bps, gaugeIndex, g.Graphic.GetFixedRealXSize(e.X)-g.Graphic.GetCurrentTextureXPosition(bps, gaugeIndex))
								g.Graphic.SetCurrentTextureXSize(bps, gaugeIndex, g.Graphic.GetFixedRealXSize(e.X)-g.Graphic.GetCurrentTextureXPosition(bps, gaugeIndex))
								g.Graphic.SetCurrentTextureXPosition(bps, cursorIndex, g.Graphic.GetCurrentTextureXPosition(bps, gaugeIndex)+g.Graphic.GetCurrentTextureXFlip(bps, gaugeIndex)-float32(g.Graphic.Image_Manager.GetImageWidth(b.CursorImageName))/2)
								b.Up(e, b.MinValue+((b.MaxValue-b.MinValue)*((g.Graphic.GetFixedRealXSize(e.X)-g.Graphic.GetCurrentTextureXPosition(bps, gaugeIndex))/(g.Graphic.GetCurrentTextureXSize(bps, frameIndex)-(b.StartPadding+b.EndPadding)))))
							}
						}
					},
					Scroll: func(e *event.EVENT_MouseMotion) {
						bps := g.screenBps[name]
						if b.IsVertical {
							if g.Graphic.GetFixedRealYSize(e.Y) >= g.Graphic.GetCurrentTextureYPosition(bps, frameIndex)+g.Graphic.GetCurrentTextureYSize(bps, frameIndex)-b.EndPadding {
								g.Graphic.SetCurrentTextureYFlip(bps, gaugeIndex, g.Graphic.GetCurrentTextureYSize(bps, frameIndex)-(b.StartPadding+b.EndPadding))
								g.Graphic.SetCurrentTextureYPosition(bps, cursorIndex, g.Graphic.GetCurrentTextureYPosition(bps, gaugeIndex)+g.Graphic.GetCurrentTextureYFlip(bps, gaugeIndex)-float32(g.Graphic.Image_Manager.GetImageHeight(b.CursorImageName))/2)
								b.Scroll(e, b.MaxValue)
							} else if g.Graphic.GetFixedRealYSize(e.Y) <= g.Graphic.GetCurrentTextureYPosition(bps, gaugeIndex) {
								g.Graphic.SetCurrentTextureYFlip(bps, gaugeIndex, 0)
								g.Graphic.SetCurrentTextureYPosition(bps, cursorIndex, g.Graphic.GetCurrentTextureYPosition(bps, gaugeIndex)-float32(g.Graphic.Image_Manager.GetImageHeight(b.CursorImageName))/2)
								b.Scroll(e, b.MinValue)
							} else {
								g.Graphic.SetCurrentTextureYFlip(bps, gaugeIndex, g.Graphic.GetFixedRealYSize(e.Y)-g.Graphic.GetCurrentTextureYPosition(bps, gaugeIndex))
								g.Graphic.SetCurrentTextureYPosition(bps, cursorIndex, g.Graphic.GetCurrentTextureYPosition(bps, gaugeIndex)+g.Graphic.GetCurrentTextureYFlip(bps, gaugeIndex)-float32(g.Graphic.Image_Manager.GetImageHeight(b.CursorImageName))/2)
								b.Scroll(e, b.MinValue+((b.MaxValue-b.MinValue)*((g.Graphic.GetFixedRealYSize(e.Y)-g.Graphic.GetCurrentTextureYPosition(bps, gaugeIndex))/(g.Graphic.GetCurrentTextureYSize(bps, frameIndex)-(b.StartPadding+b.EndPadding)))))
							}
						} else {
							if g.Graphic.GetFixedRealXSize(e.X) >= g.Graphic.GetCurrentTextureXPosition(bps, frameIndex)+g.Graphic.GetCurrentTextureXSize(bps, frameIndex)-b.EndPadding {
								g.Graphic.SetCurrentTextureXFlip(bps, gaugeIndex, g.Graphic.GetCurrentTextureXSize(bps, frameIndex)-(b.StartPadding+b.EndPadding))
								g.Graphic.SetCurrentTextureXSize(bps, gaugeIndex, g.Graphic.GetCurrentTextureXSize(bps, frameIndex)-(b.StartPadding+b.EndPadding))
								g.Graphic.SetCurrentTextureXPosition(bps, cursorIndex, g.Graphic.GetCurrentTextureXPosition(bps, gaugeIndex)+g.Graphic.GetCurrentTextureXFlip(bps, gaugeIndex)-float32(g.Graphic.Image_Manager.GetImageWidth(b.CursorImageName))/2)
								b.Scroll(e, b.MaxValue)
							} else if g.Graphic.GetFixedRealXSize(e.X) <= g.Graphic.GetCurrentTextureXPosition(bps, gaugeIndex) {
								g.Graphic.SetCurrentTextureXFlip(bps, gaugeIndex, 0)
								g.Graphic.SetCurrentTextureXSize(bps, gaugeIndex, 0)
								g.Graphic.SetCurrentTextureXPosition(bps, cursorIndex, g.Graphic.GetCurrentTextureXPosition(bps, gaugeIndex)-float32(g.Graphic.Image_Manager.GetImageWidth(b.CursorImageName))/2)
								b.Scroll(e, b.MinValue)
							} else {
								g.Graphic.SetCurrentTextureXFlip(bps, gaugeIndex, g.Graphic.GetFixedRealXSize(e.X)-g.Graphic.GetCurrentTextureXPosition(bps, gaugeIndex))
								g.Graphic.SetCurrentTextureXSize(bps, gaugeIndex, g.Graphic.GetFixedRealXSize(e.X)-g.Graphic.GetCurrentTextureXPosition(bps, gaugeIndex))
								g.Graphic.SetCurrentTextureXPosition(bps, cursorIndex, g.Graphic.GetCurrentTextureXPosition(bps, gaugeIndex)+g.Graphic.GetCurrentTextureXFlip(bps, gaugeIndex)-float32(g.Graphic.Image_Manager.GetImageWidth(b.CursorImageName))/2)
								b.Scroll(e, b.MinValue+((b.MaxValue-b.MinValue)*((g.Graphic.GetFixedRealXSize(e.X)-g.Graphic.GetCurrentTextureXPosition(bps, gaugeIndex))/(g.Graphic.GetCurrentTextureXSize(bps, frameIndex)-(b.StartPadding+b.EndPadding)))))
							}
						}
					},
				},
			)
		}
	}
}

func (g *Game) echoTransform(t obj.Transform, width, height float32) obj.Transform {
	if t.Type != nil {
		switch b := (t.Type).(type) {
		case *obj.Center:
			t.Pos = obj.Vector2{
				X: (float32(g.width) - width) / 2,
				Y: (float32(g.height) - height) / 2,
			}
		case *obj.XCenter:
			t.Pos = obj.Vector2{
				X: (float32(g.width) - width) / 2,
				Y: b.Ypos,
			}
		case *obj.YCenter:
			t.Pos = obj.Vector2{
				X: b.Xpos,
				Y: (float32(g.height) - height) / 2,
			}
		case *obj.AxisCenter:
			t.Pos = obj.Vector2{
				X: b.Axis.X - width/2,
				Y: b.Axis.Y - height/2,
			}
		}
	}
	t.Size.X = width
	t.Size.Y = height

	return t
}

/*-------------- LABEL OBJECT -------------*/

func (g *Game) evalHide(h *obj.Hide, name string, bps int) {
	if h.Anime != nil {
		if h.Anime.End != nil {
			temp := h.Anime.End
			h.Anime.End = func() {
				temp()
				g.Graphic.DeleteAnimationByTextureIndex(name, h.TextureIndex)
				g.Graphic.DeleteScreenTextureRenderBuffer(bps, h.TextureIndex)
			}
		} else {
			h.Anime.End = func() {
				g.Graphic.DeleteAnimationByTextureIndex(name, h.TextureIndex)
				g.Graphic.DeleteScreenTextureRenderBuffer(bps, h.TextureIndex)
			}
		}
		g.Graphic.AddAnimation(name, h.Anime, bps, h.TextureIndex)
	} else {
		g.Graphic.DeleteAnimationByTextureIndex(name, h.TextureIndex)
		g.Graphic.DeleteScreenTextureRenderBuffer(bps, h.TextureIndex)
	}
}

func (g *Game) evalSay(s *obj.Say) error {
	*g.nowName = s.Character.Name
	*g.nowText = s.Text

	g.Graphic.SayLock()
	g.InActiveScreen(g.SayScreenName)
	err := g.ActiveScreen(g.SayScreenName)
	g.Graphic.SayUnlock()
	if err != nil {
		return err
	}

	lock := make(chan int)
	g.Event.AddMouseClickEvent(g.SayScreenName, event.MouseClickEvent{
		Down: func(e *event.EVENT_MouseButton) {},
		Up: func(e *event.EVENT_MouseButton) {
			if e.Button != event.RENG_BUTTON_LEFT {
				return
			}
			if buttons, ok := g.Event.Button[g.SayScreenName]; !ok {
				lock <- 0
			} else {
				for _, button := range buttons {
					if button.IsNowDown {
						return
					}
				}
				lock <- 0
			}
		},
	})
	<-lock
	return nil
}
