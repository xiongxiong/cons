package cons

import "sync"

type signal struct{}

type done struct {
	c chan signal
}

type wait done

// Cons ...
type Cons struct {
	sync.Mutex
	theMap map[interface{}]*wait
}

func (cs *Cons) check(key interface{}) Con {
	cs.Lock()
	defer cs.Unlock()

	reset := func(d done, w wait) {
		<-d.c
		close(d.c)
		close(w.c)

		cs.Lock()
		defer cs.Unlock()
		delete(cs.theMap, key)
	}

	if cs.theMap[key] == nil {
		d := done{make(chan signal)}
		w := wait{make(chan signal)}
		cs.theMap[key] = &w
		go reset(d, w)
		return Con{con: &d}
	}
	return Con{con: cs.theMap[key]}
}

// GetCons ...
func GetCons() *Cons {
	return &Cons{
		theMap: make(map[interface{}]*wait),
	}
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

// Done ...
func (c Con) Done() {
	switch s := c.con.(type) {
	case *done:
		s.c <- signal{}
	}
}
