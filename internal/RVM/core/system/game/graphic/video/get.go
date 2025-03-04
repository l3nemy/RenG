package video

/*
#cgo LDFLAGS: -lSDL2 -lSDL2main

#include <SDL2/SDL.h>
*/
import "C"
import "internal/RVM/core/globaltype"

func (v *Video) GetVideoNameANDLoopByTexture(t *globaltype.SDL_Texture) (string, int) {
	v.Lock()
	defer v.Unlock()

	for name, video := range v.V {
		if video.texture == (*C.SDL_Texture)(t) {
			return name, int(video.loop)
		}
	}

	return "", 0
}

func (v *Video) GetVideoTexture(name string) *C.SDL_Texture {
	v.Lock()
	defer v.Unlock()

	if video, ok := v.V[name]; ok {
		return video.texture
	}
	return nil
}

func (v *Video) GetNowPlaying(name string) bool {
	v.Lock()
	defer v.Unlock()

	if _, ok := v.V[name]; ok {
		return v.V[name].nowPlaying == 1
	}
	return false
}
