package game

import "internal/RVM/core/object"

func (g *Game) AddAnimation(a *object.Anime, screenName string, textureIndex int) {
	g.lock.Lock()
	defer g.lock.Unlock()

	g.Graphic.AddAnimation(screenName, a, g.screenBps[screenName], textureIndex)
}
