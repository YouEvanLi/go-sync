package sync

import (
	"context"
	"fmt"
	"testing"
	"your_module_path/sync" // 请替换为实际的模块路径
)

func TestAsyncTasks(t *testing.T) {
	// 定义一个简单的任务函数
	taskFunc := func(ctx context.Context, params ...interface{}) (int, error) {
		if len(params) < 1 {
			return 0, fmt.Errorf("no parameters provided")
		}
		num, ok := params[0].(int)
		if !ok {
			return 0, fmt.Errorf("parameter is not an int")
		}
		// 模拟一个可能引发 panic 的操作
		if num == 2 {
			panic("simulated panic")
		}
		return num * 2, nil
	}

	// 创建任务
	task1 := sync.NewTask(taskFunc).WithParams(1).Start(context.Background())
	task2 := sync.NewTask(taskFunc).WithParams(2).Start(context.Background())
	task3 := sync.NewTask(taskFunc).WithParams(3).Start(context.Background())

	// 使用 Async 函数运行并等待所有任务完成
	results, errors := sync.Async(context.Background(), task1, task2, task3)

	// 验证结果
	expectedResults := []int{2, 0, 6} // task2 应该返回 0，因为它会 panic
	for i, result := range results {
		if result != expectedResults[i] {
			t.Errorf("expected result %d, got %d", expectedResults[i], result)
		}
	}

	// 验证错误
	if errors[1] == nil || errors[1].Error() != "panic occurred: simulated panic" {
		t.Errorf("expected panic error for task 2, got %v", errors[1])
	}
	for i, err := range errors {
		if i != 1 && err != nil {
			t.Errorf("unexpected error for task %d: %v", i+1, err)
		}
	}
}
