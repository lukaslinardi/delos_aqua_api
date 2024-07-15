package auth

type LoginResponse struct {
	Token    TokenRes `json:"token"`
	Fullname string   `json:"fullname" db:"fullname"`
	Nik      string   `json:"nik" db:"nik"`
	Password string   `json:"-" db:"password"`
	ID       int64    `json:"-" db:"id"`
}

type TokenRes struct {
	Access        string `json:"access"`
	AccessExpired string `json:"access_expired"`
	Renew         string `json:"renew"`
	RenewExpired  string `json:"renew_expired"`
}

type CredentialData struct {
	ID       int64  `json:"id"`
	Fullname string `json:"fullname"`
	Nik      string `json:"nik"`
}
