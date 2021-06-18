package napv2

import "net/http"

type Client struct {
	Client *http.Client
	Auth   Authentication
}
