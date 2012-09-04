package tests

import (
	"testing"
	"github.com/sunfmin/excerpt"
)

type Case struct {
	Source   []string
	Keywords []string
	Result   []string
}

var p1 = `We currently have only one sign at the outside of the building. I was just wondering where do you plan to put the sign in China. Are you going to put it at the entrance gate, or somewhere at the entrance of the building?

Are you still going to use the cutout logo somewhere near the entrance of the office upstairs?`

var cases = []Case{
	{
		[]string{p1},
		[]string{"china", "sign"},
		[]string{`We currently have only one *sign* at the outside of the building`, `I was just wondering where do you plan to put the *sign* in *China*`},
	},
	{
		[]string{p1},
		[]string{"entrance"},
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

func TestSentencesAround(t *testing.T) {

	for _, c := range cases {
		r := excerpt.SentencesAround(c.Source, c.Keywords, highlight)
		if len(r) != len(c.Result) {
			t.Error(r)
		}
		for i, line := range r {
			if c.Result[i] != line {
				t.Errorf("expected: \n%s, \n\nbut was: \n%s", c.Result[i], line)
			}
		}
	}

}
