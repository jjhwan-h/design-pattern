package main

type History struct {
	stack []*posMemento
}

func (h *History) Push(m *posMemento) {
	h.stack = append(h.stack, m)
}

func (h *History) Pop() *posMemento {
	if len(h.stack) == 0 {
		return nil
	}
	m := h.stack[len(h.stack)-1]
	h.stack = h.stack[:len(h.stack)-1]
	return m
}
