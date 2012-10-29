package excerpt

import (
	"github.com/sunfmin/exphtml"
	"strings"
	"fmt"
	"io"
)

func HighlightHtml(source string, keywords []string, hf highlightFunc) (r string, highlighted bool) {
	z := exphtml.NewTokenizer(strings.NewReader(source))
	for {
		tt := z.Next()

		fmt.Println("RAW: ", string(z.Raw()))

		switch tt {
		case exphtml.ErrorToken:
			if z.Err() != io.EOF {
				fmt.Println(z.Err())
			}
			goto exit
		case exphtml.TextToken:
			fmt.Println("TEXT: ", string(z.Text()))
		case exphtml.StartTagToken, exphtml.EndTagToken:
			tn, _ := z.TagName()
			fmt.Println(string(tn))
		}
	}
exit:
	return
}

