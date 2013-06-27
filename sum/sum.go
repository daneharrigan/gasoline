package sum

import "sort"

type Element struct {
	Value string
	Count int
}

type Results []*Element

type Stream struct {
	r map[string]*Element
	defined bool
}

func New(keys ...string) *Stream {
	r := make(map[string]*Element)
	s := new(Stream)

	if len(keys) > 0 {
		s.defined = true
		for _, k := range keys {
			r[k] = &Element{Value: k, Count: 0}
		}
	}

	s.r = r
	return s
}

func (s *Stream) Insert(x string) {
	if e, ok := s.r[x]; ok {
		e.Count++
	} else if !s.defined {
		s.r[x] = &Element{ Value: x, Count: 1 }
	}
}

func (s *Stream) Query() Results {
	var r Results
	for _, e := range s.r {
		r = append(r, e)
	}

	sort.Sort(sort.Reverse(r))
	return r
}

func (r Results) Len() int {
	return len(r)
}

func (r Results) Less(a, b int) bool {
	return r[a].Count < r[b].Count
}

func (r Results) Swap(a, b int) {
	r[a], r[b] = r[b], r[a]
}
