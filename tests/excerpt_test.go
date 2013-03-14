package tests

import (
	"github.com/sunfmin/excerpt"
	"testing"
)

type Case struct {
	Source        []string
	Keywords      []string
	Result        []string
	HighlightFunc func(word string) (r string)
}

var p1 = `We currently have only one sign at the outside of the building. I was just wondering where do you plan to put the sign in China. Are you going to put it at the entrance gate, or somewhere at the entrance of the building?

Are you still going to use the cutout logo somewhere near the entrance of the office upstairs?`

var cases = []Case{
	{
		[]string{p1},
		[]string{"entrance"},
		[]string{
			`I was just wondering where do you plan to put the sign in China. Are you going to put it at the *entrance* gate, or somewhere at the *entrance* of the building?

Are you still going to use the cutout logo somewhere near the *entrance* of the office upstairs?`,
		},
		highlight,
	},
	{
		[]string{p1},
		[]string{"china", "sign"},
		[]string{`We currently have only one *sign* at the outside of the building. I was just wondering where do you plan to put the *sign* in *China*. Are you going to put it at the entrance gate, `},
		highlight,
	},
	{
		[]string{p1},
		[]string{"outside"},
		[]string{`We currently have only one sign at the *outside* of the building. I was just wondering where do you plan to put the sign in China. `},
		highlight,
	},
	{
		[]string{p1},
		[]string{"upstairs"},
		[]string{`or somewhere at the entrance of the building?

Are you still going to use the cutout logo somewhere near the entrance of the office *upstairs*?`},
		highlight,
	},
	{
		[]string{`杭州大浪软件技术有限公司，位于杭州西湖区`},
		[]string{"大浪", "技术"},
		[]string{`杭州*大浪*软件*技术*有限公司，位于杭州西湖区`},
		highlight,
	},
	{
		[]string{`杭州大浪软件技术有限公司，位于杭州西湖区`},
		[]string{"大浪", "大浪软件技术", "技术有限"},
		[]string{`杭州*大浪软件技术有限*公司，位于杭州西湖区`},
		highlight,
	},
	{
		[]string{`杭州大浪软件技术有限公司，位于杭州西湖区`},
		[]string{"大浪", "软件", "术"},
		[]string{`杭州*大浪软件*技*术*有限公司，位于杭州西湖区`},
		highlight,
	},
	{
		[]string{`fsfds found a dsfsdfsd`},
		[]string{"found", "a"},
		[]string{`fsfds <span style="color: #ff6226">found</span> <span style="color: #ff6226">a</span> dsfsdfsd`},
		highlightspan,
	},
}

func highlight(word string) (r string) {
	r = "*" + word + "*"
	return
}

func highlightspan(word string) (r string) {
	r = `<span style="color: #ff6226">` + word + `</span>`
	return
}

func TestSentencesAround(t *testing.T) {

	for _, c := range cases {
		r := excerpt.SentencesAround(c.Source, c.Keywords, c.HighlightFunc)
		if len(r) != len(c.Result) {
			t.Errorf("expected: \n%s \n\nbut was: \n%s\n\n\n", c.Result, r)
			continue
		}
		for i, line := range r {
			if c.Result[i] != line {
				t.Errorf("expected: \n%s \n\nbut was: \n%s\n\n\n", c.Result[i], line)
			}
		}
	}

}
