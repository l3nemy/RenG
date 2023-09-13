package image

/*
#cgo CFLAGS: -I./../../../../system/game/graphic/image/c
#cgo LDFLAGS: -lSDL2 -lSDL2main -lSDL2_image

#include <filter.h>
*/
import "C"

func (i *Image) Blur(ImageName string, xrad, yrad float32) *C.SDL_Surface {
	return C.Blur(i.images[ImageName].surface, C.float(xrad), C.float(yrad))
}

func (i *Image) BlurSurface(Surface *C.SDL_Surface, xrad, yrad float32) *C.SDL_Surface {
	return C.Blur(Surface, C.float(xrad), C.float(yrad))
}
