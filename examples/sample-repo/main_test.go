package main

import "testing"

func TestUser(t *testing.T) {
	u := User{ID: 1, Name: "Alice"}
	if u.Name != "Alice" {
		t.Errorf("expected Alice, got %s", u.Name)
	}
}
