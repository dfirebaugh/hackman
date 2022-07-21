package page

type Page uint

const (
	Login = iota
	LoginError
	Loading
	HomeMenu
	MemberList
	MemberEditor
	CurrentForm
	AddMember
	Search
)
