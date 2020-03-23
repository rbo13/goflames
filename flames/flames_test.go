package flames_test

import (
	"testing"

	"github.com/rbo13/goflames/flames"
)

var nameTests = []struct {
	title     string
	names     []string
	out       string
	flamesOut string
}{
	{
		title:     "Should Be Married",
		names:     []string{"jepjeply", "beverly"},
		out:       "ichadbevely",
		flamesOut: "Married",
	},
}

func TestFlames(t *testing.T) {
	for _, tt := range nameTests {
		t.Run(tt.title, func(t *testing.T) {
			res := flames.Pair(tt.names[0], tt.names[1])

			if res != tt.out {
				t.Errorf("got %q, want %q", res, tt.out)
			}

			val := flames.Get(res)

			if val != tt.flamesOut {
				t.Errorf("got %q, want %q", val, tt.flamesOut)
			}
		})
	}
}
