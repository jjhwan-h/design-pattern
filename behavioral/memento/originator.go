package main

type Editor struct {
	x, y int
}

func (e *Editor) Save() *posMemento {
	return &posMemento{
		e.x,
		e.y,
	}
}

func (e *Editor) Restore(m *posMemento) {
	if m == nil {
		return
	}

	e.x = m.x
	e.y = m.y
}

func (e *Editor) Set(x, y int) {
	e.x = x
	e.y = y
}
