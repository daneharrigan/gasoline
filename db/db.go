package db

import (
	"github.com/bmizerany/perks/topk"
	"sync"
)

var (
	r map[string]*Record
	l sync.RWMutex
)

type Record struct {
	//Hit int64
	PageView      int64
	Visit         int64
	UniqueVisitor int64
	ReturnVisitor int64
	TopK          *topk.Stream
}

func init() {
	r = make(map[string]*Record)
}

func New(id string) *Record {
	l.Lock()
	defer l.Unlock()

	r[id] = &Record{TopK: topk.New(10)}
	return r[id]
}

func Get(id string) *Record {
	l.Lock()
	defer l.Unlock()

	return r[id]
}
