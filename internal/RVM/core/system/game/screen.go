package game

import (
	"internal/RVM/core/object"
	"internal/RVM/core/system/game/event"
	"time"
)

func (g *Game) screenEval(
	so []object.ScreenObject,
	name string, bps int,
) (err error) {
	for _, obj := range so {
		switch obj := obj.(type) {
		case *object.Code:
			g.lock.Unlock()
			obj.Func()
			g.lock.Lock()
		case *object.Show:
			g.evalShow(obj, name, bps)
		case *object.PlayMusic:
			g.NowMusic = obj.Path
			err = g.Audio.PlayMusic(g.path+obj.Path, obj.Loop, obj.Ms)
		case *object.StopMusic:
			g.NowMusic = ""
			g.Audio.StopMusic(obj.Ms)
		case *object.PlayChannel:
			err = g.Audio.PlayChannel(obj.ChanName, g.path+obj.Path)
		case *object.PlayVideo:
			g.evalPlayVideo(obj, name, bps)
		case *object.Key:
			g.Event.AddKeyEvent(name, event.KeyEvent{Down: obj.Down, Up: obj.Up})
		case *object.Button:
			g.evalButton(obj, name, bps)
		case *object.Bar:
			g.evalBar(obj, name, bps)
		case *object.Text:
			g.evalText(obj, name, bps)
		case *object.TextPointer:
			g.evalTextPointer(obj, name, bps)
		case *object.Timer:
			go func() {
				time.Sleep(time.Duration(obj.Time) * time.Second)
				obj.Do()
			}()
		}
	}
	return err
}
