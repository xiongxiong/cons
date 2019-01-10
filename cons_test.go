package cons_test

import (
	"cons"
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestWait(t *testing.T) {
	wg := sync.WaitGroup{}

	x := 0
	wait := func() {
		defer wg.Done()

		c := cons.Wait("hello")
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
	wg := sync.WaitGroup{}

	x := 0
	wait := func() {
		defer wg.Done()

		c := cons.Skip("hello")
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

func TestQueue(t *testing.T) {
	wg := sync.WaitGroup{}

	x := 0
	queue := func() {
		defer wg.Done()

		fmt.Println("waiting...")
		c := cons.Queue("hello")
		defer c.Done()

		fmt.Printf("--- in queue : %d\n", x)
		time.Sleep(1 * time.Second)
		x++
	}

	for range [5]int{} {
		wg.Add(1)
		go queue()
	}
	time.Sleep(3 * time.Second)
	for range [5]int{} {
		wg.Add(1)
		go queue()
	}
	wg.Wait()

	t.Logf("x >>> %d", x)
}
