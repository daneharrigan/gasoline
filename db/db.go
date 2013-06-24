package db

import "sync"

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
}

func init() {
	r = make(map[string]*Record)
}

func New(id string) *Record {
	l.Lock()
	defer l.Unlock()

	r[id] = new(Record)
	return r[id]
}

func Get(id string) *Record {
	l.Lock()
	defer l.Unlock()

	return r[id]
}
