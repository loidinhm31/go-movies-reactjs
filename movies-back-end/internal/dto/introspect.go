package dto

type Introspect struct {
	ClientId     string
	ClientSecret string
	Token        string `form:"token"`
}
