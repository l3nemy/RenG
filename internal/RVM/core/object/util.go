package object

type Code struct {
	Func func()
}

func (c *Code) labelObj()  {}
func (c *Code) screenObj() {}

func (c *Code) Serialize() {

}
func (c *Code) Deserialize() {

}

type Vector2 struct {
	X, Y float32
}

type Color struct {
	R uint8
	G uint8
	B uint8
	A uint8
}
