package excerpt

import (
	"github.com/sunfmin/exphtml"
	"strings"
	"io"
	"bytes"
)

func HighlightHtml(source string, keywords []string, hf highlightFunc) (r string, highlighted bool, err error) {
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

