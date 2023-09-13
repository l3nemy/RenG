package image

/*
#cgo LDFLAGS: -lSDL2 -lSDL2main -lSDL2_image

#include <SDL2/SDL.h>
#include <SDL2/SDL_image.h>
*/
import "C"
import "internal/RVM/core/globaltype"

func (i *Image) FreeSurface(sur *C.SDL_Surface) {
	C.SDL_FreeSurface(sur)
}

func (i *Image) FreeTexture(tex *globaltype.SDL_Texture) {
	C.SDL_DestroyTexture((*C.SDL_Texture)(tex))
}
