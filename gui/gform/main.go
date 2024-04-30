package main

import (
	"github.com/AllenDang/gform"
)

func main() {
	gform.Init()

	dialog := gform.NewDialogFromResId(nil, 101) //101 is the resource Id.
	dialog.Center()
	dialog.Show()

	edt = gform.AttachEdit(dialog, 1000)
	edt.SetCaption("Hello")

	btn := gform.AttachPushButton(dialog, 2)
	btn.OnLBDown().Attach(onclick)

	gform.RunMainLoop()
}
