package pool

import (
	"log/slog"
	"os"
	"sync"
)

// Task - опис інтрефейсу роботи
type Task interface {
	Execute(byWorker int)
}

// Pool - структура, нам знадобиться Мутекс, для гарантій атомарності змін самого об'єкта
// Канал вхідних завдань
// Канал скасування, для завершення роботи
// WaitGroup для контролю завершення робіт
type Pool struct {
	mu     sync.Mutex
	taskNo int
	size   int
	tasks  chan Task
	kill   chan struct{}
	wg     sync.WaitGroup
	log    *slog.Logger
}

// NewPool Сховаємо внутрішній пристрій за конструктором, користувач може впливати тільки на розмір пула
func NewPool(parallelTasks, totalTasks int) *Pool {
	pool := &Pool{
		// Канал завдань - буферизований, щоб основна програма не блокувалася під час постановки завдань
		tasks: make(chan Task, totalTasks),
		// Канал kill для вбивства "зайвих воркерів"
		kill: make(chan struct{}),
		log: slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		})),
	}
	// Викличемо метод resize, щоб встановити відповідний розмір пулу
	pool.Resize(parallelTasks)

	return pool
}

// worker Життєвий цикл воркера
func (p *Pool) worker(no int) {
	defer p.wg.Done()
	p.log.Info("started", "worker", no)
	for {
		select {
		// Якщо є завдання, його потрібно обробити
		case task, ok := <-p.tasks:
			if !ok {
				p.log.Warn("closed", "worker", no)
				return
			}
			p.log.Debug("got new task", "worker", no, "task", task)
			task.Execute(no)
			// Якщо прийшов сигнал помирати, виходимо
		case <-p.kill:
			p.log.Warn("killed", "worker", no)
			return
		}
	}
}

// Resize - зміна розміру пула
func (p *Pool) Resize(n int) {
	// Захоплюємо лок, щоб уникнути одночасної зміни стану
	p.mu.Lock()
	defer p.mu.Unlock()
	p.log.Info("resize", "from", p.size, "to", n)
	for p.size < n {
		p.taskNo++
		p.size++
		p.wg.Add(1)
		go p.worker(p.taskNo)
	}
	for p.size > n {
		p.size--
		p.kill <- struct{}{}
	}
}

// Close - закриваємо пул, закриваємо канал завдань
func (p *Pool) Close() {
	close(p.tasks)
}

// Wait - чекаємо завершення всіх завдань
func (p *Pool) Wait() {
	p.wg.Wait()
}

// Exec - виконуємо завдання (постановка в чергу)
func (p *Pool) Exec(task Task) {
	p.tasks <- task
}
