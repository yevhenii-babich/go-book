package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type PubSub struct {
	mu       sync.Mutex
	subs     map[string][]chan string
	wg       sync.WaitGroup
	msgQueue chan string
	ctx      context.Context
	cancel   context.CancelFunc
}

// NewPubSub створює новий PubSub
func NewPubSub() *PubSub {
	ctx, cancel := context.WithCancel(context.Background())
	return &PubSub{
		subs:     make(map[string][]chan string),
		msgQueue: make(chan string, 10),
		ctx:      ctx,
		cancel:   cancel,
	}
}

// Subscribe додає нового subscriber до теми
func (ps *PubSub) Subscribe(topic string) {
	ps.mu.Lock() // lock mutex
	defer ps.mu.Unlock()

	ch := make(chan string, 5) // create channel with buffer
	ps.subs[topic] = append(ps.subs[topic], ch)
	ps.wg.Add(1) // add subscriber to waitgroup
	workerName := fmt.Sprintf("Підписник %s:%d", topic, len(ps.subs[topic]))
	go func() {
		defer ps.wg.Done() // remove subscriber from waitgroup
		for {
			select {
			case <-ps.ctx.Done(): // check if context is cancelled
				close(ch)
				fmt.Printf("%s закрито\n", workerName)
				return
			case msg := <-ch:
				ps.msgQueue <- msg // send message to queue
				fmt.Printf("Повідомлення отримано [%s]: %s\n", workerName, msg)
			}
		}
	}()
}

// Publish публікує повідомлення до теми
func (ps *PubSub) Publish(topic, msg string) {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	for _, sub := range ps.subs[topic] {
		sub <- msg
	}
}

// Start запускає процес розсилки повідомлень
func (ps *PubSub) Start() {
	const timeoutDuration = time.Second
	ps.wg.Add(1)
	defer ps.wg.Done()
	timeout := time.After(timeoutDuration)
	for {
		select {
		case <-ps.msgQueue:
			timeout = time.After(timeoutDuration) //reset timeout (extend)
		case <-ps.ctx.Done():
			fmt.Println("PubSub закрито (за запитом)")
			return
		case <-timeout:
			fmt.Println("Немає нових повідомлень, закриваємо PubSub...")
			ps.cancel() // cancel context
			return
		}
	}
}

func main() {
	ps := NewPubSub()

	// Створення підписників
	ps.Subscribe("topic1")
	ps.Subscribe("topic1")
	ps.Subscribe("topic2")

	// Створення та запуск Publisher (відправник повідомлень)
	go func() {
		ps.Publish("topic1", "Hello, Topic 1!")
		ps.Publish("topic2", "Hello, Topic 2!")
		ps.Publish("topic1", "Another message for Topic 1")
		fmt.Println("Waiting 100 ms...")
		time.Sleep(100 * time.Millisecond)
		ps.Publish("topic1", "One more message for Topic 1!")
		fmt.Println("Waiting 900 ms...")
		time.Sleep(900 * time.Millisecond) // check for timeout reset
		ps.Publish("topic1", "Next message for Topic 1!")
		fmt.Println("Waiting 900 ms...")
		time.Sleep(900 * time.Millisecond) // check for timeout reset
		ps.Publish("topic1", "Again message for Topic 1!")
	}()

	go ps.Start()                      // запуск PubSub
	time.Sleep(100 * time.Millisecond) // чекаємо 100 мс
	fmt.Println("Adding new subscriber...")
	ps.Subscribe("topic1") // додавання нового підписника
	// Чекаємо завершення підписників
	ps.wg.Wait()
}
