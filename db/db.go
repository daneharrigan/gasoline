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
		"Java Applet", "Silverlight", "Retina Display", "Geolocation",
		"Touch Screen", "Tilt"}
	ps = []float64{0.5, 0.9, 0.99}
)

type Record struct {
	l             sync.RWMutex
	PageView      int64
	Visit         int64
	UniqueVisitor int64
	ReturnVisitor int64
	TopK          *topk.Stream
	Language      *sum.Stream
	Features      *sum.Stream
	Resolutions   *sum.Stream
	OS            *sum.Stream
	ViewDurations map[string]*quantile.Stream
	BrowserUsage  map[string]*sum.Stream
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
		Language:      sum.New(),
		OS:            sum.New(),
		ViewDurations: make(map[string]*quantile.Stream),
		BrowserUsage:  make(map[string]*sum.Stream),
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
	var ok bool

	if s, ok = r.ViewDurations[u]; !ok {
		s = quantile.NewTargeted(ps...)
		r.ViewDurations[u] = s
	}

	s.Insert(d)
}

func (r *Record) BrowserUsed(n, v string) {
	r.l.Lock()
	defer r.l.Unlock()

	var s *sum.Stream
	var ok bool

	if s, ok = r.BrowserUsage[n]; !ok {
		s = sum.New()
		r.BrowserUsage[n] = s
	}

	s.Insert(v)
}
