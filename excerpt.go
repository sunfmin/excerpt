package excerpt

import (
	"fmt"
	"strings"
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
			var highlighted bool
			hs, highlighted = Highlight(hs, keywords, hf)

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

func Highlight(source string, keywords []string, hf highlightFunc) (r string, highlighted bool) {
	sol := &segOrderList{}
	for _, keyword := range keywords {
		findMatchesPut(source, keyword, sol)
	}

	if len(sol.l) > 0 {
		highlighted = true
	}
	last := 0
	for _, s := range sol.l {
		r = r + source[last:s.start]
		r = r + hf(source[s.start:s.end])
		last = s.end
	}
	r = r + source[last:]

	return
}

type seg struct {
	start int
	end   int
}

func (x *seg) String() string {
	return fmt.Sprintf("[%d, %d]", x.start, x.end)
}

func (x *seg) intersect(y *seg) bool {
	if x.start > y.end {
		return false
	}
	if x.end < y.start {
		return false
	}
	return true
}

func (x *seg) smallerThan(y *seg) bool {
	return x.end < y.start
}

func (x *seg) merge(y *seg) {
	if x.start > y.start {
		x.start = y.start
	}
	if x.end < y.end {
		x.end = y.end
	}
	return
}

type segOrderList struct {
	l []*seg
}

func (sol *segOrderList) putInOrder(s *seg) {
	p := -1
	for i, s1 := range sol.l {
		if s1.intersect(s) {
			s1.merge(s)
			return
		}
		if s.smallerThan(s1) {
			p = i
			break
		}
	}
	if p == -1 {
		sol.l = append(sol.l, s)
	} else {
		sol.l = append(sol.l[:p], append([]*seg{s}, sol.l[p:]...)...)
	}
	return
}

func findMatchesPut(source, keyword string, sol *segOrderList) {
	downcaseSentence := strings.ToLower(source)
	left := downcaseSentence
	var offset int
	for {
		i := strings.Index(left, keyword)
		if i < 0 {
			break
		}
		start := offset + i
		end := (start + len(keyword))

		s := &seg{start, end}
		sol.putInOrder(s)

		left = left[end:]
		offset += end
	}
	return
}
