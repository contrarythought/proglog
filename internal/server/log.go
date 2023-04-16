package server

import (
	"errors"
	"sync"
)

type Record struct {
	Value  []byte `json:"value"`
	Offset uint64 `json:"offset"`
}

type Log struct {
	mu      sync.RWMutex
	Records []Record
}

func NewLog() *Log {
	return &Log{}
}

func (l *Log) Append(record Record) (uint64, error) {
	l.mu.Lock()
	defer l.mu.Unlock()
	record.Offset = uint64(len(l.Records))
	l.Records = append(l.Records, record)
	return record.Offset, nil
}

var errOffset error = errors.New("offset not found")

func (l *Log) Read(offset uint64) (Record, error) {
	l.mu.RLock()
	defer l.mu.RUnlock()
	if offset >= uint64(len(l.Records)) {
		return Record{}, errOffset
	}
	record := l.Records[offset]
	return record, nil
}
