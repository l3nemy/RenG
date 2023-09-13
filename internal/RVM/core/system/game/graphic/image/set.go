package image

/*
#cgo LDFLAGS: -lSDL2 -lSDL2main -lSDL2_image

#include <SDL2/SDL.h>
#include <SDL2/SDL_image.h>
*/
import "C"
import (
	"log"
)

func (i *Image) SetImageAlpha(name string, alpha uint8) {
	i.lock.Lock()
	defer i.lock.Unlock()
	if image, ok := i.images[name]; !ok {
		log.Fatalf("Image Name Error : got - %s", name)
	} else {
		i.ChangeTextureAlpha(image.texture, alpha)
	}
}
