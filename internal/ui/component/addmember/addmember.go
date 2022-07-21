package addmember

import (
	"hackman/internal/ui/component"
	"hackman/internal/ui/page"

	"github.com/rivo/tview"
)

func View(s component.Screen) *tview.Form {
	var email string
	var rfid string
	return tview.NewForm().
		AddInputField("email", email, 30, nil, func(text string) {
			email = text
		}).
		AddInputField("rfid", rfid, 30, nil, func(text string) {
			rfid = text
		}).
		AddButton("cancel", func() {
			s.Show(page.HomeMenu)
		}).
		AddButton("submit", func() {
			s.Show(page.HomeMenu)
		})
}
