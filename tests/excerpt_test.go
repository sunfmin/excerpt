package tests

import (
	"testing"
	"github.com/sunfmin/excerpt"
)

type Case struct {
	Source          []string
	Keywords        string
	AroundWordCount int
	Result          []string
}

var p1 = `We currently have only one sign at the outside of the building. I was just wondering where do you plan to put the sign in China. Are you going to put it at the entrance gate, or somewhere at the entrance of the building?

Are you still going to use the cutout logo somewhere near the entrance of the office upstairs?`

var cases = []Case{
	{
		[]string{p1},
		"china sign",
		10,
		[]string{`was just wondering where do you plan to put the *sign* in *China*. Are you going to put it at the entrance gate`},
	},
	{
		[]string{p1},
		"entrance",
		10,
		[]string{
			`Are you going to put it at the *entrance* gate, or somewhere at the *entrance* of the building?`,
			`still going to use the cutout logo somewhere near the entrance of the office upstairs?`,
		},
	},
}

func highlight(word string) (r string) {
	r = "*" + word + "*"
	return
}

func TestExcerpt(t *testing.T) {

	for _, c := range cases {
		r := excerpt.ExcerptsAround(c.Source, c.Keywords, c.AroundWordCount, highlight)
		if len(r) != len(c.Result) {
			t.Error(r)
		}
		for i, line := range r {
			if c.Result[i] != line {
				t.Error(line)
			}
		}
	}

}
