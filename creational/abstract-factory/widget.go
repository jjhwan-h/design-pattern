package main

type WidgetFactory interface {
	NewButton() Button
	NewCheckBox() CheckBox
}
