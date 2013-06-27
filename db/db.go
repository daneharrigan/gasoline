package db

import (
	"gasoline/sum"
	"github.com/daneharrigan/perks/topk"
	"sync"
)

var (
	r  map[string]*Record
	l  sync.RWMutex
	k  = 10
	fs = []string{"Cookies", "QuickTime", "Shockwave Flash", "Google Talk",
		"Java Applet", "Silverlight", "Retina Display"}
)

type Record struct {
	PageView      int64
	Visit         int64
	UniqueVisitor int64
	ReturnVisitor int64
	TopK          *topk.Stream
	Features      *sum.Stream
	Resolutions   *sum.Stream
	OS            *sum.Stream
}

func init() {
	r = make(map[string]*Record)
}

func New(id string) *Record {
	l.Lock()
	defer l.Unlock()

	r[id] = &Record{
		TopK:        topk.New(k),
		Features:    sum.New(fs...),
		Resolutions: sum.New(),
		OS:          sum.New(),
	}

	return r[id]
}

func Get(id string) *Record {
	l.Lock()
	defer l.Unlock()

	return r[id]
}
