package ping_test

import (
	"ping"
	"testing"
)

func TestSend(t *testing.T) {
	cases := []struct {
		want string
	}{
		{want: "pong"},
	}
	for _, c := range cases {
		result := ping.Send()
		if result != c.want {
			t.Fatalf("[%s] is incorrect, we want [%s]", result, c.want)
		}
	}
}