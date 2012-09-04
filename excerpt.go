package excerpt

import (
	"strings"
	// "fmt"
)

type highlightFunc func(word string) string

func sentencesStop(w rune) (r bool) {
	r = (strings.Index(",，.。?？!！\n", string(w)) >= 0)
	return
}

func SentencesAround(sources []string, keywords []string, hf highlightFunc) (r []string) {
	for _, chunk := range sources {
		sentences := strings.FieldsFunc(chunk, sentencesStop)
		previousHighlighted := false

		for _, s := range sentences {
			hs := strings.TrimSpace(s)
			highlighted := false
			for _, keyword := range keywords {
				var yes bool
				hs, yes = Highlight(hs, keyword, hf)
				if yes {
					highlighted = true
				}
			}
			if highlighted {
				r = append(r, hs)
			}
		}
	}
	return
}

func Highlight(sentence, keyword string, hf highlightFunc) (r string, highlighted bool) {
	downcaseSentence := strings.ToLower(sentence)
	left := downcaseSentence
	sentenceLeft := sentence
	for {
		i := strings.Index(left, keyword)
		if i < 0 {
			break
		}
		wordTo := (i + len(keyword))

		r = r + sentenceLeft[:i] + hf(sentenceLeft[i:wordTo])
		left = left[wordTo:]

		sentenceLeft = sentenceLeft[wordTo:]
		highlighted = true
	}
	r = r + sentenceLeft
	return
}
