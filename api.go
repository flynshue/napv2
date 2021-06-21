package napv2

import "fmt"

type API struct {
	BaseURL       string
	Client        *Client
	Resources     map[string]RestResource
	DefaultRouter *CBRouter
}

func NewAPI(baseURL string) *API {
	return &API{
		BaseURL:       baseURL,
		Client:        NewClient(),
		Resources:     make(map[string]RestResource),
		DefaultRouter: NewRouter(),
	}
}

func (a *API) AddResource(name string, resource RestResource) {
	a.Resources[name] = resource
}

func (a *API) ListResources() []string {
	resources := make([]string, 0, len(a.Resources))
	for name := range a.Resources {
		resources = append(resources, name)
	}
	return resources
}

func (a *API) SetAuth(auth Authentication) {
	a.Client.SetAuth(auth)
}

func (a *API) Call(name string, params map[string]string, payload interface{}) error {
	res, ok := a.Resources[name]
	if !ok {
		return fmt.Errorf("%s resource not found", name)
	}
	return a.Client.ProcessRequest(a.BaseURL, res, params, payload)
}
