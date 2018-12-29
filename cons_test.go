package cons_test

import (
	"cons"
	"sync"
	"testing"
	"time"
)

func TestWait(t *testing.T) {
	con := cons.GetCons()
	wg := sync.WaitGroup{}

	x := 0
	wait := func() {
		defer wg.Done()

		c := con.Wait("hello")
		defer c.Done()

		if x == 0 {
			t.Logf("--- sleeping")
			time.Sleep(1 * time.Second)
			x++
			return
		}
		t.Log("awake")
	}

	for range [100]int{} {
		wg.Add(1)
		go wait()
	}
	wg.Wait()

	t.Logf("x >>> %d", x)
}

func TestSkip(t *testing.T) {
	con := cons.GetCons()
	wg := sync.WaitGroup{}

	x := 0
	wait := func() {
		defer wg.Done()

		c := con.Skip("hello")
		defer c.Done()

		if c.Skip {
			t.Log("skip")
			return
		}
		t.Log("--- doing")
		time.Sleep(1 * time.Second)
		x++
	}

	for range [100]int{} {
		wg.Add(1)
		go wait()
	}
	wg.Wait()

	t.Logf("x >>> %d", x)
}
