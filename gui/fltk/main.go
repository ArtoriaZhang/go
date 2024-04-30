package main

import "github.com/pwiecz/go-fltk"

func main() {
	win := fltk.NewWindow(400, 300)
	win.SetLabel("Main Window")
	btn := fltk.NewButton(160, 200, 80, 30, "Click")
	btn.SetCallback(func() {
		btn.SetLabel("Clicked")
	})
	win.End()
	win.Show()
	fltk.Run()
}
