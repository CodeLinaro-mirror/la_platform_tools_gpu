package gles

import "fmt"

func (c Color) String() string {
	return fmt.Sprintf("R:% 6f, G:% 6f, B:% 6f, A:% 6f", c.Red, c.Green, c.Blue, c.Alpha)
}

func (v Vec2f) String() string {
	return fmt.Sprintf("(% 6f, % 6f)", v.X, v.Y)
}

func (v Vec3f) String() string {
	return fmt.Sprintf("(% 6f, % 6f, % 6f)", v.X, v.Y, v.Z)
}

func (v Vec4f) String() string {
	return fmt.Sprintf("(% 6f, % 6f, % 6f, % 6f)", v.X, v.Y, v.Z, v.W)
}

func (v Vec2i) String() string {
	return fmt.Sprintf("(% 6d, % 6d)", v.X, v.Y)
}

func (v Vec3i) String() string {
	return fmt.Sprintf("(% 6d, % 6d, % 6d)", v.X, v.Y, v.Z)
}

func (v Vec4i) String() string {
	return fmt.Sprintf("(% 6d, % 6d, % 6d, % 6d)", v.X, v.Y, v.Z, v.W)
}

func (m Mat2f) String() string {
	return fmt.Sprintf("[%s, %s]", m.Col0, m.Col1)
}

func (m Mat3f) String() string {
	return fmt.Sprintf("[%s, %s, %s]", m.Col0, m.Col1, m.Col2)
}

func (m Mat4f) String() string {
	return fmt.Sprintf("[%s, %s, %s, %s]", m.Col0, m.Col1, m.Col2, m.Col3)
}

func (u Uniform) String() string {
	return fmt.Sprintf("%v %v", u.Name, u.Value)
}

func (a VertexAttributeArray) String() string {
	if a.Enabled {
		return fmt.Sprintf("%d x %v", int(a.Size), a.Type)
	} else {
		return "disabled"
	}
}
