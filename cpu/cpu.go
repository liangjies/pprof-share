package main

import (
	"fmt"
	"os"
	"runtime/pprof"
)

func main() {
	// 创建性能分析文件
	f, err := os.Create("cpu.prof")
	if err != nil {
		fmt.Println("Error creating profile file:", err)
		return
	}
	// 启动性能分析器
	// pprof.WriteHeapProfile(f) // Memory profile
	err = pprof.StartCPUProfile(f)
	if err != nil {
		fmt.Println("Error starting CPU profiler:", err)
		return
	}
	defer pprof.StopCPUProfile()
	// 执行需要分析时间的代码
	fmt.Println(fib(45))
}

func fib(n int) int {
	if n < 2 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}
