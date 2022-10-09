package system

/*
#cgo CFLAGS: -I./../sdl/include
#cgo LDFLAGS: -L./../sdl/lib -lSDL2 -lSDL2main -lSDL2_image -lSDL2_ttf -lSDL2_mixer

#include <SDL.h>
#include <SDL_image.h>
#include <SDL_ttf.h>
#include <SDL_mixer.h>
*/
import "C"
import (
	"RenG/RVM/src/core/st"
	"unsafe"
)

type System struct {
	window   *st.SDL_Window
	renderer *st.SDL_Renderer
	event    st.SDL_Event
}

func Init(title string, width, height int) *System {
	system := &System{}

	if C.SDL_Init(st.SDL_INIT_EVERYTHING) < 0 {
		return nil
	}

	Ctitle := C.CString(title)
	defer C.free(unsafe.Pointer(Ctitle))

	window := C.SDL_CreateWindow(
		Ctitle, st.SDL_WINDOWPOS_CENTERED, st.SDL_WINDOWPOS_CENTERED,
		C.int(width), C.int(height),
		st.SDL_WINDOW_SHOWN|st.SDL_WINDOW_RESIZABLE|st.SDL_WINDOW_INPUT_FOCUS|st.SDL_WINDOW_MOUSE_FOCUS,
	)
	if window == nil {
		return nil
	}

	renderer := C.SDL_CreateRenderer(window, -1,
		st.SDL_RENDERER_ACCELERATED|st.SDL_RENDERER_PRESENTVSYNC|st.SDL_RENDERER_TARGETTEXTURE,
	)
	if renderer == nil {
		return nil
	}

	system.window = (*st.SDL_Window)(window)
	system.renderer = (*st.SDL_Renderer)(renderer)

	return system
}

func (s *System) Close() {
	C.SDL_DestroyWindow(s.window)
	C.SDL_DestroyRenderer(s.renderer)

	C.SDL_Quit()
}