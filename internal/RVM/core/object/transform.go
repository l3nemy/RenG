package object

type Transform struct {
	Pos    Vector2
	Size   Vector2
	Flip   Vector2
	Rotate float32

	Type SpecialTransform
}

/*--------- 특수 transform ---------*/

type SpecialTransform interface {
	specialTransformObj()
}

type Center struct {
}

func (c *Center) specialTransformObj() {}

type XCenter struct {
	Ypos float32
}

func (xc *XCenter) specialTransformObj() {}

type YCenter struct {
	Xpos float32
}

func (yc *YCenter) specialTransformObj() {}

type AxisCenter struct {
	Axis Vector2
}

func (ac *AxisCenter) specialTransformObj() {}
