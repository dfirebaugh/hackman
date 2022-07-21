package memberlist

import (
	"fmt"
	"hackman/internal/model"
	"hackman/internal/store"
	"hackman/internal/ui/component"
	"hackman/internal/ui/page"

	"github.com/rivo/tview"
)

func View(s component.Screen) tview.Primitive {
	memberList := tview.NewList()

	for _, m := range store.CurrentList {
		memberList.
			AddItem(m.Name, fmt.Sprintf("%s | %s | %s", m.Email, m.RFID, model.MemberLevelToStr[model.MemberLevel(m.MemberLevel)]), 0, func() func() {
				// using closure to keep track of which member is associated with the list element
				mem := m

				return func() {
					store.CurrentMember = mem
					s.Show(page.MemberEditor)
				}
			}())
	}
	return memberList.
		AddItem("Back", "Back to main menu", 'q', func() {
			s.Show(page.HomeMenu)
		})
}
