package excerpt

import (
	"strings"
	// "fmt"
	"unicode"
)

type highlightFunc func(word string) string

func sentencesStop(w rune) (r bool) {
	r = (strings.Index(",，.。?？!！#\n", string(w)) >= 0)
	return
}

type sentenceScanner struct {
	src     []rune
	length  int
	start   int
	pos     int
	stopped bool
}

func newScanner(src string) (r *sentenceScanner) {
	r = new(sentenceScanner)
	r.src = []rune(src)
	r.length = len(r.src)
	r.pos = -1
	return
}

func (ss *sentenceScanner) next() (r string, finished bool) {
	for {
		ss.pos = ss.pos + 1

		if ss.pos > ss.length {
			finished = true
			return
		}

		if ss.pos == ss.length {
			r = string(ss.src[ss.start:ss.pos])
			return
		}
		// fmt.Println("pos: ", string(ss.src))
		// fmt.Println("pos: ", ss.pos)
		// fmt.Println("length: ", ss.length)
		current := ss.src[ss.pos]
		if ss.stopped && !unicode.IsSpace(current) {
			r = string(ss.src[ss.start:ss.pos])
			ss.start = ss.pos
			ss.stopped = false
			return
		}

		if sentencesStop(current) {
			ss.stopped = true
		}

	}
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
	r = strings.Join(rs, "")
	return
}

// Find the highlights sentences in each chunk of sources, and highlight the sentence and find the previous and next sentence around it, all connected highlighted sentences are merged into one
func SentencesAround(sources []string, keywords []string, hf highlightFunc) (r []string) {
	for _, chunk := range sources {
		ss := newScanner(chunk)

		rc := resultChunk{}

		for {
			s, finished := ss.next()
			if finished {
				break
			}

			hs := s

			if rc.after != "" {
				r = append(r, rc.tostring())
				rc = resultChunk{}
			}

			highlighted := false
			for _, keyword := range keywords {
				var yes bool
				hs, yes = Highlight(hs, keyword, hf)
				if yes {
					highlighted = true
				}
			}

			// is highlighted or has highlighted before, but the sentence is shorter then 16 charactor.
			if highlighted || (len(rc.lines) > 0 && len(hs) < 16) {
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
