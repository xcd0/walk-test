package main

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func main() {
	Run()
	// something
	Run()
}
func Run() {
	var mw *walk.MainWindow
	w := &MainWindow{
		Title:    "test",
		Size:     Size{400, 200},
		Layout:   VBox{},
		AssignTo: &mw,
		Children: []Widget{
			PushButton{
				Text: "Exit",
				OnClicked: func() {
					mw.Close()
				},
			},
		},
	}
	(*w).Run()
}
