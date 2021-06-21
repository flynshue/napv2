package napv2

import "testing"

func TestResource_RenderEndpoint(t *testing.T) {
	resource := NewResource("GET", "/anything/{{.user}}", NewRouter())
	got := resource.RenderEndpoint(map[string]string{"user": "walker"})
	want := "/anything/walker"
	if got != want {
		t.Fatalf("error RenderEndpoint(); got %s, wanted %s\n", got, want)
	}
}
