package factory

import "fmt"

type Vehicler interface {
	SetName(name string)
	SetColor(color string)
	MoveForward()
	MoveBackward()
}

type vehicle struct {
	name  string
	color string
}

func (v *vehicle) SetName(name string) {
	v.name = name
}

func (v *vehicle) SetColor(color string) {
	v.color = color
}

func (v *vehicle) MoveForward() {
	fmt.Print("moving forward")
}

func (v *vehicle) MoveBackward() {
	fmt.Print("moving backward")
}
