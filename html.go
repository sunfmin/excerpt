package excerpt

import (
	"bytes"
	"github.com/sunfmin/exphtml"
	"io"
	"strings"
)

func HighlightHtml(source string, keywords []string, hf highlightFunc) (r string, highlighted bool, err error) {
	r = source
	if len(keywords) == 0 {
		return
	}
	z := exphtml.NewTokenizer(strings.NewReader(source))
	buf := bytes.NewBuffer(nil)
	for {
		tt := z.Next()

		switch tt {
		case exphtml.ErrorToken:
			if z.Err() != io.EOF {
				err = z.Err()
				return
			}
			goto exit
		case exphtml.TextToken:
			htext, _ := Highlight(string(z.Text()), keywords, hf)
			buf.WriteString(htext)
		case exphtml.StartTagToken, exphtml.EndTagToken:
			buf.Write(z.Raw())
		}
	}
exit:
	r = buf.String()
	return
}
