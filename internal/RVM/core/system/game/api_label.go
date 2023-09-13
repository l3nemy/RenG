package game

import (
	"internal/RVM/core/object"
	"time"
)

func (g *Game) startLabel(name string, index int, sayScreenName string) (err error) {
	g.SayScreenName = sayScreenName

	// label을 screen처럼 등록
	bps, ok := g.screenBps[name]
	if !ok {
		g.screenBps[name] = g.Graphic.GetCurrentTopRenderBps() + 1
		bps = g.screenBps[name]
		g.Graphic.AddScreenRenderBuffer()
	}

	//config 정보 삽입
	g.LabelManager.SetNowLabelName(name)
	g.LabelManager.SetNowLabelIndex(index)
	g.LabelManager.AddCallStack(name, index)
	g.Event.TopScreenName = name

	for {
		err = g.labelEval(g.LabelManager.GetNowLabelObject(), g.LabelManager.GetNowLabelName(), bps)
		if err != nil {
			return err
		}

		if !g.LabelManager.NextLabelObject() {
			break
		}
	}

	g.InActiveScreen(g.SayScreenName)
	return
}

func (g *Game) labelEval(obj object.LabelObject, name string, bps int) (err error) {
	switch obj := obj.(type) {
	case *object.Code:
		obj.Func()
	case *object.Jump:
		g.LabelManager.JumpLabel(obj.LabelName)
	case *object.Call:
		g.LabelManager.CallLabel(obj.LabelName)
	case *object.Show:
		g.evalShow(obj, name, bps)
	case *object.Hide:
		g.evalHide(obj, name, bps)
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
	case *object.Say:
		err = g.evalSay(obj)
	case *object.Pause:
		g.InActiveScreen(g.SayScreenName)
		time.Sleep(time.Duration(obj.Time) * time.Second)
	}
	return
}
