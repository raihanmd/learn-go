package test

import (
	"fmt"
	"goroutines/services"
	"goroutines/utils"
	"runtime"
	"strconv"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// ? Whatever func can be a goroutine with go keyword
func TestGoroutine(t *testing.T) {
	go t.Log("TestGoroutine")
	t.Log("Hello")
}

// ? Channel can have some storage called buffered channel
func TestChannel(t *testing.T) {
	ch := make(chan []utils.DbResult)
	bufferedCh := make(chan []utils.DbResult, 3)
	defer close(ch)
	defer close(bufferedCh)

	service := &services.Service{}

	t.Log(cap(bufferedCh))
	go utils.FetchData(ch, service)
	utils.FetchData(bufferedCh, service)
	utils.FetchData(bufferedCh, service)
	t.Log(len(bufferedCh))

	result := <-ch
	data1 := <-bufferedCh
	data2 := <-bufferedCh

	go t.Log(data1, data2)

	assert.Equal(t, []utils.DbResult{{Id: "1", Name: "John"}, {Id: "2", Name: "Doe"}}, result)
}

// ? Channel can be iterate by for ... range, and auto break the loop when channel is closed, but this method is just on one channel
func TestRangeChannel(t *testing.T) {
	ch := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- "Loop " + strconv.Itoa(i)
		}
		close(ch)
	}()

	for result := range ch {
		t.Log("Menerima data", result)
	}
}

// ? Select can be used with channel, if select not in the loop that only will be select 1 fastest
func TestSelect(t *testing.T) {
	ch1 := make(chan string)
	ch2 := make(chan string)
	defer close(ch1)
	defer close(ch2)

	go func() {
		for i := 1; i <= 10; i++ {
			ch1 <- "Loop CH1 " + strconv.Itoa(i)
		}
	}()
	go func() {
		for i := 1; i <= 15; i++ {
			ch1 <- "Loop CH2 " + strconv.Itoa(i)
		}
	}()

	counter := 0
	for {
		select {
		case data := <-ch1:
			t.Log("Menerima", data)
			counter++
		case data := <-ch2:
			t.Log("Menerima", data)
			counter++
		default:
			t.Log("Wait For Data..")
		}

		if counter == 25 {
			break
		}
	}
}

// ? Mutex can be slowewr, RWMutex dll
func TestMutex(t *testing.T) {
	x := 0
	var mutex sync.Mutex
	for i := 0; i < 100; i++ {
		go func() {
			for i := 0; i < 100; i++ {
				mutex.Lock()
				x++
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(1 * time.Second)
	t.Log("\nData:", x)
	assert.Equal(t, 10000, x)
}

func TestDeadlock(t *testing.T) {
	user1 := utils.BankAccount{Name: "Adit", Balance: 10000}
	user2 := utils.BankAccount{Name: "Gorou", Balance: 100000}

	go utils.Transfer(&user1, &user2, 5000)
	go utils.Transfer(&user2, &user1, 50000)

	fmt.Println("User the result is wrong because they are locking each other xD")
	fmt.Println("User", user1.Name, "Balance", user1.Balance)
	fmt.Println("User", user2.Name, "Balance", user2.Balance)
}

func TestWaitGroup(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(100 * time.Millisecond)
		}()
	}
	wg.Wait()
}

func TestOnce(t *testing.T) {
	var once sync.Once
	var wg sync.WaitGroup
	counter := 0
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			once.Do(func() {
				counter++
			})
		}()
	}
	wg.Wait()
	t.Log(counter)
	assert.Equal(t, 1, counter)
}

func TestPool(t *testing.T) {
	var wg sync.WaitGroup
	var pool sync.Pool
	var poolWithDefaultValue = sync.Pool{
		New: func() any { return "Default" },
	}
	user1 := utils.BankAccount{Name: "Adit", Balance: 10000}
	user2 := utils.BankAccount{Name: "Gorou", Balance: 100000}

	pool.Put(&user1)
	pool.Put(&user2)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			data := pool.Get()
			if data == nil {
				t.Log("I got nil")
				return
			}
			t.Log(data.(*utils.BankAccount).Name)
			pool.Put(data)
		}()
	}

	t.Log("Pool With Default Value:", poolWithDefaultValue.Get())

	wg.Wait()
}

// ? sync.Map is safety for concurrent
func TestSyncMap(t *testing.T) {
	var wg = &sync.WaitGroup{}
	var sm = &sync.Map{}
	var addToMap = func(data *sync.Map, value any, group *sync.WaitGroup) {
		defer group.Done()
		group.Add(1)
		data.Store(value, value)
	}
	for i := 0; i < 100; i++ {
		go addToMap(sm, i, wg)
	}
	wg.Wait()
	// ? this bool return is for go to the next loop if true
	sm.Range(func(key, value any) bool {
		t.Log(key, ":", value)
		return true
	})
}

func TestCond(t *testing.T) {
	cond := sync.NewCond(&sync.Mutex{})
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			cond.L.Lock()
			cond.Wait()
			t.Log(i)
			cond.L.Unlock()
		}(i)
	}
	// ? .Signal() for wake one goroutines, either .Broadcast() for wake all
	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(300 * time.Millisecond)
			cond.Signal()
		}
	}()

	wg.Wait()
}

// ? instead use Mutex like line 99, for primitive data type can use sync/atomic pkg
func TestAtomic(t *testing.T) {
	var counter int64 = 0
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 10000; j++ {
				atomic.AddInt64(&counter, 1)
			}
		}()
	}
	wg.Wait()
	t.Log(counter)
	assert.Equal(t, int64(1000000), counter)
}

func TestTimer(t *testing.T) {
	wg := sync.WaitGroup{}
	var timer = time.NewTimer(1 * time.Second)
	t.Log(time.Now().Local())
	myTime := <-timer.C
	t.Log(myTime.Local())

	timeAfter := time.After(1 * time.Second)
	tAfter := <-timeAfter
	t.Log(tAfter.Local())

	wg.Add(1)
	// ? this is automaticly goroutine
	time.AfterFunc(1*time.Second, func() {
		defer wg.Done()
		go t.Log("Hello World")
	})
	wg.Wait()

	ticker := time.NewTicker(1 * time.Second)
	go func() {
		defer t.Log("Stop")
		for tick := range ticker.C {
			t.Log("Tick", tick)
		}
	}()
	time.Sleep(3 * time.Second)
	t.Log("Stopping ticker")
	ticker.Stop()

	t.Log("Total  thread:", runtime.GOMAXPROCS(-1))
}
