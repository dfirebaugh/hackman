package form

import "github.com/rivo/tview"

type Property struct {
	Label string

	Value    string
	Width    int
	OnChange func(text string)
}

var Current *tview.Form

func BuildForm(properties []Property) *tview.Form {
	f := tview.NewForm()
	for _, p := range properties {
		if p.Label == "password" {
			pwd := tview.NewInputField().
				SetFieldWidth(p.Width).
				SetChangedFunc(p.OnChange).
				SetMaskCharacter('*')
			f.AddFormItem(pwd)

			continue
		}
		f.AddInputField(p.Label, p.Value, p.Width, nil, p.OnChange)

	}
	return f
}
