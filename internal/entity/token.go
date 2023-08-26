package entity

type Token struct {
	AccessToken string
	RefreshToken string `json:"refresh_token"`
}