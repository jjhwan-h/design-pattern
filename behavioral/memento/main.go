package main

import "fmt"

func main() {
	editor := &Editor{}
	history := &History{}

	editor.Set(10, 1)
	history.Push(editor.Save())

	fmt.Println(editor.x, editor.y)
	editor.Set(11, 12)

	fmt.Println(editor.x, editor.y)
	editor.Restore(history.Pop())

	fmt.Println(editor.x, editor.y)
}
