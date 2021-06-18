package napv2

type API struct {
	BaseURL       string
	Client        *Client
	Resources     []RestResource
	DefaultRouter *CBRouter
}
