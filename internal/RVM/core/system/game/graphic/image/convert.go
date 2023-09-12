package image

/*
#cgo CFLAGS: -I./../../../../sdl/include
#cgo CFLAGS: -I./../../../../system/game/graphic/image/c
#cgo LDFLAGS: -L./../../../../sdl/lib -lSDL2 -lSDL2main -lSDL2_image

#include <SDL2/SDL.h>
*/
import "C"
import (
	"internal/RVM/core/globaltype"
	"internal/RVM/core/system/game/graphic/image/pixel"
)

func (i *Image) ConvertFrameDataToYUV(data [8]*uint8, linesize [8]int32, Width, Height int64) pixel.YUV {
	yuv := pixel.NewYUV(pixel.RENG_PIXELFORMAT_YUV420)

	yuv.SetYUV(
		data[0],
		data[1],
		data[2],
	)

	yuv.SetWH(
		Width,
		Height,
	)

	return yuv
}

func (i *Image) ConvertSurfaceToTexture(sur *C.SDL_Surface) *globaltype.SDL_Texture {
	return (*globaltype.SDL_Texture)(C.SDL_CreateTextureFromSurface((*C.SDL_Renderer)(i.renderer), sur))
}
