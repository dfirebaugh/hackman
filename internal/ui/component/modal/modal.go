package modal

import (
	"hackman/internal/ui/component"
	"hackman/internal/ui/page"

	"github.com/rivo/tview"
)

type ModalStack struct {
	elements []Modal
}

type ModalButton struct {
	Label string
	Func  func()
}
type Modal struct {
	Text    string
	Buttons []ModalButton
}

var Modals = &ModalStack{}

func (ms ModalStack) OpenModal(s component.Screen, msg string, p page.Page) {
	ms.Push(Modal{
		Text: msg,
		Buttons: []ModalButton{{
			Label: "continue",
			Func: func() {
				s.Show(p)
			},
		}},
	})
}

// BuildModal pops a modal of the top of the stack and builds a *tview.Modal
func (ms ModalStack) BuildModal(s component.Screen) *tview.Modal {
	m := tview.NewModal()

	if Modals.IsEmpty() {
		return tview.NewModal().
			SetText("no response yet").
			AddButtons([]string{"refresh", "cancel"}).
			SetDoneFunc(func(index int, label string) {
				if label == "cancel" {
					s.Show(page.HomeMenu)
				}

				if label == "refresh" {
					s.Show(page.HomeMenu)
				}
			})
	}

	top := Modals.Pop()
	m.SetText(top.Text)
	for _, b := range top.Buttons {
		m.AddButtons([]string{b.Label})
		m.SetDoneFunc(func(index int, label string) {
			if label == b.Label {
				if b.Func == nil {
					return
				}
				b.Func()
			}
		})
	}
	return m
}

func (ms *ModalStack) Push(m Modal) {
	ms.elements = append(ms.elements, m)
}

func (ms *ModalStack) Pop() Modal {
	top := ms.elements[len(ms.elements)-1]
	ms.elements = ms.elements[:len(ms.elements)-1]

	return top
}

func (ms ModalStack) Len() int {
	return len(ms.elements)
}
func (ms ModalStack) IsEmpty() bool {
	return len(ms.elements) == 0
}

// Default - return a modal but don't push one to stack
func (ms ModalStack) Default(s component.Screen) Modal {
	return Modal{
		Text: "no response yet",
		Buttons: []ModalButton{
			{
				Label: "cancel",
				Func: func() {
					s.Show(page.HomeMenu)
				},
			},
			{
				Label: "refresh",
				Func: func() {
					s.Show(page.MemberList)
				},
			},
		},
	}
}
