package globaltype

/*
#cgo CFLAGS: -I./../sdl/include
#cgo LDFLAGS: -L./../sdl/lib -lSDL2 -lSDL2main -lSDL2_image -lSDL2_ttf

#include <SDL2/SDL.h>
#include <SDL2/SDL_image.h>
#include <SDL2/SDL_ttf.h>
*/
import "C"

type (
	SDL_Window   = C.SDL_Window
	SDL_Renderer = C.SDL_Renderer
	SDL_Texture  = C.SDL_Texture
	SDL_Surface  = C.SDL_Surface

	TTF_Font = C.TTF_Font
)
