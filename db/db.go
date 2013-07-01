package db

import (
	"gasoline/sum"
	"github.com/daneharrigan/perks/quantile"
	"github.com/daneharrigan/perks/topk"
	"sync"
)

var (
	r  map[string]*Record
	l  sync.RWMutex
	k  = 10
	fs = []string{"Cookies", "QuickTime", "Shockwave Flash", "Google Talk",
		"Java Applet", "Silverlight", "Retina Display"}
	ps = []float64{0.5, 0.9, 0.99}
)

type Record struct {
	l             sync.RWMutex
	PageView      int64
	Visit         int64
	UniqueVisitor int64
	ReturnVisitor int64
	TopK          *topk.Stream
	Features      *sum.Stream
	Resolutions   *sum.Stream
	OS            *sum.Stream
	ViewDurations map[string]*quantile.Stream
}

func init() {
	r = make(map[string]*Record)
}

func New(id string) *Record {
	l.Lock()
	defer l.Unlock()

	r[id] = &Record{
		TopK:          topk.New(k),
		Features:      sum.New(fs...),
		Resolutions:   sum.New(),
		OS:            sum.New(),
		ViewDurations: make(map[string]*quantile.Stream),
	}

	return r[id]
}

func Get(id string) *Record {
	l.Lock()
	defer l.Unlock()

	return r[id]
}

func (r *Record) ViewDuration(u string, d float64) {
	r.l.Lock()
	defer r.l.Unlock()

	var s *quantile.Stream
	if s, ok := r.ViewDurations[u]; !ok {
		s = quantile.NewTargeted(ps...)
		r.ViewDurations[u] = s
	}

	s.Insert(d)
}
