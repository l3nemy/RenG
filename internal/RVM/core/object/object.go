package object

type Object interface {
	ScreenObject
	LabelObject
}

type ScreenObject interface {
	screenObj()
}

type LabelObject interface {
	labelObj()
}
