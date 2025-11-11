package main

type LightButton struct{}

func (LightButton) Render() string { return "light button" }

type LightCheckBox struct{ on bool }

func (c *LightCheckBox) Check(on bool) { c.on = on }
func (c *LightCheckBox) Status() bool  { return c.on }

type LightFactory struct{}

func (LightFactory) NewButton() Button     { return LightButton{} }
func (LightFactory) NewCheckBox() CheckBox { return &LightCheckBox{} }

type DarkButton struct{}

func (DarkButton) Render() string { return "dark button" }

type DarkCheckBox struct{ on bool }

func (c *DarkCheckBox) Check(on bool) { c.on = on }
func (c *DarkCheckBox) Status() bool  { return c.on }

type DarkFactory struct{}

func (DarkFactory) NewButton() Button     { return DarkButton{} }
func (DarkFactory) NewCheckBox() CheckBox { return &DarkCheckBox{} }
