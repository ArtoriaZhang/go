package main

import (
	"os"

	"github.com/kitech/qt.go/qtwidgets"
)

func main() {
	app := qtwidgets.NewQApplication(len(os.Args), os.Args, 0)
	btn := qtwidgets.NewQPushButton1("hello qt.go", nil)
	btn.Show()
	app.Exec()
}
