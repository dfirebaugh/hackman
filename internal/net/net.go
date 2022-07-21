package net

import (
	"bytes"
	"encoding/json"
	"fmt"
	"hackman/config"
	"hackman/internal/model"
	"hackman/internal/store"
	"hackman/internal/ui/component/modal"
	"hackman/internal/ui/page"
	"io"
	"net/http"
	"time"

	"github.com/rivo/tview"
)

type screen interface {
	Show(p page.Page) tview.Primitive
	Stop()
}

type LoginResponse struct {
	Token string `json:"token"`
}

type netService struct {
	Screen  screen
	client  http.Client
	cookie  *http.Cookie
	apiRoot string
}

var (
	NetService = &netService{
		client:  http.Client{Timeout: 5 * time.Second},
		apiRoot: config.Config.ServiceURL + "",
	}
)

func (n *netService) IsLoggedIn() bool {
	return n.cookie.Valid() == nil
	// return store.CurrentUsername != "" && store.CurrentUserPass != ""
}

func (n *netService) updateCachedCred(username string, password string) {
	store.CurrentUsername = username
	store.CurrentUserPass = password
}

func (n *netService) OpenModal(msg string, p page.Page) {
	modal.Modals.Push(modal.Modal{
		Text: msg,
		Buttons: []modal.ModalButton{{
			Label: "continue",
			Func: func() {
				if !n.IsLoggedIn() {
					n.Screen.Show(page.HomeMenu)
					return
				}

				n.Screen.Show(p)
			},
		}},
	})
}

func (n *netService) Login(username string, password string) bool {
	req, err := http.NewRequest(http.MethodPost, config.Config.ServiceURL+"/api/auth/login", http.NoBody)
	if err != nil {
		n.OpenModal(err.Error(), page.HomeMenu)
		return false
	}
	n.updateCachedCred(username, password)
	req.SetBasicAuth(username, password)
	res, err := n.client.Do(req)
	if err != nil {
		n.OpenModal(err.Error(), page.HomeMenu)
		return false
	}
	defer res.Body.Close()
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		n.OpenModal(err.Error(), page.HomeMenu)
		return false
	}

	var loginResponse LoginResponse
	err = json.Unmarshal(resBody, &loginResponse)
	if err != nil {
		n.OpenModal(fmt.Sprintf("error unmarshalling token\n\n%s - \n%s\n\n%s", res.Status, err, resBody), page.Login)
		return false
	}
	n.cookie = &http.Cookie{
		Name:   "memberserver",
		Value:  loginResponse.Token,
		MaxAge: 300,
	}
	// n.OpenModal(fmt.Sprintf("%s:%s:%s --- %t", username, password, loginResponse.Token, n.IsLoggedIn()), page.Login)
	// n.OpenModal(fmt.Sprintf("login response\n\n%s --- %t", res.Status, n.IsLoggedIn()), page.Login)

	store.MemberCache = n.GetMembers()

	return res.StatusCode == http.StatusOK
}

func (n *netService) GetUser() model.User {
	req, err := http.NewRequest(http.MethodGet, "/api/user", http.NoBody)
	if err != nil {
		n.OpenModal(err.Error(), page.HomeMenu)
	}
	req.SetBasicAuth(store.CurrentUsername, store.CurrentUserPass)
	// req.Header.Add("Authorization", "Bearer "+n.token)

	req.AddCookie(n.cookie)
	n.OpenModal(`Bearer `+n.cookie.Value, page.HomeMenu)
	res, err := n.client.Do(req)
	if err != nil {
		n.OpenModal(err.Error(), page.HomeMenu)
	}
	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		n.OpenModal(err.Error(), page.HomeMenu)
	}

	var user model.User
	json.Unmarshal(resBody, &user)

	n.OpenModal(fmt.Sprintf("get user response\n\n%s", res.Status), page.HomeMenu)
	return user
}

func (n *netService) GetMembers() []model.Member {
	req, err := http.NewRequest(http.MethodGet, config.Config.ServiceURL+"/api/member", http.NoBody)
	if err != nil {
		n.OpenModal(err.Error(), page.HomeMenu)
	}
	req.SetBasicAuth(store.CurrentUsername, store.CurrentUserPass)
	// req.Header.Add("Authorization", "Bearer "+n.token)
	req.AddCookie(n.cookie)

	res, err := n.client.Do(req)
	if err != nil {
		n.OpenModal(err.Error(), page.HomeMenu)
	}
	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		n.OpenModal(err.Error(), page.HomeMenu)
	}

	// n.OpenModal(string(resBody))

	var members []model.Member
	json.Unmarshal(resBody, &members)
	// n.OpenModal(fmt.Sprintf("get members response\n\n%s", res.Status))
	store.MemberCache = members
	return members
}

func (n *netService) UpdateMember(m model.Member) {
	defer n.GetMembers()

	// marshal member update to json
	json, err := json.Marshal(m)
	if err != nil {
		n.OpenModal(err.Error(), page.HomeMenu)
	}

	req, err := http.NewRequest(http.MethodPut, config.Config.ServiceURL+"/api/member/email/"+m.Email, bytes.NewBuffer(json))
	if err != nil {
		n.OpenModal(err.Error(), page.HomeMenu)
	}
	req.SetBasicAuth(store.CurrentUsername, store.CurrentUserPass)
	// req.Header.Add("Authorization", "Bearer "+n.token)
	req.AddCookie(n.cookie)

	res, err := n.client.Do(req)
	if err != nil {
		n.OpenModal(err.Error(), page.HomeMenu)
	}
	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		n.OpenModal(err.Error(), page.HomeMenu)
	}

	n.OpenModal(fmt.Sprintf("update member response\n\n%s\n\n%s", res.Status, resBody), page.MemberEditor)
}

func (n *netService) UpdateRFID(m model.Member) {
	defer n.GetMembers()
	// marshal member update to json
	json, err := json.Marshal(m)
	if err != nil {
		n.OpenModal(err.Error(), page.HomeMenu)
	}

	req, err := http.NewRequest(http.MethodPut, config.Config.ServiceURL+"/member/assignRFID"+m.Email, bytes.NewBuffer(json))
	if err != nil {
		n.OpenModal(err.Error(), page.HomeMenu)
	}
	req.SetBasicAuth(store.CurrentUsername, store.CurrentUserPass)
	// req.Header.Add("Authorization", `Bearer `+n.token)
	req.AddCookie(n.cookie)

	res, err := n.client.Do(req)
	if err != nil {
		n.OpenModal(err.Error(), page.HomeMenu)
	}
	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		n.OpenModal(err.Error(), page.HomeMenu)
	}

	n.OpenModal(fmt.Sprintf("update member response\n\n%s\n\n%s", res.Status, resBody), page.MemberEditor)
}
