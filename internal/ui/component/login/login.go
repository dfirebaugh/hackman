package login

import (
	"hackman/internal/net"
	"hackman/internal/store"
	"hackman/internal/ui/component"
	"hackman/internal/ui/component/form"
	"hackman/internal/ui/page"

	"github.com/rivo/tview"
)

func View(s component.Screen) *tview.Form {
	return form.BuildForm([]form.Property{
		{
			Label: "username",
			Value: store.CurrentUsername,
			Width: 30,
			OnChange: func(text string) {
				store.CurrentUsername = text
			},
		},
		{
			Label: "password",
			Value: store.CurrentUserPass,
			Width: 30,
			OnChange: func(text string) {
				store.CurrentUserPass = text
			},
		},
	}).
		AddButton("submit", func() {
			s.Show(page.Loading)
			if net.NetService.Login(store.CurrentUsername, store.CurrentUserPass) {
				// user := net.NetService.GetUser()

				// modal.Modals.Push(modal.Modal{
				// 	Text: fmt.Sprintf("Welcome, %s", user.Name),
				// 	Buttons: []modal.ModalButton{{
				// 		Label: "continue",
				// 		Func: func() {
				// 			s.Show(page.HomeMenu)
				// 		},
				// 	}},
				// })
			}
		})
}
