package rest_api

import (
	"testing"
)

func TestFunc(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name string
	}{
		{
			name: "test",
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			t.Log("test")
		})
	}
}
