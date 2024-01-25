package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"sync"
	"time"
)

type (
	// MyCache is a simple memory cache with expiration time and cleanup
	MyCache[T any] struct {
		log               *slog.Logger            // logger
		items             map[string]cacheItem[T] // cache items
		defaultExpiration time.Duration           // default expiration time
		ctx               context.Context         // context for gracefull shutdown
		mu                sync.RWMutex            // mutex for concurrent access
	}
	cacheItem[T any] struct {
		item    T     // item to store
		expired int64 // expiration time in nanoseconds
	}
)

// NewCache creates new cache with default expiration time
func NewCache[T any](ctx context.Context, exp, check time.Duration, handler slog.Handler) *MyCache[T] {
	if handler == nil {
		handler = slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelDebug})
	}
	log := slog.New(handler) // create logger
	tmp := &MyCache[T]{      // create cache instance as pointer to struct
		items:             make(map[string]cacheItem[T]), // create map
		log:               log,                           // set logger
		defaultExpiration: exp,                           // set default expiration time
		ctx:               ctx,                           // set context
	}
	if check > 0 { // if cleanup time is set
		go tmp.runner(check) // run cleaner
	}
	return tmp // return cache instance
}

func (mc *MyCache[T]) runner(check time.Duration) {
	mc.log.Info("Start cache cleaner")
	ticker := time.NewTicker(check) // create ticker with cleanup time
	for {
		select {
		case <-ticker.C: // on ticker
			cleared := mc.Cleanup()                   // cleanup cache
			mc.log.Debug("cleaned", "items", cleared) // log cleanup result
		case <-mc.ctx.Done(): // on context done
			mc.log.Info("Stop cache cleaner")
			ticker.Stop() // stop ticker
			return        // exit
		}
	}
}

func (mc *MyCache[T]) Set(key string, value T) {
	mc.mu.Lock()                  // lock for concurrent access (write)
	defer mc.mu.Unlock()          // unlock on exit
	mc.items[key] = cacheItem[T]{ // set item
		item:    value,                                           // set value
		expired: time.Now().Add(mc.defaultExpiration).UnixNano()} // set expiration time in nanoseconds
}

func (mc *MyCache[T]) Get(key string) (T, bool) {
	mc.mu.RLock()                // lock for concurrent access (read)
	defer mc.mu.RUnlock()        // unlock on exit
	item, found := mc.items[key] // get item
	var empty T                  // empty value
	if !found {
		return empty, found // early exit if item not found
	}
	if time.Now().UnixNano() > item.expired {
		// if item is expired - return empty value and false as found
		return empty, false
	}
	return item.item, found // return item value and true as found
}
func (mc *MyCache[T]) Delete(key string) {
	mc.mu.Lock()          // lock for concurrent access (write)
	defer mc.mu.Unlock()  // unlock on exit
	delete(mc.items, key) // delete item from map
}

func (mc *MyCache[T]) Cleanup() int {
	mc.mu.Lock()         // lock for concurrent access (write)
	defer mc.mu.Unlock() // unlock on exit
	cur := len(mc.items) // get current items count
	for s, c := range mc.items {
		if time.Now().UnixNano() > c.expired {
			delete(mc.items, s) // delete expired item
		}
	}
	return cur - len(mc.items) // return cleared items count
}

func cacheWriter(cache *MyCache[string], cancel context.CancelFunc, out chan<- string) {
	for i := 0; i < 100; i++ {
		cache.Set(fmt.Sprintf("key_%d", i), fmt.Sprintf("value_%d", i)) // set item
		out <- fmt.Sprintf("key_%d", i)                                 // send key to channel reader
		time.Sleep(100 * time.Millisecond)                              // sleep for 100 milliseconds
	}
	close(out) // close channel on completion
	cancel()   // cancel context (stop cleaner and readers)
}
func cacheReader(in <-chan string, out chan<- string, cache *MyCache[string]) {
	for {
		select { // select on channel
		case key, ok := <-in: // read key from channel
			if !ok { // if channel closed
				slog.Warn("channel closed", "reader", "1")
				return
			}
			out <- key                  // send key to channel reader
			v, ok := cache.Get(key)     // get item from cache
			fmt.Println("<<1\t", v, ok) // print item
		case <-cache.ctx.Done(): // on context done
			close(out) // close channel
			slog.Warn("cacheReader: context done")
			return
		}
	}
}

func cacheReader2(in <-chan string, cache *MyCache[string]) {
	ticker := time.NewTicker(150 * time.Millisecond) // create ticker with read time
	var keys []string                                // create keys slice
	for {
		select {
		case key, ok := <-in: // collect keys
			if !ok { // if channel closed
				slog.Warn("channel closed", "reader", "2")
				return
			}
			keys = append(keys, key) // append key to slice
		case <-ticker.C: // read keys
			for _, key := range keys { // read keys from slice
				v, ok := cache.Get(key)          // get item from cache
				fmt.Println("<<2\t", key, v, ok) // print item
			}
			keys = nil // reset
		case <-cache.ctx.Done(): // on context done
			slog.Warn("cacheReader2: context done")
			return
		}
	}
}

func main() {
	bkg := context.Background()                  // create background context
	bkg = context.WithValue(bkg, "key", "value") // add value to context
	ctx, cancel := context.WithCancel(bkg)       // create context with cancel function
	m, _ := bkg.Value("key").(string)            // get value from context
	fmt.Println(m)
	// create string cache with expiration time 50 milliseconds and cleanup time 500 milliseconds
	cache := NewCache[string](ctx, 50*time.Millisecond, 500*time.Millisecond, nil)
	ch := make(chan string, 10)        // create channel
	ch2 := make(chan string, 10)       // create channel
	go cacheWriter(cache, cancel, ch)  // start cache writer
	go cacheReader(ch, ch2, cache)     // start cache reader
	go cacheReader2(ch2, cache)        // start cache reader
	<-ctx.Done()                       // wait for context done
	time.Sleep(100 * time.Millisecond) // sleep for 100 milliseconds to allow for cleanup
}
