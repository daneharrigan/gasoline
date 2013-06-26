package db

import (
	"gasoline/sum"
	"github.com/daneharrigan/perks/topk"
	"sync"
)

var (
	r map[string]*Record
	l sync.RWMutex
	k = 10
)

type Record struct {
	//Hit int64
	PageView      int64
	Visit         int64
	UniqueVisitor int64
	ReturnVisitor int64
	TopK          *topk.Stream
	Features      *sum.Stream
}

func init() {
	r = make(map[string]*Record)
}

func New(id string) *Record {
	l.Lock()
	defer l.Unlock()

	r[id] = &Record{TopK: topk.New(k), Features: sum.New()}
	return r[id]
}

func Get(id string) *Record {
	l.Lock()
	defer l.Unlock()

	return r[id]
}

func (rec *Record) Flush() {
	l.Lock()
	defer l.Unlock()

	rec.PageView = 0
	rec.Visit = 0
	rec.UniqueVisitor = 0
	rec.ReturnVisitor = 0
	rec.TopK = topk.New(k)
}
