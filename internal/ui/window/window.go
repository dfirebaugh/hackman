package window

import (
	"github.com/epiclabs-io/winman"
	"github.com/rivo/tview"
)

type WindowManager struct {
	App *tview.Application
}

var wm = winman.NewWindowManager()

func (w WindowManager) BuildWindow(content tview.Primitive, title string) *winman.WindowBase {
	return wm.NewWindow(). // create new window and add it to the window manager
				Show().             // make window visible
				SetRoot(content).   // have the text view above be the content of the window
				SetDraggable(true). // make window draggable around the screen
				SetResizable(true). // make the window resizable
				SetTitle(title)
}
