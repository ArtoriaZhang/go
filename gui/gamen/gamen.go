package main

import (
	"runtime"

	"github.com/rajveermalviya/gamen/display"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	d, err := display.NewDisplay()
	if err != nil {
		panic(err)
	}
	defer d.Destroy()

	w, err := display.NewWindow(d)
	if err != nil {
		panic(err)
	}
	defer w.Destroy()

	w.SetCloseRequestedCallback(func() { d.Destroy() })

	for {
		if !d.Poll() {
			break
		}

		// render()
	}
}
