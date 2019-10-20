package models

const (
	KEY_TOKEN = "token"
)

type Reduce struct {
	Id    int64  `json:"-"`
	Url   string `json:"url"`
	Token string `json:"token"`
}
