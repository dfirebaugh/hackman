package screen

import (
	"hackman/internal/net"
	"hackman/internal/ui/component/addmember"
	"hackman/internal/ui/component/form"
	"hackman/internal/ui/component/login"
	"hackman/internal/ui/component/member"
	"hackman/internal/ui/component/memberlist"
	"hackman/internal/ui/component/menu"
	"hackman/internal/ui/component/modal"
	"hackman/internal/ui/component/search"
	"hackman/internal/ui/page"
	"hackman/internal/ui/window"

	"github.com/rivo/tview"
)

type screen struct {
	wm window.WindowManager
}

var app = tview.NewApplication()
var Screen = screen{
	wm: window.WindowManager{
		App: app,
	},
}

func init() {
	net.NetService.Screen = Screen
}

func (s screen) Show(p page.Page) tview.Primitive {
	m := s.showModal()
	if m != nil {
		return m
	}

	var pageMap map[page.Page]tview.Primitive = map[page.Page]tview.Primitive{
		page.Login:        s.wm.BuildWindow(login.View(s), "login"),
		page.Loading:      modal.Modals.BuildModal(s),
		page.HomeMenu:     s.wm.BuildWindow(menu.View(s), "home"),
		page.AddMember:    s.wm.BuildWindow(addmember.View(s), "addmember"),
		page.Search:       s.wm.BuildWindow(search.View(s), "search"),
		page.MemberList:   s.wm.BuildWindow(memberlist.View(s), "members"),
		page.MemberEditor: s.wm.BuildWindow(member.View(s), "member editor"),
		page.CurrentForm:  form.Current,
	}
	if !net.NetService.IsLoggedIn() {
		loginForm := pageMap[p]
		app.SetRoot(loginForm, true)
		return loginForm
	}

	if t, ok := pageMap[p]; ok {
		app.SetRoot(t, true)
		return t
	}
	return tview.NewBox().SetBorder(true).SetTitle("Hack Manager!")
}

func (s screen) showModal() tview.Primitive {
	if modal.Modals.Len() == 0 {
		return nil
	}
	m := modal.Modals.BuildModal(s)
	app.SetRoot(m, true)
	return m
}

func (s screen) Run() {
	if err := app.SetRoot(s.Show(page.Login), true).Run(); err != nil {
		panic(err)
	}
}

func (s screen) Stop() {
	app.Stop()
}
