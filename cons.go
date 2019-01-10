package cons

import "sync"

type signal struct{}

type done struct {
	c chan signal
}

type wait done

type queue struct {
	c     chan signal
	cc    chan signal
	count int
}

// Cons ...
type Cons struct {
	sync.Mutex
	mapWait  map[interface{}]*wait
	mapQueue map[interface{}]*queue
}

func (cs *Cons) check(key interface{}) Con {
	cs.Lock()
	defer cs.Unlock()

	reset := func(d *done, w *wait) {
		<-d.c
		close(d.c)
		close(w.c)

		cs.Lock()
		defer cs.Unlock()
		delete(cs.mapWait, key)
	}

	if cs.mapWait[key] == nil {
		d := done{make(chan signal)}
		w := wait{make(chan signal)}
		cs.mapWait[key] = &w
		go reset(&d, &w)
		return Con{con: &d}
	}
	return Con{con: cs.mapWait[key]}
}

func (cs *Cons) queue(key interface{}) Con {
	cs.Lock()
	defer cs.Unlock()

	reset := func(q *queue) {
		for {
			<-q.cc

			cs.Lock()
			q.count--
			if q.count == 0 {
				close(q.c)
				close(q.cc)
				delete(cs.mapQueue, key)
				break
			} else {
				q.c <- signal{}
			}
			cs.Unlock()
		}
	}

	if cs.mapQueue[key] == nil {
		q := queue{make(chan signal, 1), make(chan signal), 0}
		q.c <- signal{}
		cs.mapQueue[key] = &q
		go reset(&q)
	}
	q := cs.mapQueue[key]
	q.count++
	return Con{con: q}
}

// GetCons ...
func GetCons() *Cons {
	return &Cons{
		mapWait:  make(map[interface{}]*wait),
		mapQueue: make(map[interface{}]*queue),
	}
}

var cons = GetCons()

// Wait ...
func Wait(key interface{}) Con {
	return cons.Wait(key)
}

// Skip ...
func Skip(key interface{}) Con {
	return cons.Skip(key)
}

// Queue ...
func Queue(key interface{}) Con {
	return cons.Queue(key)
}

// Con ...
type Con struct {
	Skip bool
	con  interface{}
}

// Wait ...
func (cs *Cons) Wait(key interface{}) Con {
	c := cs.check(key)
	switch s := c.con.(type) {
	case *wait:
		<-s.c
	}
	return c
}

// Skip ...
func (cs *Cons) Skip(key interface{}) Con {
	c := cs.check(key)
	switch c.con.(type) {
	case *wait:
		c.Skip = true
	}
	return c
}

// Queue ...
func (cs *Cons) Queue(key interface{}) Con {
	c := cs.queue(key)
	switch s := c.con.(type) {
	case *queue:
		<-s.c
	}
	return c
}

// Done ...
func (c Con) Done() {
	switch s := c.con.(type) {
	case *done:
		s.c <- signal{}
	case *queue:
		s.cc <- signal{}
	}
}
