package napv2

import (
	"encoding/base64"
	"fmt"
)

type Authentication interface {
	AuthorizationHeader() string
}

type AuthBasic struct {
	Username string
	Password string
}

func NewAuthBasic(user, pass string) AuthBasic {
	return AuthBasic{
		Username: user,
		Password: pass,
	}
}

func (ab AuthBasic) AuthorizationHeader() string {
	creds := fmt.Sprintf("%s:%s", ab.Username, ab.Password)
	b64Creds := base64.URLEncoding.EncodeToString([]byte(creds))
	return fmt.Sprintf("Basic %s", b64Creds)
}

type AuthToken struct {
	Token string
}

func (at AuthToken) AuthorizationToken() string {
	return fmt.Sprintf("token %s", at.Token)
}
