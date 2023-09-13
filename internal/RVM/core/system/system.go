package system

/*
#cgo LDFLAGS: -lSDL2 -lSDL2main -lSDL2_image -lSDL2_ttf -lSDL2_mixer

#include <SDL2/SDL.h>
#include <SDL2/SDL_image.h>
#include <SDL2/SDL_ttf.h>
#include <SDL2/SDL_mixer.h>
*/
import "C"
import (
	"internal/RVM/core/globaltype"
	"internal/RVM/core/system/game"
	"internal/RVM/core/system/game/audio"
	"internal/RVM/core/system/game/graphic"
	"log"
	"os"
	"unsafe"
)

type System struct {
	window *globaltype.SDL_Window

	Title           string
	IsNowFullScreen bool

	// Public : 게임 객체입니다. 여러가지 게임 진행에 필요한 것들이 담겨있습니다.
	game *game.Game
}

/*
Public

[title] : 게임 제목

[width], [height] : 창 너비, 높이

[CursorPath] : 커서 이미지 경로
(사이즈는 조정되지 않으므로 이미지 크기를 맞춰주세요.)

초기화 함수입니다. 이때 윈도우는 생성되지만 화면에 표시되지 않습니다.
WindowStart() 함수에서 화면에 표시되기 시작합니다.
*/
func Init(title string,
	width, height int,
	CursorPath *string,
	NowCharacter *string,
	NowText *string,
) *System {

	if C.SDL_Init(C.SDL_INIT_EVERYTHING) < 0 {
		return nil
	}

	Ctitle := C.CString(title)
	defer C.free(unsafe.Pointer(Ctitle))

	window := C.SDL_CreateWindow(
		Ctitle, C.SDL_WINDOWPOS_CENTERED, C.SDL_WINDOWPOS_CENTERED,
		C.int(width), C.int(height),
		C.SDL_WINDOW_HIDDEN|C.SDL_WINDOW_INPUT_FOCUS|C.SDL_WINDOW_MOUSE_FOCUS,
	)
	if window == nil {
		return nil
	}

	renderer := C.SDL_CreateRenderer(window, -1,
		C.SDL_RENDERER_ACCELERATED|C.SDL_RENDERER_PRESENTVSYNC|C.SDL_RENDERER_TARGETTEXTURE,
	)
	if renderer == nil {
		return nil
	}

	C.TTF_Init()

	hint1 := C.CString(C.SDL_HINT_RENDER_SCALE_QUALITY)
	defer C.free(unsafe.Pointer(hint1))

	hint2 := C.CString("2")
	defer C.free(unsafe.Pointer(hint2))

	if C.SDL_SetHint(hint1, hint2) == 0 {
		log.Println("Hint quality Error")
	}

	path, _ := os.Getwd()

	g := graphic.Init((*globaltype.SDL_Window)(window), (*globaltype.SDL_Renderer)(renderer), path, width, height)
	if CursorPath != nil {
		g.RegisterCursor(*CursorPath)
	}

	return &System{
		window:          (*globaltype.SDL_Window)(window),
		Title:           title,
		IsNowFullScreen: false,
		game: game.Init(
			g,
			audio.Init(),
			path,
			width, height,
			NowCharacter,
			NowText,
		),
	}
}

func (s *System) Close() {
	C.SDL_DestroyWindow((*C.SDL_Window)(s.window))

	s.game.Close()
	C.SDL_Quit()
}

func (s *System) WindowStart(
	firstScreen string,
) {
	s.game.ActiveScreen(firstScreen)
	C.SDL_ShowWindow((*C.SDL_Window)(s.window))

	for {
		if s.game.Event.Update() {
			break
		}
		s.game.Graphic.Update()
		s.game.Graphic.Render()
	}
}

func (s *System) ToggleFullScreen() {
	if !s.IsNowFullScreen {
		C.SDL_SetWindowFullscreen((*C.SDL_Window)(s.window), C.SDL_WINDOW_FULLSCREEN_DESKTOP)
		s.IsNowFullScreen = true
	} else {
		C.SDL_SetWindowFullscreen((*C.SDL_Window)(s.window), 0)
		s.IsNowFullScreen = false
	}
}
