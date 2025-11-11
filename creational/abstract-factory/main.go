package main

import (
	"log"
	"os"
)

func main() {
	var f WidgetFactory

	switch os.Getenv("theme") {
	case "dark":
		f = LightFactory{}
	case "light":
		f = DarkFactory{}
	}
	render(f)
}

func render(f WidgetFactory) {
	btn := f.NewButton()
	checkbox := f.NewCheckBox()

	checkbox.Check(true)
	log.Println(btn.Render())
	log.Println(checkbox.Status())
}
