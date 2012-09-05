package excerpt

import (
	"strings"
	// "fmt"
)

type highlightFunc func(word string) string

func sentencesStop(w rune) (r bool) {
	r = (strings.Index(",，.。?？!！#\n", string(w)) >= 0)
	return
}

type resultChunk struct {
	before string
	lines  []string
	after  string
}

func (rc resultChunk) tostring() (r string) {
	var rs []string
	if rc.before != "" {
		rs = append(rs, rc.before)
	}
	rs = append(rs, rc.lines...)
	if rc.after != "" {
		rs = append(rs, rc.after)
	}
	r = strings.Join(rs, ", ")
	return
}

// Find the highlights sentences in each chunk of sources, and highlight the sentence and find the previous and next sentence around it, all connected highlighted sentences are merged into one
func SentencesAround(sources []string, keywords []string, hf highlightFunc) (r []string) {
	for _, chunk := range sources {
		sentences := strings.FieldsFunc(chunk, sentencesStop)

		rc := resultChunk{}

		for _, s := range sentences {
			if rc.after != "" {
				r = append(r, rc.tostring())
				rc = resultChunk{}
			}

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
				rc.lines = append(rc.lines, hs)
			} else {
				if len(rc.lines) > 0 {
					rc.after = hs
				} else {
					rc.before = hs
				}
			}
		}
		if len(rc.lines) > 0 {
			r = append(r, rc.tostring())
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
