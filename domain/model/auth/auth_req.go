package auth

import (
	"time"
)

// type SignUp struct {
// 	Fullname string `json:"fullname"`
// 	Email    string `json:"email"`
// 	Password string `json:"password"`
// }

type SignUp struct {
	Nik        string    `json:"nik"`
	Fullname   string    `json:"fullname"`
	Legalname  string    `json:"legal_name"`
	BirthPlace string    `json:"birth_place"`
	BirthDate  time.Time `json:"birth_date"`
	Salary     int64     `json:"salary"`
	KtpPics    string    `json:"ktp_pics"`
	Pics       string    `json:"pics"`
	Password   string    `json:"password"`
}

type Login struct {
	Fullname string `json:"fullname"`
	Nik      string `json:"nik"`
	Password string `json:"password"`
}

func (r *SignUp) Validate() map[string]string {
	if r.Fullname == "" {
		return map[string]string{
			"en": "fullname not found",
			"id": "fullname tidak boleh kosong",
		}
	}

	if r.Password == "" {
		return map[string]string{
			"en": "password not found",
			"id": "kata sandi tidak boleh kosong",
		}
	}
	return nil
}

func (r *Login) Validate() map[string]string {
	if r.Fullname == "" {
		return map[string]string{
			"en": "fullname not found",
			"id": "fullname tidak boleh kosong",
		}
	}

	if r.Password == "" {
		return map[string]string{
			"en": "password not found",
			"id": "kata sandi tidak boleh kosong",
		}
	}
	return nil
}
