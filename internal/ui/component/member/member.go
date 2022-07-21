package member

import (
	"fmt"
	"hackman/internal/model"
	"hackman/internal/net"
	"hackman/internal/store"
	"hackman/internal/ui/component"
	"hackman/internal/ui/component/form"
	"hackman/internal/ui/component/modal"
	"hackman/internal/ui/page"
	"hackman/pkg/paypal"

	"github.com/rivo/tview"
)

func View(s component.Screen) *tview.List {
	return tview.NewList().
		AddItem(store.CurrentMember.Name, "edit name", 0, func() {
			form.Current = form.BuildForm([]form.Property{
				{
					Label: "name",
					Value: store.CurrentMember.Name,
					Width: 30,
					OnChange: func(value string) {
						store.CurrentMember.Name = value
					},
				},
			}).
				AddButton("submit", func() {
					// send request to update
					net.NetService.UpdateMember(model.Member{
						Email: store.CurrentMember.Email,
						Name:  store.CurrentMember.Name,
						// FullName?
						FullName:       store.CurrentMember.Name,
						SubscriptionID: store.CurrentMember.SubscriptionID,
					})
					s.Show(page.MemberEditor)
				})
			s.Show(page.CurrentForm)
		}).
		AddItem(store.CurrentMember.Email, "edit email", 0, nil).
		AddItem(store.CurrentMember.RFID, "edit rfid", 0, func() {
			form.Current = form.BuildForm([]form.Property{
				{
					Label: "rfid",
					Value: store.CurrentMember.RFID,
					Width: 30,
					OnChange: func(value string) {
						store.CurrentMember.RFID = value
					},
				},
			}).
				AddButton("submit", func() {
					// send request to update
					net.NetService.UpdateRFID(model.Member{
						Email: store.CurrentMember.Email,
						RFID:  store.CurrentMember.RFID,
					})
					s.Show(page.MemberEditor)
				})
			s.Show(page.CurrentForm)
		}).
		AddItem(store.CurrentMember.SubscriptionID, "edit subscription id", 0, func() {
			form.Current = form.BuildForm([]form.Property{
				{
					Label: "subscriptionID",
					Value: store.CurrentMember.SubscriptionID,
					Width: 30,
					OnChange: func(value string) {
						store.CurrentMember.SubscriptionID = value
					},
				},
			}).
				AddButton("submit", func() {
					// send request to update
					net.NetService.UpdateMember(model.Member{
						Email: store.CurrentMember.Email,
						Name:  store.CurrentMember.Name,
						// FullName?
						FullName:       store.CurrentMember.Name,
						SubscriptionID: store.CurrentMember.SubscriptionID,
					})
					s.Show(page.MemberEditor)
				})
			s.Show(page.CurrentForm)
		}).
		AddItem(model.MemberLevelToStr[model.MemberLevel(store.CurrentMember.MemberLevel)], "member level", 0, nil).
		AddItem("check paypal status", "fetch subscription info from paypal", 'c', func() {
			status, lastPayment, err := paypal.Paypal{}.GetSubscription(store.CurrentMember.SubscriptionID)
			resultString := fmt.Sprintf("status: %s \nlast payment: %s ", status, lastPayment.Time.Format("2006-02-01"))
			if err != nil {
				resultString = err.Error()
			}
			modal.Modals.Push(modal.Modal{
				Text: resultString,
				Buttons: []modal.ModalButton{
					{
						Label: "back",
						Func: func() {
							s.Show(page.MemberEditor)
						},
					},
				},
			})

			s.Show(page.MemberEditor)
		}).
		AddItem("back", "back to menu", 'q', func() {
			s.Show(page.HomeMenu)
		})
}
