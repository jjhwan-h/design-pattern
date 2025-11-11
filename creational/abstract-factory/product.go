package main

type Button interface {
	Render() string
}

type CheckBox interface {
	Check(on bool)
	Status() bool
}
