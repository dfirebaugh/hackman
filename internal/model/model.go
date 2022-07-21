package model

type User struct {
	Email string
	Name  string
}
type Member struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	// FullName?
	FullName    string `json:"fullName"`
	Email       string `json:"email"`
	RFID        string `json:"rfid"`
	MemberLevel uint   `json:"memberLevel"`
	// Resources []Resource
	SubscriptionID string `json:"subscriptionID"`
}

type MemberLevel uint

const (
	// Inactive $0
	Inactive MemberLevel = iota + 1
	// Credited $1
	Credited
	// Classic $30
	Classic
	// Standard $35
	Standard
	// Premium $50
	Premium
)

// MemberLevelToStr convert MemberLevel to string
var MemberLevelToStr = map[MemberLevel]string{
	Inactive: "Inactive",
	Credited: "Credited",
	Classic:  "Classic",
	Standard: "Standard",
	Premium:  "Premium",
}
