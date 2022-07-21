package search

import (
	"fmt"
	"hackman/internal/store"
	"hackman/internal/ui/component"
	"hackman/internal/ui/component/modal"
	"hackman/internal/ui/page"

	"github.com/rivo/tview"
)

func View(s component.Screen) *tview.Form {
	var searchTerm string
	return tview.NewForm().
		AddInputField("search term", searchTerm, 30, nil, func(text string) {
			searchTerm = text
		}).
		AddButton("submit", func() {
			s.Show(page.Loading)
			Search(searchTerm, s)
		})
}

func Search(term string, s component.Screen) {
	m := store.Search(term)
	modal.Modals.Push(modal.Modal{
		Text: fmt.Sprintf("found %d members", len(m)),
		Buttons: []modal.ModalButton{{
			Label: "continue",
			Func: func() {
				s.Show(page.MemberList)
			},
		}},
	})
}
