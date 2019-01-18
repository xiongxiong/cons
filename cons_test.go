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
			fmt.Printf("--- sleeping")
			time.Sleep(1 * time.Second)
			x++
			return
		}
		fmt.Printf("awake")
	}

	for range [100]int{} {
		wg.Add(1)
		go wait()
	}
	wg.Wait()

	fmt.Printf("x >>> %d\n", x)
}

func ExampleWait() {
	wg := sync.WaitGroup{}

	x := 0
	wait := func() {
		defer wg.Done()

		c := cons.Wait("hello")
		defer c.Done()

		if x == 0 {
			fmt.Println("--- sleeping")
			time.Sleep(1 * time.Second)
			x++
			return
		}
		fmt.Println("awake")
	}

	for range [100]int{} {
		wg.Add(1)
		go wait()
	}
	wg.Wait()

	fmt.Printf("x >>> %d\n", x)
}

func TestSkip(t *testing.T) {
	wg := sync.WaitGroup{}

	x := 0
	skip := func() {
		defer wg.Done()

		c := cons.Skip("hello")
		defer c.Done()

		if c.Skip {
			fmt.Println("skip")
			return
		}
		fmt.Println("--- doing")
		time.Sleep(1 * time.Second)
		fmt.Println("--- done")
		x++
	}

	for range [100]int{} {
		wg.Add(1)
		go skip()
	}
	wg.Wait()

	fmt.Printf("x >>> %d\n", x)
}

func TestSkipCurrent(t *testing.T) {
	fa := func(wg *sync.WaitGroup, cs *cons.Cons, flag string, index int) {
		defer wg.Done()

		c := cs.Skip("hello")
		defer c.Done()

		if c.Skip {
			println("skip -- ", flag, index)
			return
		}

		println("doing -- ", flag, index)
		time.Sleep(3 * time.Second)
		println("done -- ", flag, index)
	}

	csa := cons.GetCons()
	wga := sync.WaitGroup{}

	wga.Add(1)
	go func() {
		defer wga.Done()
		for i := 0; i < 10; i++ {
			wga.Add(1)
			go fa(&wga, csa, "a", i)
		}
	}()

	csb := cons.GetCons()
	wgb := sync.WaitGroup{}

	wgb.Add(1)
	go func() {
		defer wgb.Done()
		for i := 0; i < 10; i++ {
			wgb.Add(1)
			go fa(&wgb, csb, "b", i)
		}
	}()

	wga.Wait()
	wgb.Wait()
}

func ExampleSkip() {
	wg := sync.WaitGroup{}

	x := 0
	skip := func() {
		defer wg.Done()

		c := cons.Skip("hello")
		defer c.Done()

		if c.Skip {
			fmt.Println("skip")
			return
		}
		fmt.Println("--- doing")
		time.Sleep(1 * time.Second)
		x++
	}

	for range [100]int{} {
		wg.Add(1)
		go skip()
	}
	wg.Wait()

	fmt.Printf("x >>> %d\n", x)
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

	fmt.Printf("x >>> %d\n", x)
}

func TestQueueAndSkip(t *testing.T) {
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

	wg.Add(1)
	go func() {
		defer wg.Done()
		for range [5]int{} {
			wg.Add(1)
			go queue()
		}
	}()

	skip := func() {
		defer wg.Done()

		c := cons.Skip("hello")
		defer c.Done()

		if c.Skip {
			fmt.Println("skip")
			return
		}
		fmt.Println("--- doing")
		time.Sleep(1 * time.Second)
		fmt.Println("--- done")
		x++
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for range [10]int{} {
			wg.Add(1)
			go skip()
		}
	}()

	wg.Wait()

	fmt.Printf("x >>> %d\n", x)
}

func ExampleQueue() {
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

	fmt.Printf("x >>> %d\n", x)
}
