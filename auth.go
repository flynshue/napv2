package napv2

type Authentication interface {
	AuthorizationHeader() string
}

type AuthBasic struct {
	Username string
	Password string
}

func NewAuthBasic() AuthBasic {
	return AuthBasic{}
}

func (ab AuthBasic) AuthorizationHeader() string {
	return ""
}

type AuthToken struct {
	Token string
}

func (at AuthToken) AuthorizationToken() string {
	return ""
}
