package sum

import "sort"

type Element struct {
	Value string
	Count int
}

type Result []*Element

type Stream struct {
	r map[string]*Element
}

func New() *Stream {
	return &Stream{r: make(map[string]*Element)}
}

func (s *Stream) Insert(x string) {
	if e := s.r[x]; e != nil {
		e.Count++
	} else {
		s.r[x] = &Element{Value: x, Count: 1}
	}
}

func (s *Stream) Query() Result {
	var r Result
	for _, e := range s.r {
		r = append(r, s.r)
	}

	sort.Sort(sort.Reverse(r))
	return r
}

func (r Result) Len() int {
	return len(r)
}

func (r Result) Less(a, b int) bool {
	return r[a].Count < r[b].Count
}

func (r Result) Swap(a, b int) {
	r[a], r[b] = r[b], r[a]
}
