package game

func (g *Game) ChangeTextureAlpha(name string, alpha uint8) {
	g.lock.Lock()
	defer g.lock.Unlock()

	g.Graphic.Image_Manager.ChangeTextureAlpha(g.Graphic.Image_Manager.GetImageTexture(name), alpha)
}
