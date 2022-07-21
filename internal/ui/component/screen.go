package component

import (
	"hackman/internal/ui/page"

	"github.com/rivo/tview"
)

type Screen interface {
	Show(p page.Page) tview.Primitive
	Stop()
}
