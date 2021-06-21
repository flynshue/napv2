package napv2

import "testing"

func TestAuthBasic(t *testing.T) {
	auth := NewAuthBasic("testUser", "pass123")
	want := "Basic dGVzdFVzZXI6cGFzczEyMw=="
	got := auth.AuthorizationHeader()
	if got != want {
		t.Fatalf("error AuthorizationHeader(); got %s, want %s\n", got, want)
	}
}

func TestAuthToken(t *testing.T) {
	auth := AuthToken{Token: "56451208-d294-11eb-8bd0-c8f75077546c"}
	want := "token 56451208-d294-11eb-8bd0-c8f75077546c"
	got := auth.AuthorizationHeader()
	if got != want {
		t.Fatalf("error AuthorizationHeader(); got %s, want %s\n", got, want)
	}
}
