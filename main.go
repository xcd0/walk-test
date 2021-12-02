package main

import (
	. "github.com/lxn/walk/declarative"
)

func main() {
	Run()
	// something
	Run()
}

func Run() {
	var w *MainWindow
	w = &MainWindow{
		Title:    "test",
		Size:     Size{400, 200},
		Layout:   VBox{},
		AssignTo: &w,
		Children: []Widget{
			PushButton{
				Text: "Exit",
				OnClicked: func() {
					(**w.AssignTo).Close()
				},
			},
		},
	}
	(*w).Run()
}
