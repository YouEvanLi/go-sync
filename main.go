package sync

import (
	"context"
	"fmt"
	"sync"
)

// Task 表示一个异步任务，支持可选且不定类型的参数
type Task[T any] struct {
	fn     func(context.Context, ...interface{}) (T, error) // 任务的执行函数，接收 context 和不定数量的参数
	params []interface{}                                    // 存储参数
	result chan result[T]                                   // 用于返回任务执行的结果或错误
}

// result 用于封装返回的结果和错误
type result[T any] struct {
	data T
	err  error
}

// NewTask 创建一个新的异步任务，接收可选的参数
func NewTask[T any](fn func(context.Context, ...interface{}) (T, error)) *Task[T] {
	return &Task[T]{fn: fn, result: make(chan result[T], 1)}
}

// WithParams 设置任务的参数
func (t *Task[T]) WithParams(params ...interface{}) *Task[T] {
	t.params = params
	return t
}

// Start 启动任务
func (t *Task[T]) Start(ctx context.Context) *Task[T] {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				t.result <- result[T]{err: fmt.Errorf("panic occurred: %v", r)}
			}
		}()
		// 执行任务并传入 context 和可变参数
		data, err := t.fn(ctx, t.params...)
		t.result <- result[T]{data: data, err: err}
	}()
	return t
}

// Await 等待任务执行完成并返回结果
func (t *Task[T]) Await() (T, error) {
	res := <-t.result
	return res.data, res.err
}

// Async 运行多个异步任务并等待结果
func Async[T any](ctx context.Context, tasks ...*Task[T]) ([]T, []error) {
	var wg sync.WaitGroup
	results := make([]T, len(tasks))
	errors := make([]error, len(tasks))
	var mu sync.Mutex // 用于保护 results 和 errors

	// 启动所有任务
	for i, task := range tasks {
		wg.Add(1)
		go func(i int, task *Task[T]) {
			defer wg.Done()
			data, err := task.Await()
			mu.Lock()
			results[i] = data
			errors[i] = err
			mu.Unlock()
		}(i, task)
	}

	// 等待所有任务完成
	wg.Wait()

	return results, errors
}
