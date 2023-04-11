package uuidv7

import (
	"encoding/binary"
	"math/rand"
	"sync"
	"time"
)

type Generator struct {
	ts      int64
	counter uint32
	mu      sync.Mutex
	rnd     rand.Source
}

func New() *Generator {
	ts := time.Now().UnixMilli()
	return &Generator{
		ts:  ts,
		rnd: rand.NewSource(ts),
	}
}

func (u *Generator) Next() UUID {
	ts := time.Now().UnixMilli()

	u.mu.Lock()

	// Implementation does not make explicit guarantees about clock drift.
	// The counter may roll over if there is significantly negative clock corrections
	if u.ts < ts {
		u.ts = ts
		u.counter = 0
	} else {
		u.counter += 1
	}

	cnt := u.counter
	rnd1 := uint64(u.rnd.Int63())
	rnd2 := uint64(u.rnd.Int63())

	u.mu.Unlock()

	var val [16]byte

	binary.LittleEndian.PutUint64(val[0:8], (2<<62)|((uint64(cnt)&0xFFF)<<50)|(rnd1&0xFFFFFFFFFFFFF))
	binary.LittleEndian.PutUint64(val[8:16], (uint64(ts)<<16)+(7<<12)+rnd2&0xFFF)

	return val
}
