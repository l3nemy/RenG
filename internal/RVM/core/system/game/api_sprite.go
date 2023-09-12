package game

import "internal/RVM/core/object"

func (g *Game) StartSprite(
	screenName,
	spriteName string,
	duration float64,
	loop bool,
	count,
	xsize,
	ysize int,
	T object.Transform,
) {
	g.lock.Lock()
	defer g.lock.Unlock()

	g.Graphic.Sprite_Manager.CreateSpriteImages(spriteName, count, xsize, ysize)

	if T.Size.X != 0 && T.Size.Y != 0 {
		T = g.echoTransform(T, T.Size.X, T.Size.Y)
	} else {
		T = g.echoTransform(T, xsize, ysize)
	}

	g.Graphic.AddScreenTextureRenderBuffer(
		g.screenBps[screenName],
		g.Graphic.Sprite_Manager.GetSpriteImage(spriteName, 0),
		T,
	)

	index := g.Graphic.GetCurrentTopScreenIndexByBps(g.screenBps[screenName])

	g.Graphic.AddSprite(screenName, spriteName, g.screenBps[screenName], index, duration, loop)
}
