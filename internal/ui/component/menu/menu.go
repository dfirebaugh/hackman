package menu

import (
	"hackman/internal/ui/component"
	"hackman/internal/ui/page"

	"github.com/rivo/tview"
)

func View(s component.Screen) *tview.List {
	return tview.NewList().
		AddItem("Search", "Search for a member by name", 's', func() {
			s.Show(page.Search)
		}).
		// AddItem("All", "List out all active members", 'l', func() {
		// 	s.Show(page.MemberList)
		// }).
		AddItem("Add New Member", "if a member doesn't exist, we can give them temporary access", 'a', func() {
			s.Show(page.AddMember)
		}).
		AddItem("Quit", "Press to exit", 'q', func() {
			s.Stop()
		})
}
