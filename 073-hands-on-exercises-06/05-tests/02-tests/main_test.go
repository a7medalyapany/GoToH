package main

import "testing"

func TestGet(t *testing.T) {
	md := &MockDatastore{
		Users: map[int]User{
			2: {ID: 2, First: "Ahmed"},
		},
	}

	s := &Service{
		ds: md,
	}

	u, err := s.GetUser(2)
	if err != nil {
		t.Errorf("Got Error: %v", err)
	}

	if u.First != "Ahmed" {
		t.Errorf("Got: %s; Want: %s", u.First, "Ahmed")
	}
}
